package kafka_consumer

import (
	"context"
	"fmt"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain"
	"github.com/riskykurniawan15/xarch/logger"
	"github.com/segmentio/kafka-go"
)

func ConsumerRun(cfg config.Config, log logger.Logger, svc *domain.Service) {
	ctx := context.Background()
	Address := fmt.Sprintf("%s:%d",
		cfg.KAFKA.KAFKA_SERVER,
		cfg.KAFKA.KAFKA_PORT,
	)

	kafka_config := kafka.ReaderConfig{
		Brokers:  []string{Address},
		GroupID:  cfg.KAFKA.KAFKA_CONSUMER_GROUP,
		Topic:    cfg.KAFKA.TOPIC_EMAIL_SENDER,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}

	r := kafka.NewReader(kafka_config)

	log.Info("Kafka consume running")

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}

		msg := fmt.Sprintf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		log.Info(msg)
	}

	if err := r.Close(); err != nil {
		log.FatalW("failed to close reader:", err)
	}
}
