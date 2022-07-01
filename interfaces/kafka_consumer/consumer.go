package kafka_consumer

import (
	"sync"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain"
	"github.com/riskykurniawan15/xarch/interfaces/kafka_consumer/engine"
	"github.com/riskykurniawan15/xarch/interfaces/kafka_consumer/handlers"
	"github.com/riskykurniawan15/xarch/logger"
)

type Handlers struct {
	EmailSenderHandler *handlers.EmailSenderHandler
}

func StartHandlers(log logger.Logger, svc *domain.Service) *Handlers {
	email_sender := handlers.NewEmailSenderHandler(log, svc.UserService)

	return &Handlers{
		email_sender,
	}
}

func ConsumerRun(wg *sync.WaitGroup, cfg config.Config, log logger.Logger, svc *domain.Service) {
	defer wg.Done()
	var wgc sync.WaitGroup
	handler := StartHandlers(log, svc)
	email_sender := handler.EmailSenderHandler
	en := engine.NewKafkaEngine(cfg, log)

	wgc.Add(2)
	// to consume message
	go en.Consume(&wgc, cfg.KAFKA.TOPIC_EMAIL_VERIFIED, email_sender.SendVerification)
	go en.Consume(&wgc, cfg.KAFKA.TOPIC_EMAIL_VERIFIED, email_sender.SendVerification)
	wgc.Wait()
}
