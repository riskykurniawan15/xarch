package repository

import (
	"context"
	"fmt"

	"github.com/riskykurniawan15/xarch/domain/users/models"
	"github.com/segmentio/kafka-go"
)

func (repo UserRepo) VerifiedEmailPublish(ctx context.Context, user *models.User) error {
	Address := fmt.Sprintf("%s:%d",
		repo.cfg.KAFKA.KAFKA_SERVER,
		repo.cfg.KAFKA.KAFKA_PORT,
	)

	w := &kafka.Writer{
		Addr:     kafka.TCP(Address),
		Topic:    repo.cfg.KAFKA.TOPIC_EMAIL_SENDER,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(user.Email + "_verified"),
			Value: []byte(user.Email),
		},
	)

	return err
}
