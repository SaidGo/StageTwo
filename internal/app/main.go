package app

import (
	"example.com/local/Go2part/internal/web"
)

type App struct {
	LegalEntityHandler *web.LegalEntityHandler
}

func NewApp(dsn string) *App {
	return &App{
		LegalEntityHandler: InitLegalEntityHandler(dsn),
	}
}
