package web

import (
	"example.com/local/Go2part/internal/web/olegalentity"
	"github.com/gin-gonic/gin"
)

// NewRouter создает роутер и регистрирует OpenAPI-хендлеры.
func NewRouter(h *olegalentity.LegalEntityHandler) *gin.Engine {
	r := gin.Default()
	olegalentity.RegisterHandlers(r, h)
	return r
}
