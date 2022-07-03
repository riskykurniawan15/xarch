package repository

import (
	"context"
	"net/http"

	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/riskykurniawan15/xarch/config"
	"gorm.io/gorm"
)

type HealthRepo struct {
	cfg config.Config
	DB  *gorm.DB
	RDB *redis.Client
}

func NewHealthRepo(cfg config.Config, DB *gorm.DB, RDB *redis.Client) *HealthRepo {
	return &HealthRepo{
		cfg,
		DB,
		RDB,
	}
}

func (repo HealthRepo) InternetHealth(ctx context.Context) bool {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true

}

func (repo HealthRepo) DatabaseHealth(ctx context.Context) (sql.DBStats, error) {
	sqlDB, _ := repo.DB.DB()

	return sqlDB.Stats(), sqlDB.Ping()
}

func (repo HealthRepo) RedisHealth(ctx context.Context) (interface{}, error) {
	result, err := repo.RDB.Do(ctx, "MEMORY", "STATS").Result()
	return result, err
}
