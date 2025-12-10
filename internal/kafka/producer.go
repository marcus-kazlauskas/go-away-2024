package kafka

import (
	"encoding/json"
	"go-away-2024/internal/config"

	"github.com/gofiber/fiber/v2/log"
	kafka "github.com/segmentio/kafka-go"
)

const KafkaProducerName string = "producer"

type KafkaProducer struct {
	p *kafka.Conn
}

func NewProducer(cfg *config.Config) *KafkaProducer {
	conn := newConn(cfg, KafkaProducerName)
	return &KafkaProducer{
		p: conn,
	}
}

func (k *KafkaProducer) SendTask(msg TaskMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = k.p.Write(data)
	if err != nil {
		return err
	}

	log.Infof("Successfully sended task message with id=%d", msg.Id)
	return nil
}
