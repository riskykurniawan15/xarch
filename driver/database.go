package driver

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/logger"
)

var log logger.Logger = logger.NewLogger()

func openSQL(cfg config.DBServer) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USER,
		cfg.DB_PASS,
		cfg.DB_SERVER,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)

	sqlDB, _ := sql.Open(cfg.DB_DRIVER, dsn)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(cfg.DB_MAX_IDLE_CON)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(cfg.DB_MAX_OPEN_CON)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(cfg.DB_MAX_LIFE_TIME))

	return sqlDB
}

func ConnectDB(cfg config.DBServer) *gorm.DB {
	defer func() {
		if r := recover(); r != nil {
			log.Panic(fmt.Sprint(r))
		}
	}()

	log.Info("Connection to database")

	if cfg.DB_DRIVER == "mysql" {
		db, err := gorm.Open(mysql.New(mysql.Config{
			Conn: openSQL(cfg),
		}), &gorm.Config{})
		if err != nil {
			panic("Failed to Connect DB")
		}
		return db
	} else {
		panic("Failed to Connect DB but driver not exists")
	}
}
