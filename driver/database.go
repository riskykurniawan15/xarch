package driver

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/logger"
)

var log logger.Logger = logger.NewLogger()

func ConnectDB(cfg config.DBServer) *gorm.DB {
	log.Info("Connection to database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USER,
		cfg.DB_PASS,
		cfg.DB_SERVER,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.PanicW("Failed to Connect DB", err)
		panic("Failed to Connect DB")
	}

	log.Info("Database connected")

	return db
}

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
