package app

import (
	"fmt"

	"example.com/local/Go2part/internal/legalentities"
	"example.com/local/Go2part/internal/web"
)

func InitApp() (*App, error) {
	db, err := NewDB()
	if err != nil {
		return nil, fmt.Errorf("NewDB: %w", err)
	}

	if err := applyBankAccountsFix(db); err != nil {
		return nil, fmt.Errorf("applyBankAccountsFix: %w", err)
	}

	repo := legalentities.NewRepository(db)
	svc := legalentities.NewService(repo)

	router := web.NewRouter(svc)

	app := NewApp(router, svc)
	return app, nil
}
