package server

import (
	"fmt"
	"go-away-2024/internal/database"
	"time"
)

func SaveRequest(p PostTaskParams) (*TaskResponse, error) {
	request := toRequestEntity(p)
	id, err := database.SaveRequest(request)
	if err != nil {
		return nil, err
	}
	request.Id = id
	return toTaskCreatedResponse(request), err
}

func toRequestEntity(p PostTaskParams) database.RequestEntity {
	return database.RequestEntity{
		Year:      p.Year,
		Day:       p.Day,
		Part:      p.Part,
		CreatedAt: time.Now(),
	}
}

func toTaskCreatedResponse(e database.RequestEntity) *TaskResponse {
	return &TaskResponse{
		Id:        e.Id,
		Status:    CREATED,
		CreatedAt: e.CreatedAt,
		Message:   fmt.Sprintf("Day %d Part %d from AoC %d accepted", e.Day, e.Part, e.Year),
	}
}
