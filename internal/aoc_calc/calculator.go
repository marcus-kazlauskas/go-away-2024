package aoc_calc

import (
	"errors"
	"fmt"
	"go-away-2024/internal/api"
	"go-away-2024/internal/config"
	"go-away-2024/internal/database"
	"go-away-2024/internal/kafka"
	"go-away-2024/internal/minio"
	"go-away-2024/internal/puzzles"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
	kafka_go "github.com/segmentio/kafka-go"
)

type Calculator struct {
	repository      *database.Repository
	minioClient     *minio.MinioClient
	kafkaConnection *kafka.KafkaConnection
	sleep           time.Duration
}

func NewCalculator(
	repo *database.Repository,
	minio *minio.MinioClient,
	kafka *kafka.KafkaConnection,
	cfg *config.Config,
) *Calculator {
	sleep, _ := time.ParseDuration(cfg.Calculator.Sleep)

	log.Infof("Calculator created: sleep=%.1fs", sleep.Seconds())
	return &Calculator{
		repository:      repo,
		minioClient:     minio,
		kafkaConnection: kafka,
		sleep:           sleep,
	}
}

func (c *Calculator) Start() error {
	for {
		time.Sleep(c.sleep)

		// read new puzzle from kafka
		msg, err := c.kafkaConnection.ReadTask()
		if err != nil {
			if !errors.Is(err, kafka_go.RequestTimedOut) {
				return err
			}
			continue
		}

		// check if puzzle is already solved
		res, err := c.repository.GetResult(msg.Id)
		if err != nil {
			return err
		}
		if res.RequestId == msg.Id && res.Status == fmt.Sprint(api.COMPLITED) {
			log.Infof("Task id=%d is already solved!", msg.Id)
			continue
		}

		// solve puzzle
		startedAt := time.Now()
		ans, err := c.calculate(msg)
		completedAt := time.Now()

		// save result
		res.Result = ans
		res.StartedAt = &startedAt
		res.CompletedAt = &completedAt
		if err != nil {
			log.Infof("Cant't solve task id=%d: %v", msg.Id, err)
			errorResult := fmt.Sprintf("%v", err)
			res.Result = &errorResult
			res.Status = fmt.Sprint(api.ERROR)
		} else {
			res.Result = ans
			res.Status = fmt.Sprint(api.COMPLITED)
		}
		err = c.repository.SetResult(res)
		if err != nil {
			return err
		}
		log.Infof("Solved task id=%d result='%s' status=%s", res.RequestId, *res.Result, res.Status)
	}
}

func (c *Calculator) calculate(msg *kafka.TaskMessage) (*string, error) {
	tmpFile, err := os.CreateTemp("", *msg.S3Link)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	if err := c.minioClient.DownloadPuzzleInput(*msg.S3Link, tmpFile); err != nil {
		return nil, err
	}

	switch msg.Year {
	case 2024:
		switch msg.Day {
		case 1:
			switch msg.Part {
			case 1:
				return puzzles.Year2024Day1Part1(tmpFile)
			default:
				return nil, fmt.Errorf("puzzle year=%d day=%d part=%d is not supported", msg.Year, msg.Day, msg.Part)
			}
		default:
			return nil, fmt.Errorf("puzzle year=%d day=%d is not supported", msg.Year, msg.Day)
		}
	// TODO: create year 2025 implementation
	default:
		return nil, fmt.Errorf("puzzle year=%d is not supported", msg.Year)
	}
}
