// Code generated manually to replace wire on MSYS. DO NOT EDIT BY WIRE.

package app

import (
	"fmt"

	"example.com/local/Go2part/internal/legalentities"
	"example.com/local/Go2part/internal/web"
)

// InitApp инициализирует зависимости приложения без google/wire.
// Порядок:
//
//	NewDB -> legalentities.NewRepository -> legalentities.NewService ->
//	web.NewRouter(service) -> NewApp(router, service)
func InitApp() (*App, error) {
	db, err := NewDB()
	if err != nil {
		return nil, fmt.Errorf("NewDB: %w", err)
	}

	repo := legalentities.NewRepository(db)
	svc := legalentities.NewService(repo)

	router := web.NewRouter(svc)

	app := NewApp(router, svc)
	return app, nil
}
