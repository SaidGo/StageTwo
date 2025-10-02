package dto

// LegalEntity — DTO для выдачи сущности
type LegalEntity struct {
	UUID    string `json:"uuid"    ru:"UUID"`
	Name    string `json:"name"    ru:"Название"`
	INN     string `json:"inn"     ru:"ИНН"`
	KPP     string `json:"kpp,omitempty"  ru:"КПП"`
	OGRN    string `json:"ogrn,omitempty" ru:"ОГРН"`
	Address string `json:"address,omitempty" ru:"Адрес"`
}

// LegalEntityCreate — DTO для создания
type LegalEntityCreate struct {
	Name    string `json:"name"              validate:"required,trim,name"        ru:"Название"`
	INN     string `json:"inn"               validate:"required,legal_entity_field" ru:"ИНН"`
	KPP     string `json:"kpp,omitempty"     validate:"omitempty,legal_entity_field" ru:"КПП"`
	OGRN    string `json:"ogrn,omitempty"    validate:"omitempty,legal_entity_field" ru:"ОГРН"`
	Address string `json:"address,omitempty" validate:"omitempty,trim"            ru:"Адрес"`
}

// LegalEntityUpdate — DTO для частичного обновления
type LegalEntityUpdate struct {
	Name    *string `json:"name,omitempty"    validate:"omitempty,trim,name"          ru:"Название"`
	INN     *string `json:"inn,omitempty"     validate:"omitempty,legal_entity_field" ru:"ИНН"`
	KPP     *string `json:"kpp,omitempty"     validate:"omitempty,legal_entity_field" ru:"КПП"`
	OGRN    *string `json:"ogrn,omitempty"    validate:"omitempty,legal_entity_field" ru:"ОГРН"`
	Address *string `json:"address,omitempty" validate:"omitempty,trim"               ru:"Адрес"`
}

// Совместимость со старым именем
type LegalEntityDTO = LegalEntity
