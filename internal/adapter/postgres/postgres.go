package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"go-test-api/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(cfg model.PostgresClient) *gorm.DB {
	db, err := setupDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect postgres client: %v", err)
	}
	return db
}

func setupDB(conf model.PostgresClient) (*gorm.DB, error) {
	newLoggerCfg := logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Silent,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      true,
		Colorful:                  false,
	}
	if conf.DebugMode {
		newLoggerCfg.LogLevel = logger.Info
		newLoggerCfg.IgnoreRecordNotFoundError = false
		newLoggerCfg.Colorful = true
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		newLoggerCfg,
	)

	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s", conf.Db, conf.Username, conf.Password, conf.Host, conf.Port, conf.SslMode)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(conf.MaxIdleConn)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConn)
	return db, nil
}
