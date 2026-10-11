package kafka

import (
	"context"
	"encoding/json"
	"go-away-2024/internal/config"
	"time"

	"github.com/gofiber/fiber/v2/log"
	kafka "github.com/segmentio/kafka-go"
)

type TaskReader struct {
	r *kafka.Reader
}

func NewTaskReader(cfg *config.Config) *TaskReader {
	maxWait, err := time.ParseDuration(cfg.Kafka.MaxWait)
	if err != nil {
		config.Fatal(err)
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{address(cfg)},
		Topic:       cfg.Kafka.Topic,
		GroupID:     cfg.Kafka.GroupId,
		MinBytes:    1,
		MaxBytes:    cfg.Kafka.ReadBatchMaxSize,
		MaxWait:     maxWait,
		ErrorLogger: kafka.LoggerFunc(log.Errorf),
		// after a broker restart the reader rejoins in ~1s instead of ~15s with defaults (30s, 30s, 5s);
		// Close waits for a pending join, so shutdown is bounded by RebalanceTimeout
		SessionTimeout:   10 * time.Second,
		RebalanceTimeout: 10 * time.Second,
		JoinGroupBackoff: time.Second,
	})

	log.Infof("Kafka reader created: address=%s, topic=%s, group=%s", address(cfg), cfg.Kafka.Topic, cfg.Kafka.GroupId)
	return &TaskReader{
		r: r,
	}
}

func (k *TaskReader) Close() error {
	return k.r.Close()
}

func (k *TaskReader) FetchTask(ctx context.Context) (*TaskMessage, kafka.Message, error) {
	for {
		msg, err := k.r.FetchMessage(ctx)
		if err != nil {
			return nil, msg, err
		}

		task := &TaskMessage{}
		if err := json.Unmarshal(msg.Value, task); err != nil {
			log.Errorf("Skipped invalid task message at offset=%d: %v", msg.Offset, err)
			if err := k.Commit(ctx, msg); err != nil {
				return nil, msg, err
			}
			continue
		}

		log.Infof("Read task message with id=%d, offset=%d", task.Id, msg.Offset)
		return task, msg, nil
	}
}

func (k *TaskReader) Commit(ctx context.Context, msg kafka.Message) error {
	return k.r.CommitMessages(ctx, msg)
}
