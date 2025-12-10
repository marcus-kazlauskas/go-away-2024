package kafka

import (
	"context"
	"go-away-2024/internal/config"
	"net"

	"github.com/gofiber/fiber/v2/log"
	kafka "github.com/segmentio/kafka-go"
)

type TaskMessage struct {
	Id     int64   `json:"id"`
	Year   int32   `json:"year"`
	Day    int32   `json:"day"`
	Part   int32   `json:"part"`
	S3Link *string `json:"s3_link"`
}

func newConn(cfg *config.Config, name string) *kafka.Conn {
	address := net.JoinHostPort(cfg.Kafka.Host, cfg.Kafka.Port)
	conn, err := kafka.DialLeader(context.Background(), cfg.Kafka.Network, address, cfg.Kafka.Topic, cfg.Kafka.Partition)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Connection to kafka created: name=%s, address=%s", name, address)
	return conn
}
