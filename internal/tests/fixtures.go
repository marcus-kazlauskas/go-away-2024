package tests

import (
	"go-away-2024/internal/database"
	"time"

	"github.com/jmoiron/sqlx"
)

const MaxInsertedRows int32 = 10

// Save new SaveRequest into database
func SaveRequest(
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

func DeleteRequest(db *sqlx.DB, id int64) {
	db.MustExec(`delete from request where id = $1`, id)
}
