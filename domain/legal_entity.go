package domain

import "time"

type LegalEntity struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UUID      string
	Name      string
	INN       string
	KPP       string
	OGRN      string
	Address   string
}
