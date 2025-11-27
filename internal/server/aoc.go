//go:generate go tool oapi-codegen -config oapi-codegen.yml ../../api/openapi-go-away-2024.yml

package server

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	Task
}

func NewAoCServer() *Server {
	return &Server{
		Task: Task{
			Id:          "1",
			Status:      CREATED,
			Result:      nil,
			Message:     "Day 1 Part 1 from AoC 2024 accepted",
			CreatedAt:   time.Now(),
			StartedAt:   nil,
			CompletedAt: nil,
		},
	}
}

func SendServerError(c *fiber.Ctx, code int, message string) error {
	serverErr := Error{
		Code:    int32(code),
		Message: message,
	}
	return c.Status(code).JSON(serverErr)
}

func (s *Server) PostTask(c *fiber.Ctx, params PostTaskParams) error {
	// logic is under construction)
	return c.Status(http.StatusOK).JSON(s.Task)
}

func (s *Server) GetTask(c *fiber.Ctx, id string) error {
	// logic is under construction)
	return c.Status(http.StatusOK).JSON(s.Task)
}
