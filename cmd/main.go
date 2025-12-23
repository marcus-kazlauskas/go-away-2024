package main

import (
	"go-away-2024/internal/aoc_calc"
	"go-away-2024/internal/aoc_server"
	"go-away-2024/internal/config"
	"go-away-2024/internal/database"
	"go-away-2024/internal/kafka"
	"go-away-2024/internal/minio"
	"net"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	log.SetLevel(log.LevelInfo)
	config := config.GetConfig(config.MainPath)
	repository := database.NewRepository(config)
	minio := minio.NewClient(config)
	kafka := kafka.NewKafkaConnection(config)

	adventOfCodeServer := aoc_server.NewServer(repository, minio, kafka)
	app := aoc_server.NewServerApp(adventOfCodeServer)
	addr := net.JoinHostPort(config.Server.Host, config.Server.Port)

	adventOfCodeCalculator := aoc_calc.NewCalculator(repository, minio, kafka, config)

	var wg sync.WaitGroup
	wg.Add(2)
	go startServer(&wg, app, addr)
	go startCalculator(&wg, adventOfCodeCalculator)
	wg.Wait()
}

func startServer(wg *sync.WaitGroup, app *fiber.App, addr string) {
	defer wg.Done()
	log.Fatal(app.Listen(addr))
}

func startCalculator(wg *sync.WaitGroup, calc *aoc_calc.Calculator) {
	defer wg.Done()
	log.Fatal(calc.Start())
}
