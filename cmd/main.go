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
	config := config.GetConfig(config.MainPath)
	repository := database.NewRepository(config)
	minio := minio.NewClient(config)
	kafka := kafka.NewKafkaConnection(config)

	adventOfCodeServer := aoc_server.NewServer(repository, minio, kafka)
	app := aoc_server.NewServerApp(adventOfCodeServer)

	adventOfCodeCalculator := aoc_calc.NewCalculator(repository, minio, kafka, config)

	// TODO: run two continuous functions in two goroutines
	adventOfCodeCalculator.Start()

	addr := net.JoinHostPort(config.Server.Host, config.Server.Port)
	log.Fatal(app.Listen(addr))
}
