package arch

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/riskykurniawan15/xarch/config"
	gate "github.com/riskykurniawan15/xarch/domain"
	"github.com/riskykurniawan15/xarch/driver"
	"github.com/riskykurniawan15/xarch/exec/elsa"
	echo "github.com/riskykurniawan15/xarch/interfaces/http/engine"
	"gorm.io/gorm"
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
	logo := fmt.Sprint(`Welcome to:
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
	`)

	fmt.Print(string("\033[32m"))
	fmt.Println(logo, string("\033[34m"))
	fmt.Println("By: Risky Kurniawan | https://risoftinc.com | mailto:riskykurniawan@risoftinc.com")
	fmt.Println("Please wait, your program is running in 5 seconds", string("\033[0m"))
	time.Sleep(time.Duration(5) * time.Second)
	EngineSwitch()
}

func StartDriver(cfg config.Config) (*gorm.DB, *redis.Client) {
	db := driver.ConnectDB(cfg.DB)
	rdb := driver.ConnectRedis(cfg.RDB)

	return db, rdb
}

func EngineSwitch() {
	engine_run := flag.String("engine", "*", "type your egine")
	flag.Parse()

	cfg := config.Configuration()
	DB, _ := StartDriver(cfg)
	svc := gate.StartService(DB)

	switch *engine_run {
	case "http":
		echo.Start(cfg, svc)
	case "*":
		fmt.Println("Please type your engine")
		fmt.Println("go run main.go -engine=type")
	default:
		fmt.Println("Failed run engine")
	}
}
