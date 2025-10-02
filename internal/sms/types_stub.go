//go:build !disable_extras

package sms

// Минимальный тип ответа, которого требует contract.go.
type Response struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
}
