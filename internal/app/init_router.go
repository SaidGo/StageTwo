package app

import (
	"example.com/local/Go2part/internal/web"
	"github.com/gin-gonic/gin"
)

// InitializeRouter возвращает базовый *gin.Engine.
// Маршруты добавляются внутри web.NewRouter().
func InitializeRouter() *gin.Engine {
	return web.NewRouter()
}
