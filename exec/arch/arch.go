package arch

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/domain"
	"github.com/riskykurniawan15/xarch/driver"
	"github.com/riskykurniawan15/xarch/exec/elsa"
	echo "github.com/riskykurniawan15/xarch/interfaces/http/engine"
	"github.com/riskykurniawan15/xarch/logger"
)

var log logger.Logger = logger.NewLogger()

func Startup() {
	log.Info("Running startup")
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

func StartDriver(cfg config.Config) (*gorm.DB, *redis.Client) {
	log.Info("Start all driver")
	db := driver.ConnectDB(cfg.DB)
	rdb := driver.ConnectRedis(cfg.RDB)

	return db, rdb
}

func EngineSwitch() {
	log.Info("Running Switch Engine")
	engine_run := flag.String("engine", "*", "type your egine")
	flag.Parse()

	cfg := config.Configuration()
	DB, _ := StartDriver(cfg)
	svc := domain.StartService(DB)

	switch *engine_run {
	case "http":
		log.Info("Starting Http Engine")
		echo.Start(cfg, svc)
	case "*":
		log.Error("Failed run engine")
		fmt.Println("Please type your engine")
		fmt.Println("go run main.go -engine=type")
	default:
		fmt.Println("Failed run engine")
		log.Error("Failed run engine")
	}
}
