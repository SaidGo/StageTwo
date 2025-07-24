package legalentities

import (
	"example.com/local/Go2part/domain"

	"gorm.io/gorm"
)

type LegalEntity struct {
	gorm.Model
	UUID    string `gorm:"type:uuid;uniqueIndex"`
	Name    string
	INN     string
	KPP     string
	OGRN    string
	Address string
}

func (e *LegalEntity) ToDomain() domain.LegalEntity {
	return domain.LegalEntity{
		UUID:    e.UUID,
		Name:    e.Name,
		INN:     e.INN,
		KPP:     e.KPP,
		OGRN:    e.OGRN,
		Address: e.Address,
	}
}
