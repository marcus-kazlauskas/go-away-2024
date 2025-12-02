package database

import (
	"time"
)

type RequestEntity struct {
	Id        int64     `db:"id"`
	Year      int32     `db:"year"`
	Day       int32     `db:"day"`
	Part      int32     `db:"part"`
	CreatedAt time.Time `db:"created_at"`
	S3Link    *string   `db:"s3_link"`
}

type ResultEntity struct {
	Id          int64      `db:"id"`
	Status      string     `db:"status"`
	Result      *string    `db:"result"`
	StartedAt   *time.Time `db:"started_at"`
	CompletedAt *time.Time `db:"completed_at"`
	RequestId   int64      `db:"request_id"`
}

type RequestWithResultEntity struct {
	RequestId   int64      `db:"request_id"`
	Year        int32      `db:"year"`
	Day         int32      `db:"day"`
	Part        int32      `db:"part"`
	CreatedAt   time.Time  `db:"created_at"`
	StartedAt   *time.Time `db:"started_at"`
	CompletedAt *time.Time `db:"completed_at"`
	Status      *string    `db:"status"`
	Result      *string    `db:"result"`
	S3Link      *string    `db:"s3_link"`
}
