package olegalentity

import (
	"net/http"

	"example.com/local/Go2part/dto"
	"example.com/local/Go2part/internal/legalentities"
	"github.com/gin-gonic/gin"
)

type LegalEntityBankAccountsHandler struct {
	svc legalentities.ServiceInterface
}

func NewLegalEntityBankAccountsHandler(s legalentities.ServiceInterface) *LegalEntityBankAccountsHandler {
	return &LegalEntityBankAccountsHandler{svc: s}
}

func (h *LegalEntityBankAccountsHandler) Register(r *gin.Engine) {
	g := r.Group("/legal-entities/:uuid/bank-accounts")
	g.GET("", h.ListByLE)
}

func (h *LegalEntityBankAccountsHandler) ListByLE(c *gin.Context) {
	le := c.Param("uuid")
	items, err := h.svc.ListBankAccounts(c.Request.Context(), le)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	out := make([]dto.BankAccountView, 0, len(items))
	for i := range items {
		out = append(out, dto.NewBankAccountView(&items[i]))
	}
	c.JSON(http.StatusOK, gin.H{"accounts": out})
}
