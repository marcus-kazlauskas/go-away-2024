//go:generate go tool oapi-codegen -config oapi-codegen.yml ../../api/openapi-go-away-2024.yml

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	TaskResponse
}

func NewAoCServer() *Server {
	return &Server{
		TaskResponse: TaskResponse{
			Id:          1312,
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
	serverErr := ErrorResponse{
		Code:    int32(code),
		Message: message,
	}
	return c.Status(code).JSON(serverErr)
}

func (s *Server) PostTask(c *fiber.Ctx, params PostTaskParams) error {
	response, err := SaveRequest(params)
	if err != nil {
		return SendServerError(c, http.StatusInternalServerError, fmt.Sprintf("%v", err))
	}
	return c.Status(http.StatusOK).JSON(response)
}

func (s *Server) GetTask(c *fiber.Ctx, id int64) error {
	// logic is under construction)
	return c.Status(http.StatusOK).JSON(s.TaskResponse)
}
