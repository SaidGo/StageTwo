package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("legalentities.db"), &gorm.Config{})
}
