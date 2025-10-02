package domain

import (
	"time"

	"gorm.io/gorm"
)

// Meta — общий встраиваемый блок временных меток/soft delete.
type Meta struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
