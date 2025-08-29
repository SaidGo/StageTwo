package web

import (
	"log"

	"github.com/gin-gonic/gin"

	"example.com/local/Go2part/internal/legalentities"
	"example.com/local/Go2part/internal/web/obankaccount"
	"example.com/local/Go2part/internal/web/olegalentity"
)

// NewRouter — инициализация маршрутов приложения.
func NewRouter(leSvc legalentities.ServiceInterface) *gin.Engine {
	r := gin.Default()

	// LegalEntity (oapi-codegen)
	leHandler := olegalentity.NewLegalEntityHandler(leSvc)
	olegalentity.RegisterLegalEntityRoutes(r, leHandler)

	// BankAccount — 5 независимых ручек:
	obankaccount.RegisterRoutes(r, leSvc)

	// Единственная связанная с LE ручка:
	olegalentity.NewLegalEntityBankAccountsHandler(leSvc).Register(r)

	// Диагностика: печать таблицы маршрутов
	for _, ri := range r.Routes() {
		log.Printf("[ROUTE] %s %s -> %s", ri.Method, ri.Path, ri.Handler)
	}
	return r
}
