package app

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/riskykurniawan15/xarch/app/elsa"
	"github.com/riskykurniawan15/xarch/app/xarch"
)

func RunApp() {
	version := "1.0.0"
	duration := 1
	logo := fmt.Sprintf(`Welcome to:
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
    \|__|         \/__/         \|__|         \/__/         \/__/    V %s
	`, version)

	fmt.Print(string("\033[32m"))
	fmt.Println(logo, string("\033[34m"))
	fmt.Println("By: Risky Kurniawan | https://risoftinc.com | mailto:riskykurniawan@risoftinc.com")
	fmt.Println("Please wait, your program is running in", duration, "seconds", string("\033[0m"))
	time.Sleep(time.Duration(duration) * time.Second)
	Startup()
}

func Startup() {
	var app_instance string = ""
	if len(os.Args) > 1 {
		app_instance = strings.ToLower(os.Args[1])
	}

	switch app_instance {
	case "elsa":
		elsa.RunElsa()
	case "xarch":
		xarch.RunXarch()
	default:
		fmt.Println("please run app with app instance [elsa|xarch], example:")
		fmt.Println("go run main.go ${app_instance}")
	}
}
