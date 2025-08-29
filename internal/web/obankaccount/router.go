package obankaccount

import (
	"example.com/local/Go2part/internal/legalentities"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, svc legalentities.ServiceInterface) {
	NewHandler(svc).Register(r)
}
