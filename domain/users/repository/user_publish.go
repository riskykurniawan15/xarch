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
		Topic:    repo.cfg.KAFKA.TOPIC_EMAIL_VERIFIED,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("verification_" + fmt.Sprint(user.ID)),
			Value: []byte(user.Email),
		},
	)

	return err
}

func (repo UserRepo) ForgotPassByEmailPublish(ctx context.Context, user *models.User) error {
	Address := fmt.Sprintf("%s:%d",
		repo.cfg.KAFKA.KAFKA_SERVER,
		repo.cfg.KAFKA.KAFKA_PORT,
	)

	w := &kafka.Writer{
		Addr:     kafka.TCP(Address),
		Topic:    repo.cfg.KAFKA.TOPIC_PASS_FORGOT,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("forgot_" + fmt.Sprint(user.ID)),
			Value: []byte(user.Email),
		},
	)

	return err
}
