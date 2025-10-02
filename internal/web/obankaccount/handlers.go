package obankaccount

import (
	"net/http"

	"example.com/local/Go2part/internal/legalentities"
	"github.com/gin-gonic/gin"
)

type Handler struct{ svc *legalentities.Service }

func NewHandler(s *legalentities.Service) *Handler { return &Handler{svc: s} }

func (h *Handler) Register(r *gin.Engine) {
	r.GET("/bank-accounts", func(c *gin.Context) {
		accounts, err := h.svc.GetAllBankAccounts(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, accounts)
	})
}
