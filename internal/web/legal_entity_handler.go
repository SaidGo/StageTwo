package web

import (
	"net/http"

	"example.com/local/Go2part/internal/legalentities"

	"github.com/labstack/echo/v4"
)

type LegalEntityHandler struct {
	service *legalentities.Service
}

func NewLegalEntityHandler(service *legalentities.Service) *LegalEntityHandler {
	return &LegalEntityHandler{service: service}
}

func (h *LegalEntityHandler) GetLegalEntities(ctx echo.Context) error {
	entities, err := h.service.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, entities)
}
