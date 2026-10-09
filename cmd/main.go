package main

import (
	"context"
	"fmt"
	"go-away-2024/internal/aoc_calc"
	"go-away-2024/internal/aoc_server"
	"go-away-2024/internal/config"
	"go-away-2024/internal/database"
	"go-away-2024/internal/kafka"
	"go-away-2024/internal/minio"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.SetLevel(log.LevelInfo)

	if err := run(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	log.Info("Stopped")
}

func run() error {
	config := config.GetConfig(config.MAIN_PATH)
	shutdownTimeout, err := time.ParseDuration(config.Server.ShutdownTimeout)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	repository := database.NewRepository(config)
	minio := minio.NewClient(config)
	kafkaWriter := kafka.NewTaskWriter(config)
	kafkaReader := kafka.NewTaskReader(config)

	adventOfCodeServer := aoc_server.NewServer(repository, minio, kafkaWriter)
	app := aoc_server.NewServerApp(adventOfCodeServer)
	addr := net.JoinHostPort(config.Server.Host, config.Server.Port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	adventOfCodeCalculator := aoc_calc.NewCalculator(repository, minio, kafkaReader)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error { return app.Listener(ln) })
	g.Go(func() error {
		<-ctx.Done()
		log.Info("Shutting down")
		err := app.ShutdownWithTimeout(shutdownTimeout)
		// stops Listener if shutdown came before it started serving
		_ = ln.Close()
		return err
	})
	g.Go(func() error { return adventOfCodeCalculator.Start(ctx) })

	err = g.Wait()

	if err := kafkaWriter.Close(); err != nil {
		log.Error(err)
	}
	if err := kafkaReader.Close(); err != nil {
		log.Error(err)
	}
	if err := repository.Close(); err != nil {
		log.Error(err)
	}
	return err
}
