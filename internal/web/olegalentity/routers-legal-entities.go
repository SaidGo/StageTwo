package olegalentity

import (
	"github.com/gin-gonic/gin"
)

// Регистрация маршрутов для банковских счетов юрлица
func RegisterBankAccountRoutes(r *gin.Engine, leHandler *LegalEntityHandler) {
	group := r.Group("/legal-entities/:uuid/bank-accounts")
	{
		group.GET("", leHandler.GetAllBankAccounts)
		group.POST("", leHandler.CreateBankAccount)
		group.PUT("/:accountId", leHandler.UpdateBankAccount)
		group.DELETE("/:accountId", leHandler.DeleteBankAccount)
	}
}
