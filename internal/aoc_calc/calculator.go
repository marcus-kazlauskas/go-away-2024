package aoc_calc

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-away-2024/internal/api"
	"go-away-2024/internal/database"
	"go-away-2024/internal/kafka"
	"go-away-2024/internal/minio"
	"go-away-2024/internal/puzzles"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

const fetchRetry = time.Second

type Calculator struct {
	repository  *database.Repository
	minioClient *minio.MinioClient
	kafkaReader *kafka.TaskReader
}

func NewCalculator(
	repo *database.Repository,
	minio *minio.MinioClient,
	kafka *kafka.TaskReader,
) *Calculator {
	log.Info("Calculator created")
	return &Calculator{
		repository:  repo,
		minioClient: minio,
		kafkaReader: kafka,
	}
}

// Start solves tasks until ctx is cancelled. The current task is always finished,
// saved and committed before returning.
func (c *Calculator) Start(ctx context.Context) error {
	for {
		// wait for new puzzle from kafka
		task, msg, err := c.kafkaReader.FetchTask(ctx)
		// a task fetched at the moment of cancellation is not committed and is read again after restart
		select {
		case <-ctx.Done():
			log.Info("Calculator stopped")
			return nil
		default:
		}

		// the reader reconnects by itself, so kafka errors are transient
		if err != nil {
			log.Errorf("Couldn't read task, retry in %.0fs: %v", fetchRetry.Seconds(), err)
			select {
			case <-ctx.Done():
			case <-time.After(fetchRetry):
			}
			continue
		}

		if err := c.solve(task); err != nil {
			return err
		}

		// commit only after the result is saved, so a crash means solving it again, not losing it;
		// an uncommitted task is read again and skipped as already solved
		if err := c.kafkaReader.Commit(context.WithoutCancel(ctx), msg); err != nil {
			log.Errorf("Couldn't commit task id=%d: %v", task.Id, err)
		}
	}
}

func (c *Calculator) solve(msg *kafka.TaskMessage) error {
	// check if puzzle is already solved
	res, err := c.repository.GetResult(msg.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Infof("Task id=%d is unknown", msg.Id)
			return nil
		}
		return err
	}
	if res.Status == fmt.Sprint(api.COMPLETED) || res.Status == fmt.Sprint(api.ERROR) {
		log.Infof("Task id=%d is already solved", msg.Id)
		return nil
	}

	// solve puzzle
	startedAt := time.Now()
	ans, err := c.calculate(msg)
	completedAt := time.Now()

	// save result
	res.StartedAt = &startedAt
	res.CompletedAt = &completedAt
	if err != nil {
		log.Infof("Couldn't solve task id=%d: %v", msg.Id, err)
		errorResult := fmt.Sprintf("%v", err)
		res.Result = &errorResult
		res.Status = fmt.Sprint(api.ERROR)
	} else {
		res.Result = ans
		res.Status = fmt.Sprint(api.COMPLETED)
	}
	err = c.repository.SetResult(res)
	if err != nil {
		return err
	}
	log.Infof("Solved task id=%d result='%s' status=%s", res.RequestId, *res.Result, res.Status)
	return nil
}

func (c *Calculator) calculate(msg *kafka.TaskMessage) (*string, error) {
	tmpFile, err := os.CreateTemp("", *msg.S3Link)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := os.Remove(tmpFile.Name()); err != nil {
			log.Error(err)
		}
	}()

	if err := c.minioClient.DownloadPuzzleInput(*msg.S3Link, tmpFile); err != nil {
		return nil, err
	}

	if _, err = tmpFile.Seek(0, 0); err != nil {
		log.Error(err)
	}

	scan := bufio.NewScanner(tmpFile)

	switch msg.Year {
	case 2024:
		switch msg.Day {
		case 1:
			switch msg.Part {
			case 1:
				return puzzles.Year2024Day1Part1(scan)
			case 2:
				return puzzles.Year2024Day1Part2(scan)
			default:
				return nil, PartError(msg.Year, msg.Day, msg.Part)
			}
		default:
			return nil, DayError(msg.Year, msg.Day)
		}
	case 2025:
		switch msg.Day {
		case 1:
			switch msg.Part {
			case 1:
				return puzzles.Year2025Day1Part1(scan)
			case 2:
				return puzzles.Year2025Day1Part2(scan)
			default:
				return nil, PartError(msg.Year, msg.Day, msg.Part)
			}
		case 2:
			switch msg.Part {
			case 1:
				return puzzles.Year2025Day2Part1(scan)
			case 2:
				return puzzles.Year2025Day2Part2(scan)
			default:
				return nil, PartError(msg.Year, msg.Day, msg.Part)
			}
		case 3:
			switch msg.Part {
			case 1:
				return puzzles.Year2025Day3Part1(scan)
			case 2:
				return puzzles.Year2025Day3Part2(scan)
			default:
				return nil, PartError(msg.Year, msg.Day, msg.Part)
			}
		case 4:
			switch msg.Part {
			case 1:
				return puzzles.Year2025Day4Part1(scan)
			case 2:
				return puzzles.Year2025Day4Part2(scan)
			default:
				return nil, PartError(msg.Year, msg.Day, msg.Part)
			}
		case 5:
			switch msg.Part {
			case 1:
				return puzzles.Year2025Day5Part1(scan)
			case 2:
				return puzzles.Year2025Day5Part2(scan)
			default:
				return nil, PartError(msg.Year, msg.Day, msg.Part)
			}
		case 6:
			switch msg.Part {
			case 1:
				return puzzles.Year2025Day6Part1(scan)
			case 2:
				return puzzles.Year2025Day6Part2(scan)
			default:
				return nil, PartError(msg.Year, msg.Day, msg.Part)
			}
		default:
			return nil, DayError(msg.Year, msg.Day)
		}
	default:
		return nil, YearError(msg.Year)
	}
}

func YearError(year int32) error {
	return fmt.Errorf("puzzle year=%d is not supported", year)
}

func DayError(year int32, day int32) error {
	return fmt.Errorf("puzzle year=%d day=%d is not supported", year, day)
}

func PartError(year int32, day int32, part int32) error {
	return fmt.Errorf("puzzle year=%d day=%d part=%d is not supported", year, day, part)
}
