package legalentities

import (
	"time"

	"example.com/local/Go2part/domain"
	"gorm.io/gorm"
)

// ORM-модель.
type LegalEntityORM struct {
	UUID      string    `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

// ЯВНОЕ имя таблицы, чтобы не было расхождений.
func (LegalEntityORM) TableName() string { return "legal_entity_orms" }

// Маппинги домен <-> ORM.
func (o *LegalEntityORM) ToDomain() *domain.LegalEntity {
	return &domain.LegalEntity{
		UUID:      o.UUID,
		Name:      o.Name,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}

func LegalEntityFromDomain(d *domain.LegalEntity) *LegalEntityORM {
	return &LegalEntityORM{
		UUID:      d.UUID,
		Name:      d.Name,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

// Автомиграция таблицы.
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&LegalEntityORM{})
}
