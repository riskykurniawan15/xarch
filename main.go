package main

import (
	"github.com/riskykurniawan15/xarch/exec/arch"
	"github.com/riskykurniawan15/xarch/logger"
)

func main() {
	logger := logger.NewLogger()
	logger.Info("Starting Application")
	arch.Startup()
}
