package utils

import (
	"go-away-2024/internal/database"
	"go-away-2024/internal/kafka"
)

func RequestEntityToTaskMessage(e database.RequestEntity) kafka.TaskMessage {
	return kafka.TaskMessage{
		Id:     e.Id,
		Year:   e.Year,
		Day:    e.Day,
		Part:   e.Part,
		S3Link: e.S3Link,
	}
}
