package dto

import (
	"time"

	"example.com/local/Go2part/domain"
	"github.com/google/uuid"
)

type BankAccountCreateDTO struct {
	LegalEntityUUID *uuid.UUID `json:"legal_entity_uuid,omitempty"`
	BIK             string     `json:"bik"`
	Bank            string     `json:"bank"`
	Address         string     `json:"address"`
	CorrAccount     string     `json:"corr_account"`
	Account         string     `json:"account"`
	Currency        string     `json:"currency"`
	Comment         string     `json:"comment"`
	IsPrimary       bool       `json:"is_primary"`
}

type BankAccountUpdateDTO = BankAccountCreateDTO

type BankAccountView struct {
	UUID            uuid.UUID  `json:"uuid"`
	LegalEntityUUID *uuid.UUID `json:"legal_entity_uuid,omitempty"`
	BIK             string     `json:"bik"`
	Bank            string     `json:"bank"`
	Address         string     `json:"address"`
	CorrAccount     string     `json:"corr_account"`
	Account         string     `json:"account"`
	Currency        string     `json:"currency"`
	Comment         string     `json:"comment"`
	IsPrimary       bool       `json:"is_primary"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

func (d BankAccountCreateDTO) ToDomain() *domain.BankAccount {
	var le uuid.UUID
	if d.LegalEntityUUID != nil {
		le = *d.LegalEntityUUID
	}
	return &domain.BankAccount{
		LegalEntityUUID: le,
		BIK:             d.BIK,
		Bank:            d.Bank,
		Address:         d.Address,
		CorrAccount:     d.CorrAccount,
		Account:         d.Account,
		Currency:        d.Currency,
		Comment:         d.Comment,
		IsPrimary:       d.IsPrimary,
	}
}

func ApplyUpdate(dst *domain.BankAccount, d BankAccountUpdateDTO) {
	if d.LegalEntityUUID != nil {
		dst.LegalEntityUUID = *d.LegalEntityUUID
	}
	dst.BIK = d.BIK
	dst.Bank = d.Bank
	dst.Address = d.Address
	dst.CorrAccount = d.CorrAccount
	dst.Account = d.Account
	dst.Currency = d.Currency
	dst.Comment = d.Comment
	dst.IsPrimary = d.IsPrimary
}

func NewBankAccountView(a *domain.BankAccount) BankAccountView {
	var le *uuid.UUID
	if a.LegalEntityUUID != uuid.Nil {
		tmp := a.LegalEntityUUID
		le = &tmp
	}
	return BankAccountView{
		UUID:            a.UUID,
		LegalEntityUUID: le,
		BIK:             a.BIK,
		Bank:            a.Bank,
		Address:         a.Address,
		CorrAccount:     a.CorrAccount,
		Account:         a.Account,
		Currency:        a.Currency,
		Comment:         a.Comment,
		IsPrimary:       a.IsPrimary,
		CreatedAt:       a.CreatedAt,
		UpdatedAt:       a.UpdatedAt,
		DeletedAt:       a.DeletedAt,
	}
}
