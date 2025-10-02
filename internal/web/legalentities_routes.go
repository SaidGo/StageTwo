//go:build disable_extras

package web

import (
	"net/http"

	"example.com/local/Go2part/internal/legalentities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterLegalEntitiesRoutes(r *gin.Engine, svc *legalentities.Service) {
	g := r.Group("/legal-entities")

	g.GET("", func(c *gin.Context) {
		items, _ := svc.GetAllLegalEntities(c.Request.Context())
		c.JSON(http.StatusOK, items)
	})

	type createReq struct {
		Name string `json:"name" binding:"required"`
	}
	g.POST("", func(c *gin.Context) {
		var in createReq
		if err := c.ShouldBindJSON(&in); err != nil || in.Name == "" {
			c.Status(http.StatusBadRequest)
			return
		}
		item, _ := svc.CreateLegalEntity(c.Request.Context(), in.Name)
		c.JSON(http.StatusCreated, item)
	})

	g.GET("/:id", func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		item, err := svc.GetLegalEntity(c.Request.Context(), id)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, item)
	})

	type updateReq struct {
		Name string `json:"name" binding:"required"`
	}
	g.PUT("/:id", func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		var in updateReq
		if err := c.ShouldBindJSON(&in); err != nil || in.Name == "" {
			c.Status(http.StatusBadRequest)
			return
		}
		item, err := svc.UpdateLegalEntity(c.Request.Context(), id, in.Name)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, item)
	})

	g.DELETE("/:id", func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		if err := svc.DeleteLegalEntity(c.Request.Context(), id); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusNoContent)
	})
}
