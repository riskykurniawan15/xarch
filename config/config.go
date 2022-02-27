package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Http HttpServer
}

type HttpServer struct {
	Server string
	Port   string
}

func Configuration() Config {
	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	cfg.Http.Server = os.Getenv("SERVER")
	cfg.Http.Port = os.Getenv("PORT")

	return cfg
}
