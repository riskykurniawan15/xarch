package kafka_consumer

import (
	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain"
	"github.com/riskykurniawan15/xarch/interfaces/kafka_consumer/engine"
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
	en := engine.NewKafkaEngine(cfg, log)

	// to consume message
	en.Consume(cfg.KAFKA.TOPIC_EMAIL_SENDER, handler.EmailSenderHandler.SendVerification)
}
