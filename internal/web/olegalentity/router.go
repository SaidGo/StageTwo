package olegalentity

import (
	"github.com/gin-gonic/gin"
)

// В oapi-codegen сгенерирована функция:
//
//	func RegisterHandlers(router gin.IRouter, si ServerInterface)
//
// Здесь вызов-обёртка под наш роутер.
func RegisterLegalEntityRoutes(r *gin.Engine, si ServerInterface) {
	RegisterHandlers(r, si)
}
