package domain

import (
	"github.com/google/uuid"
	"time"
)

type LegalEntity struct {
	UUID      uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
