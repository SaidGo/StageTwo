package domain

import "time"

type LegalEntity struct {
	UUID      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
