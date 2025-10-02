package app

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

func NewDB() (*gorm.DB, error) {
	pgDsn := os.Getenv("POSTGRES_DSN")

	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	}

	var (
		db  *gorm.DB
		err error
	)

	if pgDsn == "" {
		sqlitePath := "legalentities.db"
		db, err = gorm.Open(sqlite.Open(sqlitePath), cfg)
		if err != nil {
			log.Printf("NewDB: open sqlite failed: %v", err)
			return nil, err
		}
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(1)
		sqlDB.SetConnMaxIdleTime(5 * time.Minute)
		log.Printf("NewDB: SQLite fallback used: %s", sqlitePath)
		return db, nil
	}

	db, err = gorm.Open(postgres.Open(pgDsn), cfg)
	if err != nil {
		log.Printf("NewDB: open postgres failed: %v", err)
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	return db, nil
}
