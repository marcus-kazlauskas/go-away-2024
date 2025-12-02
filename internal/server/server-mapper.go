package server

import (
	"fmt"
	"go-away-2024/internal/database"
	"time"
)

func postTaskParamsToRequestEntity(p PostTaskParams) database.RequestEntity {
	return database.RequestEntity{
		Year:      p.Year,
		Day:       p.Day,
		Part:      p.Part,
		CreatedAt: time.Now(),
	}
}

func requestEntityToTaskCreatedResponse(e database.RequestEntity) *TaskResponse {
	return &TaskResponse{
		Id:        e.Id,
		Status:    CREATED,
		CreatedAt: e.CreatedAt,
		Message:   yearDayPartToMessage(e.Day, e.Part, e.Year),
	}
}

func yearDayPartToMessage(year int32, day int32, part int32) string {
	return fmt.Sprintf("Day %d Part %d from AoC %d accepted", day, part, year)
}

func requestWithResultEntityToTaskCreatedResponse(e database.RequestWithResultEntity) *TaskResponse {
	var status TaskResponseStatus
	if e.Status == nil {
		status = CREATED
	} else {
		status = TaskResponseStatus(*e.Status)
	}
	return &TaskResponse{
		Id: e.RequestId,
		Status: status,
		CreatedAt: e.CreatedAt,
		StartedAt: e.StartedAt,
		CompletedAt: e.CompletedAt,
		Result: e.Result,
		Message: yearDayPartToMessage(e.Day, e.Part, e.Year),
	}
}
