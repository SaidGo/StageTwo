package postgres

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db
}
