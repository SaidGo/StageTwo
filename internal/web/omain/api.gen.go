package omain

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

type About struct {
	BuildTime string `json:"build_time"`
	StartedAt string `json:"started_at"`
	Tag       string `json:"tag"`
	Uuid      string `json:"uuid"`
	Version   string `json:"version"`
}

type Health struct {
	Postgres string `json:"postgres"`
	Redis    string `json:"redis"`
	Status   string `json:"status"`
}

type EntityUUID = openapi_types.UUID

type Uuid = openapi_types.UUID

type ServerInterface interface {
	GetAbout(ctx echo.Context) error

	GetHealth(ctx echo.Context) error
}

type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

func (w *ServerInterfaceWrapper) GetAbout(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetAbout(ctx)
	return err
}

func (w *ServerInterfaceWrapper) GetHealth(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetHealth(ctx)
	return err
}

type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/about", wrapper.GetAbout)
	router.GET(baseURL+"/health", wrapper.GetHealth)

}

type GetAboutRequestObject struct {
}

type GetAboutResponseObject interface {
	VisitGetAboutResponse(w http.ResponseWriter) error
}

type GetAbout200JSONResponse About

func (response GetAbout200JSONResponse) VisitGetAboutResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetHealthRequestObject struct {
}

type GetHealthResponseObject interface {
	VisitGetHealthResponse(w http.ResponseWriter) error
}

type GetHealth200JSONResponse Health

func (response GetHealth200JSONResponse) VisitGetHealthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type StrictServerInterface interface {
	GetAbout(ctx context.Context, request GetAboutRequestObject) (GetAboutResponseObject, error)

	GetHealth(ctx context.Context, request GetHealthRequestObject) (GetHealthResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

func (sh *strictHandler) GetAbout(ctx echo.Context) error {
	var request GetAboutRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetAbout(ctx.Request().Context(), request.(GetAboutRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetAbout")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetAboutResponseObject); ok {
		return validResponse.VisitGetAboutResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) GetHealth(ctx echo.Context) error {
	var request GetHealthRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetHealth(ctx.Request().Context(), request.(GetHealthRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetHealth")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetHealthResponseObject); ok {
		return validResponse.VisitGetHealthResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
