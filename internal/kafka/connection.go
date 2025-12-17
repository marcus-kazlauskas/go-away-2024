package kafka

import (
	"encoding/json"
	"go-away-2024/internal/config"
	"time"

	"github.com/gofiber/fiber/v2/log"
	kafka "github.com/segmentio/kafka-go"
)

type KafkaConnection struct {
	conn          *kafka.Conn
	writeDeadline time.Duration
	readDeadLine  time.Duration
	maxBytes      int
}

func NewKafkaConnection(cfg *config.Config) *KafkaConnection {
	conn := Connect(cfg)

	writeDeadline, err := time.ParseDuration(cfg.Kafka.WriteDeadline)
	if err != nil {
		config.Fatal(err)
	}
	readDeadLine, err := time.ParseDuration(cfg.Kafka.ReadDeadLine)
	if err != nil {
		config.Fatal(err)
	}

	return &KafkaConnection{
		conn:          conn,
		writeDeadline: writeDeadline,
		readDeadLine:  readDeadLine,
		maxBytes:      cfg.Kafka.ReadBatchMaxSize,
	}
}

func (k *KafkaConnection) WriteTask(msg *TaskMessage) error {
	k.conn.SetWriteDeadline(time.Now().Add(k.writeDeadline))

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = k.conn.Write(data)
	if err != nil {
		return err
	}

	log.Infof("Wrote task message with id=%d, writeDeadline=%.1fs", msg.Id, k.writeDeadline.Seconds())
	return nil
}

func (k *KafkaConnection) ReadTask() (*TaskMessage, error) {
	k.conn.SetReadDeadline(time.Now().Add(k.readDeadLine))

	msg, err := k.conn.ReadMessage(k.maxBytes)
	if err != nil {
		return nil, err
	}

	task := &TaskMessage{}
	err = json.Unmarshal(msg.Value, task)
	if err != nil {
		return nil, err
	}

	log.Infof("Read task message with id=%d, readDeadLine=%.1fs", task.Id, k.readDeadLine.Seconds())
	return task, nil
}
