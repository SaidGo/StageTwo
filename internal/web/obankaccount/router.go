//go:build disable_extras

package obankaccount

import (
	"example.com/local/Go2part/internal/legalentities"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, svc *legalentities.Service) {
	NewHandler(svc).Register(r)
}
