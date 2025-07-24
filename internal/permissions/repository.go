package permissions

import (
	"context"

	"example.com/local/Go2part/domain"
	"example.com/local/Go2part/dto"
	"example.com/local/Go2part/pkg/postgres"
	"example.com/local/Go2part/pkg/redis"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

type Repository struct {
	gorm *postgres.GDB
	rds  *redis.RDS
}

func NewRepository(db *postgres.GDB, rds *redis.RDS) *Repository {
	return &Repository{
		gorm: db,
		rds:  rds,
	}
}

func (r *Repository) PubUpdate() {
	err := r.rds.Publish(context.Background(), "update", "acl")
	logrus.Debug("pub update acl")
	if err != nil {
		logrus.Error(err)
	}
}

func (r *Repository) UCreateGroup(crtr domain.Creator, name string, permissions []string) (uuid.UUID, error) {
	group := &Group{
		Name: name,
		State: Meta{
			"permissions": permissions,
		},
		CreatedBy: crtr.UUID,
	}

	err := r.gorm.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "uuid"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"state":      group.State,
			"name":       group.Name,
			"updated_at": "now()",
			"deleted_at": nil,
		}),
	}).Create(&group).Error

	return group.UUID, err
}

func (r *Repository) DeleteGroup(uid uuid.UUID) error {
	res := r.gorm.DB.
		Model(&Group{}).
		Where("uuid = ?", uid).
		Where("deleted_at is null").
		Update("deleted_at", "now()")

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return dto.NotFoundErr("группа не найдена или уже удалена")
	}

	return nil
}

func (r *Repository) AddUserToGroup(userUUID, groupUUID uuid.UUID) error {
	user := &User{
		UserUUID: userUUID,
		Groups:   groupUUID,
	}

	err := r.gorm.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_uuid"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"groups":     user.Groups,
			"updated_at": "now()",
			"deleted_at": nil,
		}),
	}).Create(&user).Error

	return err
}
