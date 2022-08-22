package driver

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/riskykurniawan15/xarch/config"
)

func openSQL(cfg config.DBServer) *sql.DB {
	var driver, dsn string
	driver = cfg.DB_DRIVER
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USER,
		cfg.DB_PASS,
		cfg.DB_SERVER,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)

	sqlDB, _ := sql.Open(driver, dsn)
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

	log.Println("Connection to database")

	if cfg.DB_DRIVER == "mysql" {
		db, err := gorm.Open(mysql.New(mysql.Config{
			Conn: openSQL(cfg),
		}), &gorm.Config{})
		if err != nil {
			panic("Failed to Connect mysql")
		}
		return db
	} else if cfg.DB_DRIVER == "sqlserver" {
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			cfg.DB_USER,
			cfg.DB_PASS,
			cfg.DB_SERVER,
			cfg.DB_PORT,
			cfg.DB_NAME,
		)
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to Connect postgresql")
		}
		return db
	} else {
		panic("Failed to Connect DB but driver not exists")
	}
}
