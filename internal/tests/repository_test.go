package tests

import (
	"go-away-2024/internal/config"
	"go-away-2024/internal/database"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	a := assert.New(t)

	count := int32(0)
	requestIds := make([]int64, MaxInsertedRows)
	addRequestId := func(id int64) {
		requestIds[count] = id
		count++
	}
	lastRequestId := func() int64 {
		return requestIds[count-1]
	}

	config := config.GetConfig(config.TestPath)
	repository := database.NewRepository(database.Connect(config))
	deleteRequest := func() {
		for i := range count {
			repository.DeleteRequest(requestIds[i])
		}
	}

	defer deleteRequest()

	t.Run("Should find saved task", func(t *testing.T) {
		now := time.Now().UTC()
		request := Request(repository, 2024, 1, 1, now)
		addRequestId(request.Id)

		result, err := repository.GetRequestWithResult(lastRequestId())

		a.Nil(err)
		a.NotNil(result)
		a.Equal(request.Id, result.RequestId)
		a.Equal(request.Year, result.Year)
		a.Equal(request.Day, result.Day)
		a.Equal(request.Part, result.Part)
		a.Equal(request.CreatedAt, result.CreatedAt)
	})

	t.Run("Should not find another task", func(t *testing.T) {
		_, err := repository.GetRequestWithResult(lastRequestId() + 1)

		a.NotNil(err)
	})
}
