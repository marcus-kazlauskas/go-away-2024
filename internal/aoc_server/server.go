package aoc_server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"go-away-2024/internal/api"
	"go-away-2024/internal/database"
	"go-away-2024/internal/mappers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/minio/minio-go/v7"
	middleware "github.com/oapi-codegen/fiber-middleware"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	Repository  *database.Repository
	MinioClient *minio.Client
}

func NewServer(repo *database.Repository, minio *minio.Client) *Server {
	return &Server{
		Repository:  repo,
		MinioClient: minio,
	}
}

func NewServerApp(s *Server) *fiber.App {
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("Error loading swagger spec: %v", err)
	}
	swagger.Servers = nil

	app := fiber.New()
	app.Use(logger.New())
	app.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(app, s)
	return app
}

func SendServerError(c *fiber.Ctx, err error) error {
	var code int
	var message string
	switch {
	case errors.Is(err, sql.ErrNoRows):
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}
	message = fmt.Sprintf("%v", err)

	serverErr := api.ErrorResponse{
		Code:    int32(code),
		Message: message,
	}
	return c.Status(code).JSON(serverErr)
}

// Create task to solve
// (POST /task/create)
func (s *Server) PostTask(c *fiber.Ctx, params api.PostTaskParams) error {
	response, err := s.saveRequest(c, params)
	if err != nil {
		return SendServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON(response)
}

func (s *Server) saveRequest(c *fiber.Ctx, p api.PostTaskParams) (*api.TaskResponse, error) {
	request := mappers.PostTaskParamsToRequestEntity(p)
	id, err := s.Repository.SaveRequest(request)
	if err != nil {
		return nil, err
	}
	request.Id = id
	return mappers.RequestEntityToTaskCreatedResponse(request), err
}

// Get task status
// (GET /task/{id})
func (s *Server) GetTask(c *fiber.Ctx, id int64) error {
	response, err := s.getRequestWithResult(id)
	if err != nil {
		return SendServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON(response)
}

func (s *Server) getRequestWithResult(id int64) (*api.TaskResponse, error) {
	rqRes, err := s.Repository.GetRequestWithResult(id)
	if err != nil {
		return nil, err
	}
	return mappers.RequestWithResultEntityToTaskCreatedResponse(rqRes), err
}
