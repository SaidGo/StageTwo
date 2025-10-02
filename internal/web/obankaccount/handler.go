//go:build disable_extras

package obankaccount

import (
	"net/http"

	"example.com/local/Go2part/dto"
	"example.com/local/Go2part/internal/legalentities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	svc *legalentities.Service
}

func NewHandler(s *legalentities.Service) *Handler { return &Handler{svc: s} }

func (h *Handler) Register(r *gin.Engine) {
	g := r.Group("/bank_accounts")
	g.POST("", h.Create)
	g.GET("", h.List)
	g.GET("/:uuid", h.Get)
	g.PUT("/:uuid", h.Update)
	g.DELETE("/:uuid", h.Delete)
}

func (h *Handler) Create(c *gin.Context) {
	var in dto.BankAccountCreateDTO
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	acc := in.ToDomain()
	created, err := h.svc.CreateBankAccount(c.Request.Context(), acc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.NewBankAccountView(&created))
}

func (h *Handler) List(c *gin.Context) {
	items, err := h.svc.ListAllBankAccounts(c.Request.Context())
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

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("uuid")
	acc, err := h.svc.GetBankAccount(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, dto.NewBankAccountView(&acc))
}

func (h *Handler) Update(c *gin.Context) {
	raw := c.Param("uuid")
	u, err := uuid.Parse(raw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	var in dto.BankAccountUpdateDTO
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	acc := in.ToDomain()
	acc.UUID = u
	updated, err := h.svc.UpdateBankAccount(c.Request.Context(), acc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, dto.NewBankAccountView(&updated))
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("uuid")
	if err := h.svc.DeleteBankAccount(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
