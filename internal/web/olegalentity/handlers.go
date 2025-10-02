package olegalentity

import (
	"context"
	"net/http"

	"example.com/local/Go2part/domain"
	"example.com/local/Go2part/internal/legalentities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LegalEntityHandler struct {
	svc *legalentities.Service
}

// Register регистрирует REST ручки для юр. лиц
func (h *LegalEntityHandler) Register(r *gin.Engine) {
	// если хендлер создан без сервиса, поднимем минимальный (in-memory)
	if h.svc == nil {
		h.svc = legalentities.NewService(nil)
	}

	r.GET("/legal-entities", h.getAll)
	r.POST("/legal-entities", h.create)
	r.GET("/legal-entities/:id", h.getOne)
	r.PUT("/legal-entities/:id", h.update)
	r.DELETE("/legal-entities/:id", h.delete)
}

func (h *LegalEntityHandler) getAll(c *gin.Context) {
	ents, err := h.svc.GetAllLegalEntities(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ents)
}

func (h *LegalEntityHandler) create(c *gin.Context) {
	var in domain.LegalEntity
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	out, err := h.svc.CreateLegalEntity(c.Request.Context(), in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *LegalEntityHandler) getOne(c *gin.Context) {
	id := c.Param("id") // строковый UUID
	e, err := h.svc.GetLegalEntity(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, e)
}

func (h *LegalEntityHandler) update(c *gin.Context) {
	id := c.Param("id")
	var in domain.LegalEntity
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsed, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	in.UUID = parsed

	out, err := h.svc.UpdateLegalEntity(context.Background(), in)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *LegalEntityHandler) delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.DeleteLegalEntity(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
