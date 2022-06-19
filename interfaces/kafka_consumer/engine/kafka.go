package engine

import (
	"context"
	"fmt"
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

func (EL EngineLib) Consume(topic string, handler func(context.Context, string, string) error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx := context.Background()
		EL.kfk.Topic = topic
		r := kafka.NewReader(EL.kfk)
		EL.log.InfoW("Kafka starting to consume", map[string]string{
			"topic": topic,
		})

		for {
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

		if err := r.Close(); err != nil {
			EL.log.FatalW("failed to close reader:", err)
		}
	}()
	wg.Wait()
}
