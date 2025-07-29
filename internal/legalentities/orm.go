package legalentities

import (
	"time"
)

type LegalEntity struct {
	UUID      string     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"uuid"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}
