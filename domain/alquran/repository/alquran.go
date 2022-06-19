package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/riskykurniawan15/xarch/config"
)

type AlquranRepo struct {
	cfg config.Config
	RDB *redis.Client
}

func NewAlquranRepo(cfg config.Config, RDB *redis.Client) *AlquranRepo {
	return &AlquranRepo{
		cfg,
		RDB,
	}
}
