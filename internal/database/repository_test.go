package database

import (
	"fmt"
	"go-away-2024/internal/api"
	"go-away-2024/internal/config"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	a := assert.New(t)

	config := config.GetConfig(config.TEST_PATH)
	db := Connect(config)
	repository := NewRepository(config)

	requestIds := make([]int64, 0)
	addRequestId := func(id int64) {
		requestIds = append(requestIds, id)
	}
	// delete results cascade
	deleteRequests := func() {
		for i := range len(requestIds) {
			DeleteRequest(db, requestIds[i])
		}
	}
	defer deleteRequests()

	requestEntity := RequestEntity{
		Year:      2024,
		Day:       1,
		Part:      1,
		CreatedAt: time.Now(),
	}
	s3Link := "S3Link"

	resultEntity := ResultEntity{}
	status := fmt.Sprint(api.COMPLETED)
	result := "answer"
	startedAt := time.Now()
	completedAt := time.Now()

	t.Run("Should save request", func(t *testing.T) {
		id, err := repository.SaveRequest(requestEntity)

		a.Nil(err)

		requestEntity.Id = id
		resultEntity.RequestId = id
		addRequestId(id)
	})

	t.Run("Should update request S3 link", func(t *testing.T) {
		err := repository.UpdateRequestS3Link(requestEntity.Id, s3Link)

		a.Nil(err)

		requestEntity.S3Link = &s3Link
	})

	t.Run("Should save result", func(t *testing.T) {
		err := repository.SaveResult(resultEntity.RequestId)

		a.Nil(err)
	})

	t.Run("Should set result", func(t *testing.T) {
		resultEntity.Status = status
		resultEntity.Result = &result
		resultEntity.StartedAt = &startedAt
		resultEntity.CompletedAt = &completedAt

		err := repository.SetResult(resultEntity)

		a.Nil(err)
	})

	t.Run("Should find updated result", func(t *testing.T) {
		res, err := repository.GetResult(requestEntity.Id)

		a.Nil(err)
		a.NotNil(res)
		a.Equal(resultEntity.Status, res.Status)
		a.Equal(*resultEntity.Result, *res.Result)
		a.Equal(resultEntity.StartedAt.Local(), res.StartedAt.Local())
		a.Equal(resultEntity.CompletedAt.Local(), res.CompletedAt.Local())
	})

	t.Run("Should find saved request with result", func(t *testing.T) {
		result, err := repository.GetRequestWithResult(requestEntity.Id)

		a.Nil(err)
		a.NotNil(result)
		a.Equal(requestEntity.Id, result.RequestId)
		a.Equal(requestEntity.Year, result.Year)
		a.Equal(requestEntity.Day, result.Day)
		a.Equal(requestEntity.Part, result.Part)
		a.Equal(requestEntity.CreatedAt.Local(), result.CreatedAt.Local())
		a.Equal(resultEntity.StartedAt.Local(), result.StartedAt.Local())
		a.Equal(resultEntity.CompletedAt.Local(), result.CompletedAt.Local())
		a.Equal(resultEntity.Status, *result.Status)
		a.Equal(resultEntity.Result, result.Result)
		a.Equal(*requestEntity.S3Link, *result.S3Link)
	})

	t.Run("Should not find another request", func(t *testing.T) {
		_, err := repository.GetRequestWithResult(requestEntity.Id + 1)

		a.NotNil(err)
	})
}

func DeleteRequest(db *sqlx.DB, id int64) {
	db.MustExec(`delete from request where id = $1`, id)
}
