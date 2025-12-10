package mappers

import (
	"fmt"
	"go-away-2024/internal/api"
	"go-away-2024/internal/database"
	"time"
)

func PostTaskParamsToRequestEntity(p api.PostTaskParams) database.RequestEntity {
	return database.RequestEntity{
		Year:      p.Year,
		Day:       p.Day,
		Part:      p.Part,
		CreatedAt: time.Now(),
	}
}

func RequestEntityToTaskCreatedResponse(e database.RequestEntity) api.TaskResponse {
	return api.TaskResponse{
		Id:        e.Id,
		Status:    api.CREATED,
		CreatedAt: e.CreatedAt,
		Message:   yearDayPartToMessage(e.Day, e.Part, e.Year),
	}
}

func yearDayPartToMessage(year int32, day int32, part int32) string {
	return fmt.Sprintf("Day %d Part %d from AoC %d accepted", day, part, year)
}

func RequestWithResultEntityToTaskCreatedResponse(e database.RequestWithResultEntity) api.TaskResponse {
	var status api.TaskResponseStatus
	if e.Status == nil {
		status = api.CREATED
	} else {
		status = api.TaskResponseStatus(*e.Status)
	}
	return api.TaskResponse{
		Id:          e.RequestId,
		Status:      status,
		CreatedAt:   e.CreatedAt,
		StartedAt:   e.StartedAt,
		CompletedAt: e.CompletedAt,
		Result:      e.Result,
		Message:     yearDayPartToMessage(e.Day, e.Part, e.Year),
	}
}
