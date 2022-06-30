package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/riskykurniawan15/xarch/logger"
)

var log logger.Logger = logger.NewLogger()

type Config struct {
	Http  HttpServer
	DB    DBServer
	RDB   RDBServer
	Email EmailSender
	KAFKA KafkaConfig
	JWT   JWTConfig
	OTHER Other
}

type HttpServer struct {
	Server string
	Port   int
	URL    string
}

type DBServer struct {
	DB_DRIVER        string
	DB_USER          string
	DB_PASS          string
	DB_SERVER        string
	DB_PORT          int
	DB_NAME          string
	DB_MAX_IDLE_CON  int
	DB_MAX_OPEN_CON  int
	DB_MAX_LIFE_TIME int
}

type RDBServer struct {
	RDB_ADDRESS    string
	RDB_PORT       int
	RDB_USER       string
	RDB_PASS       string
	RDB_DB_DEFAULT int
}

type EmailSender struct {
	EMAIL_HOST     string
	EMAIL_PORT     int
	EMAIL_NAME     string
	EMAIL_EMAIL    string
	EMAIL_PASSWORD string
}

type KafkaConfig struct {
	KAFKA_SERVER         string
	KAFKA_PORT           int
	KAFKA_CONSUMER_GROUP string
	TOPIC_EMAIL_VERIFIED string
	TOPIC_PASS_FORGOT    string
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
	if strings.ToLower(os.Getenv("USING_SECURE")) == "true" {
		cfg.Http.URL = "https://" + cfg.Http.Server
	} else {
		cfg.Http.URL = "http://" + cfg.Http.Server
	}

	if cfg.Http.Port != 0 {
		cfg.Http.URL += fmt.Sprintf(":%d", cfg.Http.Port)
	}
	cfg.Http.URL += "/"

	cfg.DB.DB_DRIVER = strings.ToLower(os.Getenv("DB_DRIVER"))
	cfg.DB.DB_USER = os.Getenv("DB_USER")
	cfg.DB.DB_PASS = os.Getenv("DB_PASS")
	cfg.DB.DB_SERVER = os.Getenv("DB_SERVER")
	cfg.DB.DB_PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.PanicW("DB_PORT must be number", err)
		panic("DB_PORT must be number")
	}
	cfg.DB.DB_NAME = os.Getenv("DB_NAME")
	cfg.DB.DB_MAX_IDLE_CON, err = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CON"))
	if err != nil {
		log.PanicW("DB_MAX_IDLE_CON must be number", err)
		panic("DB_MAX_IDLE_CON must be number")
	}
	cfg.DB.DB_MAX_OPEN_CON, err = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CON"))
	if err != nil {
		log.PanicW("DB_MAX_OPEN_CON must be number", err)
		panic("DB_MAX_OPEN_CON must be number")
	}
	cfg.DB.DB_MAX_LIFE_TIME, err = strconv.Atoi(os.Getenv("DB_MAX_LIFE_TIME"))
	if err != nil {
		log.PanicW("DB_MAX_LIFE_TIME must be number", err)
		panic("DB_MAX_LIFE_TIME must be number")
	}

	cfg.RDB.RDB_ADDRESS = os.Getenv("RDB_ADDRESS")
	cfg.RDB.RDB_PORT, err = strconv.Atoi(os.Getenv("RDB_PORT"))
	if err != nil {
		log.PanicW("RDB_PORT must be number", err)
		panic("RDB_PORT must be number")
	}
	cfg.RDB.RDB_USER = os.Getenv("RDB_USER")
	cfg.RDB.RDB_PASS = os.Getenv("RDB_PASS")
	cfg.RDB.RDB_DB_DEFAULT, err = strconv.Atoi(os.Getenv("RDB_DB_DEFAULT"))
	if err != nil {
		log.PanicW("RDB_DB_DEFAULT must be number", err)
		panic("RDB_DB_DEFAULT must be number")
	}

	cfg.Email.EMAIL_HOST = os.Getenv("EMAIL_HOST")
	cfg.Email.EMAIL_PORT, err = strconv.Atoi(os.Getenv("EMAIL_PORT"))
	if err != nil {
		log.PanicW("EMAIL_PORT must be number", err)
		panic("EMAIL_PORT must be number")
	}
	cfg.Email.EMAIL_NAME = os.Getenv("EMAIL_NAME")
	cfg.Email.EMAIL_EMAIL = os.Getenv("EMAIL_EMAIL")
	cfg.Email.EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")

	cfg.KAFKA.KAFKA_SERVER = os.Getenv("KAFKA_SERVER")
	cfg.KAFKA.KAFKA_PORT, err = strconv.Atoi(os.Getenv("KAFKA_PORT"))
	if err != nil {
		log.PanicW("KAFKA_PORT must be number", err)
		panic("KAFKA_PORT must be number")
	}
	cfg.KAFKA.KAFKA_CONSUMER_GROUP = os.Getenv("KAFKA_CONSUMER_GROUP")
	cfg.KAFKA.TOPIC_EMAIL_VERIFIED = os.Getenv("TOPIC_EMAIL_VERIFIED")
	cfg.KAFKA.TOPIC_PASS_FORGOT = os.Getenv("TOPIC_PASS_FORGOT")

	cfg.OTHER.AlQuranAPI = os.Getenv("ALQURAN_API")

	cfg.JWT.SecretKey = os.Getenv("JWT_SECRET_KEY")
	cfg.JWT.Expired, err = strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	if err != nil {
		log.PanicW("JWT_EXPIRED must be number", err)
		panic("JWT_EXPIRED must be number")
	}

	log.Info("Success for load all configuration")
	return cfg
}
