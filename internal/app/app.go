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

func NewApp(router *gin.Engine, leService legalentities.ServiceInterface) *App {
	return &App{
		Router:             router,
		LegalEntityService: leService,
	}
}
