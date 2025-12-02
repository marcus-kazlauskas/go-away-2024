//go:generate go tool oapi-codegen -config server-codegen.yml ../../api/openapi-go-away-2024.yml
//go:generate go tool oapi-codegen -config types-codegen.yml ../../api/openapi-go-away-2024.yml

package server

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	middleware "github.com/oapi-codegen/fiber-middleware"
)

var _ ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func NewServerApp(s *Server) *fiber.App {
	swagger, err := GetSwagger()
	if err != nil {
		log.Fatalf("Error loading swagger spec: %v", err)
	}
	swagger.Servers = nil

	app := fiber.New()
	app.Use(logger.New())
	app.Use(middleware.OapiRequestValidator(swagger))
	RegisterHandlers(app, s)
	return app
}

func SendServerError(c *fiber.Ctx, code int, message string) error {
	serverErr := ErrorResponse{
		Code:    int32(code),
		Message: message,
	}
	return c.Status(code).JSON(serverErr)
}

func (s *Server) PostTask(c *fiber.Ctx, params PostTaskParams) error {
	response, err := SaveRequest(c, params)
	if err != nil {
		return SendServerError(c, http.StatusInternalServerError, fmt.Sprintf("%v", err))
	}
	return c.Status(http.StatusOK).JSON(response)
}

func (s *Server) GetTask(c *fiber.Ctx, id int64) error {
	response, err := GetRequestWithResult(id)
	if err != nil {
		return SendServerError(c, http.StatusInternalServerError, fmt.Sprintf("%v", err))
	}
	return c.Status(http.StatusOK).JSON(response)
}
