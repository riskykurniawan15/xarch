package arch

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain"
	"github.com/riskykurniawan15/xarch/driver"
	"github.com/riskykurniawan15/xarch/exec/elsa"
	echo "github.com/riskykurniawan15/xarch/interfaces/http/engine"
	consumer "github.com/riskykurniawan15/xarch/interfaces/kafka_consumer"
	"github.com/riskykurniawan15/xarch/logger"
)

var (
	cfg config.Config = config.Configuration()
	log logger.Logger = logger.NewLogger()
)

func Startup() {
	var cli string = ""
	if len(os.Args) > 1 {
		cli = strings.ToLower(os.Args[1])
	}

	switch cli {
	case "elsa":
		elsa.RunElsa()
	default:
		RunApp()
	}
}

func RunApp() {
	log.Info("Running Application Logo")
	logo := `Welcome to:
     ___           ___           ___           ___           ___     
    |\__\         /\  \         /\  \         /\  \         /\__\    
    |:|  |       /::\  \       /::\  \       /::\  \       /:/  /    
    |:|  |      /:/\:\  \     /:/\:\  \     /:/\:\  \     /:/__/     
    |:|__|__   /::\~\:\  \   /::\~\:\  \   /:/  \:\  \   /::\  \ ___ 
____/::::\__\ /:/\:\ \:\__\ /:/\:\ \:\__\ /:/__/ \:\__\ /:/\:\  /\__\
\::::/~~/~    \/__\:\/:/  / \/_|::\/:/  / \:\  \  \/__/ \/__\:\/:/  /
 ~~|:|~~|          \::/  /     |:|::/  /   \:\  \            \::/  / 
   |:|  |          /:/  /      |:|\/__/     \:\  \           /:/  /  
   |:|  |         /:/  /       |:|  |        \:\__\         /:/  /   
    \|__|         \/__/         \|__|         \/__/         \/__/    V 1.0
	`

	fmt.Print(string("\033[32m"))
	fmt.Println(logo, string("\033[34m"))
	fmt.Println("By: Risky Kurniawan | https://risoftinc.com | mailto:riskykurniawan@risoftinc.com")
	fmt.Println("Please wait, your program is running in 5 seconds", string("\033[0m"))
	time.Sleep(time.Duration(5) * time.Second)
	EngineSwitch()
}

func StartDriver() (*gorm.DB, *redis.Client) {
	log.Info("Start all driver")
	db := driver.ConnectDB(cfg.DB)
	rdb := driver.ConnectRedis(cfg.RDB)

	return db, rdb
}

func EngineSwitch() {
	var wg sync.WaitGroup
	log.Info("Running Switch Engine")
	engine_run := flag.String("engine", "*", "type your egine")
	flag.Parse()

	switch *engine_run {
	case "http":
		DB, RDB := StartDriver()
		svc := domain.StartService(cfg, DB, RDB)
		log.Info("Starting Http Engine")
		echo.Start(cfg, svc)
	case "consumer":
		DB, RDB := StartDriver()
		svc := domain.StartService(cfg, DB, RDB)
		log.Info("Starting Consumer Engine")
		consumer.ConsumerRun(cfg, log, svc)
	case "*":
		DB, RDB := StartDriver()
		svc := domain.StartService(cfg, DB, RDB)
		log.Info("Starting All Engine")
		wg.Add(1)
		go func() { defer wg.Done(); echo.Start(cfg, svc) }()
		wg.Add(1)
		go func() { defer wg.Done(); consumer.ConsumerRun(cfg, log, svc) }()
	default:
		log.Error("Failed run engine")
	}

	wg.Wait()
}
