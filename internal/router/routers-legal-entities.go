package router

import (
	"example.com/local/Go2part/internal/web"

	"github.com/labstack/echo/v4"
)

func RegisterLegalEntityRoutes(e *echo.Echo, handler *web.LegalEntityHandler) {
    e.GET("/legal-entities", handler.GetLegalEntities)
}

