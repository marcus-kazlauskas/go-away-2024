package server

import (
	"go-away-2024/internal/database"

	"github.com/gofiber/fiber/v2"
)

func SaveRequest(c *fiber.Ctx, p PostTaskParams) (*TaskResponse, error) {
	request := postTaskParamsToRequestEntity(p)
	id, err := database.SaveRequest(request)
	if err != nil {
		return nil, err
	}
	request.Id = id
	return requestEntityToTaskCreatedResponse(request), err
}

func GetRequestWithResult(id int64) (*TaskResponse, error) {
	rqRes, err := database.GetRequestWithResult(id)
	if err != nil {
		return nil, err
	}
	return requestWithResultEntityToTaskCreatedResponse(rqRes), err
}
