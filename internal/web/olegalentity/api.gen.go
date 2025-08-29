package olegalentity

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type CreateLegalEntity struct {
	Name string `json:"name"`
}

type LegalEntity struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Name      *string    `json:"name,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Uuid      *UUID      `json:"uuid,omitempty"`
}

type UUID = openapi_types.UUID

type UpdateLegalEntity struct {
	Name string `json:"name"`
}

type Limit = int

type Offset = int

type GetLegalEntitiesParams struct {
	Limit *Limit `form:"limit,omitempty" json:"limit,omitempty"`

	Offset *Offset `form:"offset,omitempty" json:"offset,omitempty"`
}

type PostLegalEntitiesJSONRequestBody = CreateLegalEntity

type PutLegalEntitiesUuidJSONRequestBody = UpdateLegalEntity

type ServerInterface interface {
	GetLegalEntities(c *gin.Context, params GetLegalEntitiesParams)

	PostLegalEntities(c *gin.Context)

	DeleteLegalEntitiesUuid(c *gin.Context, uuid UUID)

	GetLegalEntitiesUuid(c *gin.Context, uuid UUID)

	PutLegalEntitiesUuid(c *gin.Context, uuid UUID)
}

type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

func (siw *ServerInterfaceWrapper) GetLegalEntities(c *gin.Context) {

	var err error

	var params GetLegalEntitiesParams

	err = runtime.BindQueryParameter("form", true, false, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter limit: %w", err), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, false, "offset", c.Request.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter offset: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetLegalEntities(c, params)
}

func (siw *ServerInterfaceWrapper) PostLegalEntities(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostLegalEntities(c)
}

func (siw *ServerInterfaceWrapper) DeleteLegalEntitiesUuid(c *gin.Context) {

	var err error

	var uuid UUID

	err = runtime.BindStyledParameter("simple", false, "uuid", c.Param("uuid"), &uuid)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uuid: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteLegalEntitiesUuid(c, uuid)
}

func (siw *ServerInterfaceWrapper) GetLegalEntitiesUuid(c *gin.Context) {

	var err error

	var uuid UUID

	err = runtime.BindStyledParameter("simple", false, "uuid", c.Param("uuid"), &uuid)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uuid: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetLegalEntitiesUuid(c, uuid)
}

func (siw *ServerInterfaceWrapper) PutLegalEntitiesUuid(c *gin.Context) {

	var err error

	var uuid UUID

	err = runtime.BindStyledParameter("simple", false, "uuid", c.Param("uuid"), &uuid)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uuid: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutLegalEntitiesUuid(c, uuid)
}

type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/legal-entities", wrapper.GetLegalEntities)
	router.POST(options.BaseURL+"/legal-entities", wrapper.PostLegalEntities)
	router.DELETE(options.BaseURL+"/legal-entities/:uuid", wrapper.DeleteLegalEntitiesUuid)
	router.GET(options.BaseURL+"/legal-entities/:uuid", wrapper.GetLegalEntitiesUuid)
	router.PUT(options.BaseURL+"/legal-entities/:uuid", wrapper.PutLegalEntitiesUuid)
}
