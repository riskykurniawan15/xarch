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

	sqlDB, _ := sql.Open("mysql", dsn)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return sqlDB
}

func ConnectDB(cfg config.DBServer) *gorm.DB {
	log.Info("Connection to database")

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: openSQL(cfg),
	}), &gorm.Config{})
	if err != nil {
		log.PanicW("Failed to Connect DB", err)
		panic("Failed to Connect DB")
	}

	log.Info("Database connected")

	return db
}
