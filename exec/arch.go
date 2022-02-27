package exec

import (
	"flag"
	"fmt"
	"time"

	"github.com/riskykurniawan15/xarch/config"
	echo "github.com/riskykurniawan15/xarch/interfaces/http/engine"
)

func Startup() {
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
	fmt.Println("By: Risky Kurniawan | https://risoftinc.com")
	fmt.Println("Please wait, your program is running in 5 seconds", string("\033[0m"))
	time.Sleep(time.Duration(5) * time.Second)
}

func EngineSwitch() {
	engine_run := flag.String("engine", "*", "type your egine")
	flag.Parse()

	cfg := config.Configuration()

	switch *engine_run {
	case "http":
		echo.Start(cfg)
	case "*":
		fmt.Println("Please type your engine")
		fmt.Println("go run main.go -engine=type")
	default:
		fmt.Println("Failed run engine")
	}
}
