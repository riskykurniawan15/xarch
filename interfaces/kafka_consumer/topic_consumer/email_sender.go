package topic_consumer

import (
	"context"
	"strings"

	"github.com/segmentio/kafka-go"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/interfaces/kafka_consumer/handlers"
	"github.com/riskykurniawan15/xarch/logger"
)

type ConsumerEmailSender struct {
	handler handlers.EmailSenderHandler
	cfg     config.Config
	log     logger.Logger
}

func NewConsumerEmailSender(handler handlers.EmailSenderHandler, cfg config.Config, log logger.Logger) *ConsumerEmailSender {
	return &ConsumerEmailSender{
		handler,
		cfg,
		log,
	}
}

func (C ConsumerEmailSender) Consume(ctx context.Context, kafka_config kafka.ReaderConfig) {
	kafka_config.Topic = C.cfg.KAFKA.TOPIC_EMAIL_SENDER
	r := kafka.NewReader(kafka_config)
	C.log.Info("Kafka consume running")

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}

		route := strings.Split(string(m.Key), "_")

		if route[0] == "verification" {
			C.log.InfoW("incomming request to send verification", map[string]interface{}{
				"key":   string(m.Key),
				"value": string(m.Value),
			})
			C.handler.SendVerification(ctx, string(m.Key), string(m.Value))
		} else {
			C.log.Error("method not exists")
		}

	}

	if err := r.Close(); err != nil {
		C.log.FatalW("failed to close reader:", err)
	}
}
