package app

import (
	"example.com/local/Go2part/internal/web/olegalentity"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler olegalentity.ServerInterface) *gin.Engine {
	r := gin.Default()
	olegalentity.RegisterHandlers(r, handler)
	return r
}
