package legalentities

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]*LegalEntity, error)
	Create(entity *LegalEntity) error
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	db.AutoMigrate(&LegalEntity{})
	return &GormRepository{db: db}
}

func (r *GormRepository) GetAll() ([]*LegalEntity, error) {
	var entities []*LegalEntity
	err := r.db.Find(&entities).Error
	return entities, err
}

func (r *GormRepository) Create(entity *LegalEntity) error {
	return r.db.Create(entity).Error
}

