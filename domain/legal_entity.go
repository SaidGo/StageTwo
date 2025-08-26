package domain

import "time"

// LegalEntity — доменная модель юрлица.
// Расширено для 2.2: добавлены CompanyUUID и BankAccounts.
type LegalEntity struct {
	UUID         string        `json:"uuid"`
	Name         string        `json:"name"`
	CompanyUUID  string        `json:"company_uuid,omitempty"`
	BankAccounts []BankAccount `json:"bank_accounts,omitempty"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at,omitempty"`
}
