package main

import (
	"go-away-2024/internal/config"
	"go-away-2024/internal/database"
	"go-away-2024/internal/server"
	"go-away-2024/internal/s3"
	"net"

	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.Load()
	database.Connect()
	s3.CreateClient()

	aoc := server.NewAoCServer()
	app := fiber.New()
	app.Use(logger.New())
	server.RegisterHandlers(app, aoc)
	addr := net.JoinHostPort(config.ServerCfg.Host, config.ServerCfg.Port)
	log.Fatal(app.Listen(addr))
}
