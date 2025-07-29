package app

import (
	"github.com/gin-gonic/gin"
	"example.com/local/Go2part/internal/web/olegalentity"
)

func NewRouter(handler olegalentity.ServerInterface) *gin.Engine {
	r := gin.Default()
	olegalentity.RegisterHandlers(r, handler)
	return r
}
