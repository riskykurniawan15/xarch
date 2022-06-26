package driver

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/riskykurniawan15/xarch/config"
)

func ConnectRedis(cfg config.RDBServer) *redis.Client {
	log.Info("Connection to redis")
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RDB_ADDRESS, cfg.RDB_PORT),
		Username: cfg.RDB_USER,
		Password: cfg.RDB_PASS,
		DB:       cfg.RDB_DB_DEFAULT,
	})

	log.Info("Redis connected")

	return rdb
}
