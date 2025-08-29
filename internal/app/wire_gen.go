// Code generated manually to replace wire on MSYS. DO NOT EDIT BY WIRE.

package app

import (
	"fmt"

	"example.com/local/Go2part/internal/legalentities"
	"example.com/local/Go2part/internal/web"
)

// InitApp — ручная инициализация зависимостей без google/wire.
func InitApp() (*App, error) {
	db, err := NewDB()
	if err != nil {
		return nil, fmt.Errorf("NewDB: %w", err)
	}

	// ВАЖНО: фиксация схемы BA (снятие NOT NULL/FK и DEFAULT для uuid)
	if err := applyBankAccountsFix(db); err != nil {
		return nil, fmt.Errorf("applyBankAccountsFix: %w", err)
	}

	repo := legalentities.NewRepository(db)
	svc := legalentities.NewService(repo)

	router := web.NewRouter(svc)

	app := NewApp(router, svc)
	return app, nil
}
