package services

import (
	"context"

	"github.com/riskykurniawan15/xarch/domain/health/models"
	"github.com/riskykurniawan15/xarch/domain/health/repository"
)

type HealthService struct {
	HealthRepo *repository.HealthRepo
}

func NewHealthService(Repo *repository.HealthRepo) *HealthService {
	return &HealthService{
		Repo,
	}
}

func (svc HealthService) HealthMetric(ctx context.Context) *models.HealthMetric {
	var metric models.HealthMetric
	status := map[string]interface{}{
		"internet": "connected",
		"database": "connected",
		"redis":    "connected",
	}

	if !svc.HealthRepo.InternetHealth(ctx) {
		status["internet"] = "refused"
	}

	databaseHealth, err := svc.HealthRepo.DatabaseHealth(ctx)
	if err != nil {
		status["database"] = "refused"
		metric.DB = err.Error()
	} else {
		metric.DB = databaseHealth
	}
	redisHealth, err := svc.HealthRepo.RedisHealth(ctx)
	if err != nil {
		status["redis"] = "refused"
		metric.RDB = err.Error()
	} else {
		metric.RDB = redisHealth
	}
	metric.Status = status

	return &metric
}
