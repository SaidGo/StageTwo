package app

import (
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("legalentities.db"), &gorm.Config{})
}
