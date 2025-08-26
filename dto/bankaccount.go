package dto

type CreateBankAccountRequest struct {
	CompanyID uint   `json:"company_id"`
	IBAN      string `json:"iban"`
	BankName  string `json:"bank_name"`
	Currency  string `json:"currency"`
}

type UpdateBankAccountRequest struct {
	BankName string `json:"bank_name"`
	Currency string `json:"currency"`
}
