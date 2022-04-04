package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Http HttpServer
	DB   DBServer
}

type HttpServer struct {
	Server string
	Port   string
}

type DBServer struct {
	DB_USER   string
	DB_PASS   string
	DB_SERVER string
	DB_PORT   string
	DB_NAME   string
}

func Configuration() Config {
	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	cfg.Http.Server = os.Getenv("SERVER")
	cfg.Http.Port = os.Getenv("PORT")

	cfg.DB.DB_USER = os.Getenv("DB_USER")
	cfg.DB.DB_PASS = os.Getenv("DB_PASS")
	cfg.DB.DB_SERVER = os.Getenv("DB_SERVER")
	cfg.DB.DB_PORT = os.Getenv("DB_PORT")
	cfg.DB.DB_NAME = os.Getenv("DB_NAME")

	return cfg
}
