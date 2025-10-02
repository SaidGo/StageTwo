//go:build !disable_extras

package obankaccount

import (
	"example.com/local/Go2part/internal/legalentities"
	"github.com/gin-gonic/gin"
)

// Заглушка для совместимости с web/router.go
func RegisterRoutes(r *gin.Engine, svc *legalentities.Service) {
	// no-op
}
