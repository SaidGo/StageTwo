package olegalentity

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"example.com/local/Go2part/domain"
	"example.com/local/Go2part/internal/legalentities"
)

type LegalEntityHandler struct {
	service legalentities.ServiceInterface
}

func NewLegalEntityHandler(service legalentities.ServiceInterface) *LegalEntityHandler {
	return &LegalEntityHandler{service: service}
}

var _ ServerInterface = &LegalEntityHandler{}

// uuidToString конвертирует [16]byte UUID в строку с дефисами.
func uuidToString(uuid UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		uuid[0:4],
		uuid[4:6],
		uuid[6:8],
		uuid[8:10],
		uuid[10:16])
}

func (h *LegalEntityHandler) GetLegalEntities(c *gin.Context, params GetLegalEntitiesParams) {
	entities, err := h.service.GetAllLegalEntities(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var dtos []LegalEntityDTO
	for _, e := range entities {
		dtos = append(dtos, LegalEntityDTO{
			UUID:      e.UUID,
			Name:      e.Name,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, dtos)
}

func (h *LegalEntityHandler) CreateLegalEntity(c *gin.Context) {
	var input CreateLegalEntity
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity := &domain.LegalEntity{
		Name: input.Name,
	}
	if err := h.service.CreateLegalEntity(c.Request.Context(), entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dto := LegalEntityDTO{
		UUID:      entity.UUID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	c.JSON(http.StatusCreated, dto)
}

// ✔ Теперь читаем по UUID через доменный сервис, а не через List()+поиск
func (h *LegalEntityHandler) GetLegalEntityByUUID(c *gin.Context, uuid UUID) {
	id := uuidToString(uuid)

	entity, err := h.service.GetLegalEntity(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if entity == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	dto := LegalEntityDTO{
		UUID:      entity.UUID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	c.JSON(http.StatusOK, dto)
}

func (h *LegalEntityHandler) UpdateLegalEntity(c *gin.Context, uuid UUID) {
	var input UpdateLegalEntity
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity := &domain.LegalEntity{
		UUID: uuidToString(uuid),
		Name: input.Name,
		// UpdatedAt выставит доменный сервис
		UpdatedAt: time.Now(),
	}
	if err := h.service.UpdateLegalEntity(c.Request.Context(), entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.GetLegalEntity(c.Request.Context(), entity.UUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dto := LegalEntityDTO{
		UUID:      updated.UUID,
		Name:      updated.Name,
		CreatedAt: updated.CreatedAt,
		UpdatedAt: updated.UpdatedAt,
	}
	c.JSON(http.StatusOK, dto)
}

func (h *LegalEntityHandler) DeleteLegalEntity(c *gin.Context, uuid UUID) {
	err := h.service.DeleteLegalEntity(c.Request.Context(), uuidToString(uuid))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// Aliases for OpenAPI compatibility.
func (h *LegalEntityHandler) DeleteLegalEntitiesUuid(c *gin.Context, uuid UUID) {
	h.DeleteLegalEntity(c, uuid)
}

func (h *LegalEntityHandler) PostLegalEntities(c *gin.Context) {
	h.CreateLegalEntity(c)
}

func (h *LegalEntityHandler) GetLegalEntitiesUuid(c *gin.Context, uuid UUID) {
	h.GetLegalEntityByUUID(c, uuid)
}

func (h *LegalEntityHandler) PutLegalEntitiesUuid(c *gin.Context, uuid UUID) {
	h.UpdateLegalEntity(c, uuid)
}

type LegalEntityDTO struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
