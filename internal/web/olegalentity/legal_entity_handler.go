package olegalentity

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/local/Go2part/internal/legalentities"
)

// Хендлер для LegalEntity-ручек (совместим со сгенерированным ServerInterface).
type LegalEntityHandler struct {
	service legalentities.ServiceInterface
}

func NewLegalEntityHandler(s legalentities.ServiceInterface) *LegalEntityHandler {
	return &LegalEntityHandler{service: s}
}

// ===== LegalEntity CRUD =====
// Временно 501 — чтобы не требовать отсутствующих методов сервиса.
func (h *LegalEntityHandler) GetLegalEntities(c *gin.Context, params GetLegalEntitiesParams) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
func (h *LegalEntityHandler) PostLegalEntities(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
func (h *LegalEntityHandler) GetLegalEntitiesUuid(c *gin.Context, uuid UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
func (h *LegalEntityHandler) PutLegalEntitiesUuid(c *gin.Context, uuid UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
func (h *LegalEntityHandler) DeleteLegalEntitiesUuid(c *gin.Context, uuid UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ===== Bank Accounts под юрлицом =====

// GET /legal-entities/:c.Param("uuid")/bank-accounts — единственная живая LE-связанная ручка
func (h *LegalEntityHandler) GetAllBankAccounts(c *gin.Context, uuid UUID) {
	leStr := uuid.String()
	items, err := h.service.ListBankAccounts(c.Request.Context(), leStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accounts": items})
}

// Остальные BA-ручки под LE переехали в /bank_accounts — отдаём 410 Gone
