package legalentities

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	List(ctx context.Context) ([]LegalEntity, error)
	Create(ctx context.Context, entity *LegalEntity) error
	Update(ctx context.Context, entity *LegalEntity) error
	Delete(ctx context.Context, uuid string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) List(ctx context.Context) ([]LegalEntity, error) {
	var entities []LegalEntity
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&entities).Error
	return entities, err
}

func (r *repository) Create(ctx context.Context, entity *LegalEntity) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *repository) Update(ctx context.Context, entity *LegalEntity) error {
	entity.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Model(&LegalEntity{}).
		Where("uuid = ? AND deleted_at IS NULL", entity.UUID).
		Updates(map[string]interface{}{
			"name":       entity.Name,
			"updated_at": entity.UpdatedAt,
		}).Error
}

func (r *repository) Delete(ctx context.Context, uuid string) error {
	now := time.Now()
	res := r.db.WithContext(ctx).Model(&LegalEntity{}).
		Where("uuid = ? AND deleted_at IS NULL", uuid).
		Update("deleted_at", now)

	log.Printf("Delete: UUID=%s, RowsAffected=%d, Error=%v", uuid, res.RowsAffected, res.Error)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
