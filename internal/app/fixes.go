package app

import "gorm.io/gorm"

// applyBankAccountsFix — заглушка. Нужна для совместимости с wire_gen.go.
// Если есть реальная логика миграции — перенеси её сюда. Сейчас просто no-op.
func applyBankAccountsFix(_ *gorm.DB) error {
	return nil
}
