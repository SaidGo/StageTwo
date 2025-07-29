package olegalentity

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes регистрирует маршруты LegalEntityHandler в роутере Gin.
func RegisterRoutes(router *gin.Engine, handler *LegalEntityHandler) {
	RegisterHandlers(router, handler)
}
