package web

import (
	"github.com/gin-gonic/gin"

	"example.com/local/Go2part/internal/legalentities"
	"example.com/local/Go2part/internal/web/olegalentity"
)

// NewRouter — инициализация маршрутов.
func NewRouter(leSvc legalentities.ServiceInterface) *gin.Engine {
	r := gin.Default()

	// LegalEntity (2.1)
	leHandler := olegalentity.NewLegalEntityHandler(leSvc)
	olegalentity.RegisterLegalEntityRoutes(r, leHandler)

	// BankAccount (2.2)
	olegalentity.RegisterBankAccountRoutes(r, leHandler)

	return r
}
