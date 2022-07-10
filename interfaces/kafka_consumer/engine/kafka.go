package engine

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/logger"
	"github.com/segmentio/kafka-go"
)

type EngineLib struct {
	log logger.Logger
	kfk kafka.ReaderConfig
}

func NewKafkaEngine(cfg config.Config, log logger.Logger) *EngineLib {
	return &EngineLib{
		log,
		kafka.ReaderConfig{
			Brokers: []string{fmt.Sprintf("%s:%d",
				cfg.KAFKA.KAFKA_SERVER,
				cfg.KAFKA.KAFKA_PORT,
			)},
			GroupID: cfg.KAFKA.KAFKA_CONSUMER_GROUP,
		},
	}
}

func (EL EngineLib) Consume(wg *sync.WaitGroup, topic string, handler func(context.Context, string, string) error) {
	defer wg.Done()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "topic", topic)
	EL.kfk.Topic = topic
	r := kafka.NewReader(EL.kfk)
	EL.log.InfoW("Kafka starting to consume", map[string]string{
		"topic": topic,
	})

	run := true
	go func() {
		for run {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				break
			}

			EL.log.InfoW("Incomming Message", map[string]string{
				"topic": topic,
				"key":   string(m.Key),
				"value": string(m.Value),
			})
			handler(ctx, string(m.Key), string(m.Value))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	run = false
	if err := r.Close(); err != nil {
		EL.log.FatalW("failed to close reader:", err)
	} else {
		EL.log.InfoW("success to close reader:", map[string]string{
			"topic": topic,
		})
	}
}
