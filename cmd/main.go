package main

import (
	"go-away-2024/internal/aoc_server"
	"go-away-2024/internal/config"
	"go-away-2024/internal/database"
	"go-away-2024/internal/minio"
	"net"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	config := config.GetConfig(config.MainPath)
	repository := database.NewRepository(database.Connect(config))
	minio := minio.NewClient(config)

	adventOfCodeServer := aoc_server.NewServer(repository, minio)
	app := aoc_server.NewServerApp(adventOfCodeServer)
	addr := net.JoinHostPort(config.Server.Host, config.Server.Port)
	log.Fatal(app.Listen(addr))
}
