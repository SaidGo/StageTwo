package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"example.com/local/Go2part/dto"
	"example.com/local/Go2part/internal/legalentities"
	"example.com/local/Go2part/internal/web/obankaccount"
	olegalentity "example.com/local/Go2part/internal/web/olegalentity"
)

type leServer struct {
	svc legalentities.ServiceInterface
}

func toUUID(id olegalentity.UUID) (uuid.UUID, error) {
	if s, ok := any(id).(string); ok {
		return uuid.Parse(s)
	}
	if b, ok := any(id).([16]byte); ok {
		return uuid.UUID(b), nil
	}
	if bs, ok := any(id).([]byte); ok {
		if len(bs) == 16 {
			var u uuid.UUID
			copy(u[:], bs)
			return u, nil
		}
		return uuid.Parse(string(bs))
	}
	return uuid.Parse(fmt.Sprintf("%v", id))
}

func (h *leServer) GetLegalEntities(c *gin.Context, _ olegalentity.GetLegalEntitiesParams) {
	items, err := h.svc.ListLegalEntities(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *leServer) PostLegalEntities(c *gin.Context) {
	var body dto.LegalEntityCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{Error: "invalid json"})
		return
	}
	created, err := h.svc.CreateLegalEntity(c.Request.Context(), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *leServer) GetLegalEntitiesUuid(c *gin.Context, id olegalentity.UUID) {
	uid, err := toUUID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{Error: "invalid uuid"})
		return
	}
	item, err := h.svc.GetLegalEntity(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Error{Error: "not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *leServer) PutLegalEntitiesUuid(c *gin.Context, id olegalentity.UUID) {
	uid, err := toUUID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{Error: "invalid uuid"})
		return
	}
	var body dto.LegalEntityUpdate
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{Error: "invalid json"})
		return
	}
	updated, err := h.svc.UpdateLegalEntity(c.Request.Context(), uid, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *leServer) DeleteLegalEntitiesUuid(c *gin.Context, id olegalentity.UUID) {
	uid, err := toUUID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Error{Error: "invalid uuid"})
		return
	}
	if err := h.svc.DeleteLegalEntity(c.Request.Context(), uid); err != nil {
		c.JSON(http.StatusNotFound, dto.Error{Error: "not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func NewRouter(leSvc legalentities.ServiceInterface) *gin.Engine {
	r := gin.Default()

	olegalentity.RegisterHandlersWithOptions(
		r,
		&leServer{svc: leSvc},
		olegalentity.GinServerOptions{},
	)

	obankaccount.RegisterRoutes(r, leSvc)

	olegalentity.NewLegalEntityBankAccountsHandler(leSvc).Register(r)

	for _, ri := range r.Routes() {
		log.Printf("[ROUTE] %s %s -> %s", ri.Method, ri.Path, ri.Handler)
	}
	return r
}
