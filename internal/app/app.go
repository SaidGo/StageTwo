package app

import (
	"github.com/gin-gonic/gin"

	"example.com/local/Go2part/internal/legalentities"
)

type App struct {
	Router             *gin.Engine
	LegalEntityService legalentities.ServiceInterface
}

func (a *App) Run(addr string) error {
	return a.Router.Run(addr)
}

// NewApp — принимает готовый роутер и сервисы.
func NewApp(router *gin.Engine, leService legalentities.ServiceInterface) *App {
	return &App{
		Router:             router,
		LegalEntityService: leService,
	}
}
