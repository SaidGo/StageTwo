package olegalentity

import "github.com/gin-gonic/gin"

func RegisterBankAccountRoutes(_ *gin.Engine, _ *LegalEntityHandler) {
	// no-op: BA вынесены в отдельный роутер, связанная ручка регистрируется через LegalEntityBankAccountsHandler
}
