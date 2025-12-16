package aoc_server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"go-away-2024/internal/api"
	"go-away-2024/internal/database"
	"go-away-2024/internal/kafka"
	"go-away-2024/internal/mappers"
	"go-away-2024/internal/minio"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	middleware "github.com/oapi-codegen/fiber-middleware"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	repository      *database.Repository
	minioClient     *minio.MinioClient
	kafkaConnection *kafka.KafkaConnection
}

func NewServer(
	repo *database.Repository,
	minio *minio.MinioClient,
	kafka *kafka.KafkaConnection,
) *Server {
	log.Info("Server created")
	return &Server{
		repository:      repo,
		minioClient:     minio,
		kafkaConnection: kafka,
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
	request, err := s.saveRequest(params)
	if err != nil {
		return err
	}

	s3Link, err := s.uploadPuzzleInput(c, params, request.Id)
	if err != nil {
		return err
	}

	request.S3Link = &s3Link
	err = s.writeTask(request)
	if err != nil {
		return err
	}

	response := mappers.RequestEntityToTaskCreatedResponse(request)
	return c.Status(http.StatusOK).JSON(response)
}

func (s *Server) saveRequest(p api.PostTaskParams) (database.RequestEntity, error) {
	request := mappers.PostTaskParamsToRequestEntity(p)
	id, err := s.repository.SaveRequest(request)
	if err != nil {
		return request, err
	}
	request.Id = id
	return request, err
}

func (s *Server) uploadPuzzleInput(c *fiber.Ctx, p api.PostTaskParams, id int64) (string, error) {
	pattern := minio.NewPattern(id, p.Year, p.Day, p.Part)
	tmpFile, err := os.CreateTemp("", pattern)
	if err != nil {
		return pattern, err
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Write(c.Body())

	err = s.minioClient.UploadPuzzleInput(pattern, tmpFile)
	if err != nil {
		return pattern, err
	}

	return pattern, s.repository.UpdateRequestS3Link(id, pattern)
}

func (s *Server) writeTask(rq database.RequestEntity) error {
	msg := mappers.RequestEntityToTaskMessage(rq)
	err := s.kafkaConnection.WriteTask(&msg)
	if err != nil {
		return err
	}

	return s.repository.SaveResult(rq.Id)
}

// Get task status
// (GET /task/{id})
func (s *Server) GetTask(c *fiber.Ctx, id int64) error {
	rqRes, err := s.getRequestWithResult(id)
	if err != nil {
		return SendServerError(c, err)
	}
	response := mappers.RequestWithResultEntityToTaskCreatedResponse(rqRes)
	return c.Status(http.StatusOK).JSON(response)
}

func (s *Server) getRequestWithResult(id int64) (database.RequestWithResultEntity, error) {
	rqRes, err := s.repository.GetRequestWithResult(id)
	if err != nil {
		return rqRes, err
	}
	return rqRes, err
}
