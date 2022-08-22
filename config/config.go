package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/riskykurniawan15/xarch/helpers/env"
)

type Config struct {
	Http       HttpServer
	DB         DBServer
	RDB        RDBServer
	Email      EmailSender
	KAFKA      KafkaConfig
	Cloudinary Cloudinary
	JWT        JWTConfig
	OTHER      Other
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

type Cloudinary struct {
	CLOUD_NAME string
	API_KEY    string
	API_SECRET string
}

type JWTConfig struct {
	SecretKey string
	Expired   int
}

type Other struct {
	AlQuranAPI string
}

func Configuration() Config {

	err := env.LoadEnv(".env")
	if err != nil {
		log.Println(fmt.Sprintf("error read .env file %w", err.Error()))
	}

	env.LoadFlags()

	cfg := Config{
		Http:       loadHttpServer(),
		DB:         loadDBServer(),
		RDB:        loadRDBServer(),
		Email:      loadEmailSender(),
		KAFKA:      loadKafkaConfig(),
		Cloudinary: loadCloudinary(),
		JWT:        loadJWTConfig(),
		OTHER:      loadOther(),
	}

	fmt.Printf("%-v", loadHttpServer())

	log.Println("Success for load all configuration")

	return cfg
}

func loadHttpServer() HttpServer {
	var cfg HttpServer

	cfg.Server = env.GetString("SERVER")
	cfg.Port = env.GetInt("PORT")
	if env.GetBoolean("USING_SECURE") {
		cfg.URL = "https://" + cfg.Server
	} else {
		cfg.URL = "http://" + cfg.Server
	}
	if cfg.Port != 0 {
		cfg.URL += fmt.Sprintf(":%d", cfg.Port)
	}
	cfg.URL += "/"

	return cfg
}

func loadDBServer() DBServer {
	return DBServer{
		DB_DRIVER:        strings.ToLower(env.GetString("DB_DRIVER")),
		DB_USER:          env.GetString("DB_USER"),
		DB_PASS:          env.GetString("DB_PASS"),
		DB_SERVER:        env.GetString("DB_SERVER"),
		DB_PORT:          env.GetInt("DB_PORT"),
		DB_NAME:          env.GetString("DB_NAME"),
		DB_MAX_IDLE_CON:  env.GetInt("DB_MAX_IDLE_CON"),
		DB_MAX_OPEN_CON:  env.GetInt("DB_MAX_OPEN_CON"),
		DB_MAX_LIFE_TIME: env.GetInt("DB_MAX_LIFE_TIME"),
	}
}

func loadRDBServer() RDBServer {
	return RDBServer{
		RDB_ADDRESS:    env.GetString("RDB_ADDRESS"),
		RDB_PORT:       env.GetInt("RDB_PORT"),
		RDB_USER:       env.GetString("RDB_USER"),
		RDB_PASS:       env.GetString("RDB_PASS"),
		RDB_DB_DEFAULT: env.GetInt("RDB_DB_DEFAULT"),
	}
}

func loadEmailSender() EmailSender {
	return EmailSender{
		EMAIL_HOST:     env.GetString("EMAIL_HOST"),
		EMAIL_PORT:     env.GetInt("EMAIL_PORT"),
		EMAIL_NAME:     env.GetString("EMAIL_NAME"),
		EMAIL_EMAIL:    env.GetString("EMAIL_EMAIL"),
		EMAIL_PASSWORD: env.GetString("EMAIL_PASSWORD"),
	}
}

func loadKafkaConfig() KafkaConfig {
	return KafkaConfig{
		KAFKA_SERVER:         env.GetString("KAFKA_SERVER"),
		KAFKA_PORT:           env.GetInt("KAFKA_PORT"),
		KAFKA_CONSUMER_GROUP: env.GetString("KAFKA_CONSUMER_GROUP"),
		TOPIC_EMAIL_VERIFIED: env.GetString("TOPIC_EMAIL_VERIFIED"),
		TOPIC_PASS_FORGOT:    env.GetString("TOPIC_PASS_FORGOT"),
	}
}

func loadJWTConfig() JWTConfig {
	return JWTConfig{
		SecretKey: env.GetString("JWT_SECRET_KEY"),
		Expired:   env.GetInt("JWT_EXPIRED"),
	}
}

func loadOther() Other {
	return Other{
		AlQuranAPI: env.GetString("ALQURAN_API"),
	}
}

func loadCloudinary() Cloudinary {
	return Cloudinary{
		CLOUD_NAME: env.GetString("CLOUD_NAME"),
		API_KEY:    env.GetString("API_KEY"),
		API_SECRET: env.GetString("API_SECRET"),
	}
}
