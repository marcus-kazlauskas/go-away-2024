package main

import (
	"flag"
	"go-away-2024/internal/server"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := flag.String("port", "8090", "Port for Advent of Code server")
	flag.Parse()
	s := server.NewAoCServer()
	app := fiber.New()
	server.RegisterHandlers(app, s)
	log.Fatal(app.Listen(net.JoinHostPort("0.0.0.0", *port)))
}
