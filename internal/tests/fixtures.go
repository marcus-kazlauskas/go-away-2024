package tests

import (
	"go-away-2024/internal/database"
	"time"
)

const MaxInsertedRows int32 = 10

// Save new Request into database
func Request(
	r *database.Repository,
	year int32, day int32, part int32, createdAt time.Time,
) database.RequestEntity {
	request := database.RequestEntity{
		Year:      year,
		Day:       day,
		Part:      part,
		CreatedAt: createdAt,
	}
	id, err := r.SaveRequest(request)
	if err != nil {
		panic(err)
	}
	request.Id = id
	return request
}
