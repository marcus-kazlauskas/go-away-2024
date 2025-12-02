package main

import (
	"go-away-2024/internal/config"
	"go-away-2024/internal/database"
	"go-away-2024/internal/minio"
	"go-away-2024/internal/server"
	"net"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	config.Load()
	database.Connect()
	minio.CreateClient()

	adventOfCodeServer := server.NewServer()
	app := server.NewServerApp(adventOfCodeServer)
	addr := net.JoinHostPort(config.ServerCfg.Host, config.ServerCfg.Port)
	log.Fatal(app.Listen(addr))
}
