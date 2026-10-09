package kafka

import (
	"context"
	"encoding/json"
	"go-away-2024/internal/config"
	"time"

	"github.com/gofiber/fiber/v2/log"
	kafka "github.com/segmentio/kafka-go"
)

type TaskWriter struct {
	w            *kafka.Writer
	writeTimeout time.Duration
}

func NewTaskWriter(cfg *config.Config) *TaskWriter {
	batchTimeout, err := time.ParseDuration(cfg.Kafka.BatchTimeout)
	if err != nil {
		config.Fatal(err)
	}
	writeTimeout, err := time.ParseDuration(cfg.Kafka.WriteTimeout)
	if err != nil {
		config.Fatal(err)
	}

	w := &kafka.Writer{
		Addr:                   kafka.TCP(address(cfg)),
		Topic:                  cfg.Kafka.Topic,
		RequiredAcks:           kafka.RequireAll,
		BatchTimeout:           batchTimeout,
		AllowAutoTopicCreation: true,
		ErrorLogger:            kafka.LoggerFunc(log.Errorf),
	}

	log.Infof("Kafka writer created: address=%s, topic=%s", address(cfg), cfg.Kafka.Topic)
	return &TaskWriter{
		w:            w,
		writeTimeout: writeTimeout,
	}
}

func (k *TaskWriter) Close() error {
	return k.w.Close()
}

func (k *TaskWriter) WriteTask(ctx context.Context, msg *TaskMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, k.writeTimeout)
	defer cancel()
	if err := k.w.WriteMessages(ctx, kafka.Message{Value: data}); err != nil {
		return err
	}

	log.Infof("Wrote task message with id=%d", msg.Id)
	return nil
}
