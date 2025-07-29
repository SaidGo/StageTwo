package app

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func (a *App) Run(addr string) error {
	return a.Router.Run(addr)
}

func NewApp(router *gin.Engine) *App {
	return &App{Router: router}
}
