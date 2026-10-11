package kafka

import (
	"go-away-2024/internal/config"
	"net"
)

type TaskMessage struct {
	Id     int64   `json:"id"`
	Year   int32   `json:"year"`
	Day    int32   `json:"day"`
	Part   int32   `json:"part"`
	S3Link *string `json:"s3_link"`
}

func address(cfg *config.Config) string {
	return net.JoinHostPort(cfg.Kafka.Host, cfg.Kafka.Port)
}
