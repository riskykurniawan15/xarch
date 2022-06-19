package kafka_consumer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain"
	"github.com/riskykurniawan15/xarch/interfaces/kafka_consumer/handlers"
	"github.com/riskykurniawan15/xarch/logger"
)

type Handlers struct {
	EmailSenderHandler *handlers.EmailSenderHandler
}

func StartHandlers(svc *domain.Service) *Handlers {
	email_sender := handlers.NewEmailSenderHandler(svc.UserService)

	return &Handlers{
		email_sender,
	}
}

func ConsumerRun(cfg config.Config, log logger.Logger, svc *domain.Service) {
	handler := StartHandlers(svc)

	Address1 := fmt.Sprintf("%s:%d",
		cfg.KAFKA.KAFKA_SERVER,
		cfg.KAFKA.KAFKA_PORT,
	)

	kafka_config := kafka.ReaderConfig{
		Brokers: []string{Address1},
		GroupID: cfg.KAFKA.KAFKA_CONSUMER_GROUP,
	}

	Consume(log, kafka_config, cfg.KAFKA.TOPIC_EMAIL_SENDER, handler.EmailSenderHandler.SendVerification)
}

func Consume(log logger.Logger, kafka_config kafka.ReaderConfig, topic string, handler func(context.Context, string, string) error) {
	go func() {
		ctx := context.Background()
		kafka_config.Topic = topic
		r := kafka.NewReader(kafka_config)
		log.InfoW("Kafka consume running", map[string]string{
			"topic": topic,
		})

		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				break
			}

			log.InfoW("Incomming Message", map[string]string{
				"topic": topic,
				"key":   string(m.Key),
				"value": string(m.Value),
			})
			handler(ctx, string(m.Key), string(m.Value))
		}

		if err := r.Close(); err != nil {
			log.FatalW("failed to close reader:", err)
		}
	}()
}
