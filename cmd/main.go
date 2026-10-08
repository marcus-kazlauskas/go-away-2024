package main

import (
	"go-away-2024/internal/aoc_calc"
	"go-away-2024/internal/aoc_server"
	"go-away-2024/internal/config"
	"go-away-2024/internal/database"
	"go-away-2024/internal/kafka"
	"go-away-2024/internal/minio"
	"net"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	log.SetLevel(log.LevelInfo)

	config := config.GetConfig(config.MAIN_PATH)

	repository := database.NewRepository(config)
	minio := minio.NewClient(config)
	kafka := kafka.NewKafkaConnection(config)

	adventOfCodeServer := aoc_server.NewServer(repository, minio, kafka)
	app := aoc_server.NewServerApp(adventOfCodeServer)
	addr := net.JoinHostPort(config.Server.Host, config.Server.Port)

	adventOfCodeCalculator := aoc_calc.NewCalculator(repository, minio, kafka, config)

	errCh := make(chan error, 2)
	go func() { errCh <- app.Listen(addr) }()
	go func() { errCh <- adventOfCodeCalculator.Start() }()
	log.Fatal(<-errCh)
}
