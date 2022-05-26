package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/riskykurniawan15/xarch/logger"
)

var log logger.Logger = logger.NewLogger()

type Config struct {
	Http  HttpServer
	DB    DBServer
	RDB   RDBServer
	JWT   JWTConfig
	OTHER Other
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

type JWTConfig struct {
	SecretKey string
	Expired   int
}

type Other struct {
	AlQuranAPI string
}

func Configuration() Config {
	log.Info("Load all configuration")
	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		log.PanicW("Error loading .env file", err)
		panic("Error loading .env file")
	}

	cfg.Http.Server = os.Getenv("SERVER")
	cfg.Http.Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.PanicW("PORT must be number", err)
		panic("PORT must be number")
	}

	cfg.DB.DB_USER = os.Getenv("DB_USER")
	cfg.DB.DB_PASS = os.Getenv("DB_PASS")
	cfg.DB.DB_SERVER = os.Getenv("DB_SERVER")
	cfg.DB.DB_PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.PanicW("DB_PORT must be number", err)
		panic("DB_PORT must be number")
	}
	cfg.DB.DB_NAME = os.Getenv("DB_NAME")

	cfg.RDB.RDB_ADDRESS = os.Getenv("RDB_ADDRESS")
	cfg.RDB.RDB_PORT, err = strconv.Atoi(os.Getenv("RDB_PORT"))
	if err != nil {
		log.PanicW("RDB_PORT must be number", err)
		panic("RDB_PORT must be number")
	}
	cfg.RDB.RDB_PASS = os.Getenv("RDB_PASS")
	cfg.RDB.RDB_DB_DEFAULT, err = strconv.Atoi(os.Getenv("RDB_DB_DEFAULT"))
	if err != nil {
		log.PanicW("RDB_DB_DEFAULT must be number", err)
		panic("RDB_DB_DEFAULT must be number")
	}

	cfg.JWT.SecretKey = os.Getenv("JWT_SECRET_KEY")
	cfg.JWT.Expired, err = strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	if err != nil {
		log.PanicW("JWT_EXPIRED must be number", err)
		panic("JWT_EXPIRED must be number")
	}

	cfg.OTHER.AlQuranAPI = os.Getenv("ALQURAN_API")

	log.Info("Success for load all configuration")
	return cfg
}
