//go:build disable_extras

package activities

import (
	"example.com/local/Go2part/pkg/postgres"
	"github.com/google/uuid"
)

type Repository struct {
	gorm *postgres.GDB
}

func NewRepository(db *postgres.GDB) *Repository {
	return &Repository{
		gorm: db,
	}
}

func (r *Repository) CreateActivity(activity *Activity) error {
	return r.gorm.DB.Create(activity).Error
}

func (r *Repository) GetTaskActivities(taskUID uuid.UUID, limit, offset int) (orms []Activity, total int64, err error) {
	err = r.gorm.DB.
		Select("*, count(*) OVER() AS total").
		Where("entity_uuid = ?", taskUID).
		Where("entity_type = ?", "task").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&orms).
		Error

	if len(orms) > 0 {
		total = orms[0].Total
	}

	return orms, total, err
}
