//go:build disable_extras

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"example.com/local/Go2part/internal/kafka"
	"example.com/local/Go2part/internal/legalentities"
	"example.com/local/Go2part/internal/web"
)

// Собираем все зависимости и возвращаем уже сконфигурированный *gin.Engine.
func InitializeRouter() (*gin.Engine, error) {
	wire.Build(
		NewDB,                       // *sql.DB
		legalentities.NewRepository, // repo
		kafka.NewLegalEntitySender,  // producer for legal-entities-created
		kafka.NewBankAccountSender,  // producer for bank-accounts-created
		legalentities.NewService,    // service
		web.NewRouter,               // -> *gin.Engine
	)
	return nil, nil
}
