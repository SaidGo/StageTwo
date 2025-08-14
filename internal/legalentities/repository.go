package legalentities

import (
	"context"

	"example.com/local/Go2part/domain"
	"gorm.io/gorm"
)

type Repository interface {
	List(ctx context.Context) ([]*domain.LegalEntity, error)
	Get(ctx context.Context, uuid string) (*domain.LegalEntity, error)
	Create(ctx context.Context, entity *domain.LegalEntity) error
	Update(ctx context.Context, entity *domain.LegalEntity) error
	Delete(ctx context.Context, uuid string) error
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	// миграция на всякий случай
	if err := AutoMigrate(db); err != nil {
		panic(err)
	}
	return &GormRepository{db: db}
}

// Обёртка, если где-то используется именно NewRepository.
func NewRepository(db *gorm.DB) Repository {
	if err := AutoMigrate(db); err != nil {
		panic(err)
	}
	return &GormRepository{db: db}
}

func (r *GormRepository) List(ctx context.Context) ([]*domain.LegalEntity, error) {
	var ormEntities []LegalEntityORM
	if err := r.db.WithContext(ctx).Find(&ormEntities).Error; err != nil {
		return nil, err
	}
	var result []*domain.LegalEntity
	for _, o := range ormEntities {
		result = append(result, o.ToDomain())
	}
	return result, nil
}

func (r *GormRepository) Get(ctx context.Context, uuid string) (*domain.LegalEntity, error) {
	var orm LegalEntityORM
	if err := r.db.WithContext(ctx).
		Where("uuid = ?", uuid).
		First(&orm).Error; err != nil {
		return nil, err
	}
	return orm.ToDomain(), nil
}

func (r *GormRepository) Create(ctx context.Context, entity *domain.LegalEntity) error {
	return r.db.WithContext(ctx).Create(LegalEntityFromDomain(entity)).Error
}

func (r *GormRepository) Update(ctx context.Context, entity *domain.LegalEntity) error {
	return r.db.WithContext(ctx).
		Model(&LegalEntityORM{}).
		Where("uuid = ?", entity.UUID).
		Updates(map[string]any{
			"name":       entity.Name,
			"updated_at": entity.UpdatedAt,
		}).Error
}

func (r *GormRepository) Delete(ctx context.Context, uuid string) error {
	return r.db.WithContext(ctx).Delete(&LegalEntityORM{}, "uuid = ?", uuid).Error
}
