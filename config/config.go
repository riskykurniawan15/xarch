package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Http HttpServer
	DB   DBServer
	RDB  RDBServer
}

type HttpServer struct {
	Server string
	Port   int
}

type DBServer struct {
	DB_USER   string
	DB_PASS   string
	DB_SERVER string
	DB_PORT   int
	DB_NAME   string
}

type RDBServer struct {
	RDB_ADDRESS    string
	RDB_PORT       int
	RDB_PASS       string
	RDB_DB_DEFAULT int
}

func Configuration() Config {
	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	cfg.Http.Server = os.Getenv("SERVER")
	cfg.Http.Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic("PORT must be number")
	}

	cfg.DB.DB_USER = os.Getenv("DB_USER")
	cfg.DB.DB_PASS = os.Getenv("DB_PASS")
	cfg.DB.DB_SERVER = os.Getenv("DB_SERVER")
	cfg.DB.DB_PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("DB_PORT must be number")
	}
	cfg.DB.DB_NAME = os.Getenv("DB_NAME")

	cfg.RDB.RDB_ADDRESS = os.Getenv("RDB_ADDRESS")
	cfg.RDB.RDB_PORT, err = strconv.Atoi(os.Getenv("RDB_PORT"))
	if err != nil {
		panic("RDB_PORT must be number")
	}
	cfg.RDB.RDB_PASS = os.Getenv("RDB_PASS")
	cfg.RDB.RDB_DB_DEFAULT, err = strconv.Atoi(os.Getenv("RDB_DB_DEFAULT"))
	if err != nil {
		panic("RDB_DB_DEFAULT must be number")
	}

	return cfg
}
