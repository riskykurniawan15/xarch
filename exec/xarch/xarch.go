package xarch

import (
	"flag"
	"sync"

	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain"
	"github.com/riskykurniawan15/xarch/driver"
	echo "github.com/riskykurniawan15/xarch/interfaces/http/engine"
	consumer "github.com/riskykurniawan15/xarch/interfaces/kafka_consumer"
	"github.com/riskykurniawan15/xarch/logger"
)

var (
	cfg config.Config = config.Configuration()
	log logger.Logger = logger.NewLogger()
)

func StartDriver() (*gorm.DB, *redis.Client) {
	log.Info("Start all driver")
	db := driver.ConnectDB(cfg.DB)
	rdb := driver.ConnectRedis(cfg.RDB)

	return db, rdb
}

func ShutdownDriver(DB *gorm.DB, RDB *redis.Client) {
	dbCon, _ := DB.DB()
	if err := dbCon.Close(); err != nil {
		log.ErrorW("failed to close database connection", err)
	} else {
		log.Info("success to close database connection")
	}

	if err := RDB.Close(); err != nil {
		log.ErrorW("failed to close redis connection", err)
	} else {
		log.Info("success to close redis connection")
	}
}

func RunXarch() {
	DB, RDB := StartDriver()
	var wg sync.WaitGroup
	svc := domain.StartService(cfg, DB, RDB)

	log.Info("Running Switch Engine")
	engine := flag.String("engine", "*", "type your egine")
	flag.Parse()

	switch *engine {
	case "*":
		log.Info("Starting All Engine")
		wg.Add(1)
		go consumer.ConsumerRun(&wg, cfg, log, svc)
		wg.Add(1)
		go echo.Start(&wg, cfg, log, svc)
	case "http":
		log.Info("Starting Http Engine")
		wg.Add(1)
		go echo.Start(&wg, cfg, log, svc)
	case "consumer":
		log.Info("Starting Consumer Engine")
		wg.Add(1)
		go consumer.ConsumerRun(&wg, cfg, log, svc)
	default:
		log.Error("Failed run engine")
	}

	wg.Wait()

	ShutdownDriver(DB, RDB)
	log.Info("application is shutdown")
}
