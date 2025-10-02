package gates

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"example.com/local/Go2part/domain"
	"github.com/google/uuid"
)

type Permission struct {
	UserUUID       uuid.UUID `json:"user_uuid" gorm:"<-:create;"`
	FederationUUID uuid.UUID `json:"federation_uuid" gorm:"<-:create;"`

	Rules domain.PermissionRules `json:"rules"`

	CreatedAt time.Time  `json:"created_at" gorm:"<-:create;"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type JSON domain.PermissionRules

func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := domain.PermissionRules{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}
