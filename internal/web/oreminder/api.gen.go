package oreminder

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"example.com/local/Go2part/dto"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

type ReminderCreateRequest struct {
	DateFrom    *time.Time          `json:"date_from,omitempty"`
	DateTo      *time.Time          `json:"date_to,omitempty"`
	Description string              `json:"description" validate:"trim,name,min=0,max=2000"`
	TaskUuid    openapi_types.UUID  `json:"task_uuid" validate:"uuid"`
	Type        string              `json:"type" validate:"trim,name,min=0,max=50"`
	UserUuid    *openapi_types.UUID `json:"user_uuid,omitempty"`
}

type ReminderDTO = dto.ReminderDTO

type ReminderPutRequest struct {
	Comment     string              `json:"comment" validate:"trim,min=0,max=5000"`
	DateFrom    *time.Time          `json:"date_from,omitempty"`
	DateTo      *time.Time          `json:"date_to,omitempty"`
	Description string              `json:"description" validate:"trim,name,min=0,max=2000"`
	Type        string              `json:"type" validate:"trim,name,min=0,max=50"`
	UserUuid    *openapi_types.UUID `json:"user_uuid,omitempty"`
}

type StatusRequest struct {
	Comment string `json:"comment" validate:"trim,min=0,max=300"`
	Status  int    `json:"status" validate:"gte=0,lte=20"`
}

type UUIDResponse struct {
	Uuid openapi_types.UUID `json:"uuid"`
}

type EntityUUID = openapi_types.UUID

type Uuid = openapi_types.UUID

type PostReminderJSONRequestBody = ReminderCreateRequest

type PutReminderUUIDJSONRequestBody = ReminderPutRequest

type PatchReminderUUIDStatusJSONRequestBody = StatusRequest

type ServerInterface interface {
	GetReminder(ctx echo.Context) error

	PostReminder(ctx echo.Context) error

	DeleteReminderUUID(ctx echo.Context, uUID Uuid) error

	PutReminderUUID(ctx echo.Context, uUID Uuid) error

	PatchReminderUUIDStatus(ctx echo.Context, uUID Uuid) error
}

type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

func (w *ServerInterfaceWrapper) GetReminder(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetReminder(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostReminder(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostReminder(ctx)
	return err
}

func (w *ServerInterfaceWrapper) DeleteReminderUUID(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.DeleteReminderUUID(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PutReminderUUID(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PutReminderUUID(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchReminderUUIDStatus(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchReminderUUIDStatus(ctx, uUID)
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

	router.GET(baseURL+"/reminder", wrapper.GetReminder)
	router.POST(baseURL+"/reminder", wrapper.PostReminder)
	router.DELETE(baseURL+"/reminder/:UUID", wrapper.DeleteReminderUUID)
	router.PUT(baseURL+"/reminder/:UUID", wrapper.PutReminderUUID)
	router.PATCH(baseURL+"/reminder/:UUID/status", wrapper.PatchReminderUUIDStatus)

}

type GetReminderRequestObject struct {
}

type GetReminderResponseObject interface {
	VisitGetReminderResponse(w http.ResponseWriter) error
}

type GetReminder200JSONResponse struct {
	Count int           `json:"count"`
	Items []ReminderDTO `json:"items"`
}

func (response GetReminder200JSONResponse) VisitGetReminderResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostReminderRequestObject struct {
	Body *PostReminderJSONRequestBody
}

type PostReminderResponseObject interface {
	VisitPostReminderResponse(w http.ResponseWriter) error
}

type PostReminder200JSONResponse UUIDResponse

func (response PostReminder200JSONResponse) VisitPostReminderResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteReminderUUIDRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type DeleteReminderUUIDResponseObject interface {
	VisitDeleteReminderUUIDResponse(w http.ResponseWriter) error
}

type DeleteReminderUUID200Response struct {
}

func (response DeleteReminderUUID200Response) VisitDeleteReminderUUIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PutReminderUUIDRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PutReminderUUIDJSONRequestBody
}

type PutReminderUUIDResponseObject interface {
	VisitPutReminderUUIDResponse(w http.ResponseWriter) error
}

type PutReminderUUID200Response struct {
}

func (response PutReminderUUID200Response) VisitPutReminderUUIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchReminderUUIDStatusRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PatchReminderUUIDStatusJSONRequestBody
}

type PatchReminderUUIDStatusResponseObject interface {
	VisitPatchReminderUUIDStatusResponse(w http.ResponseWriter) error
}

type PatchReminderUUIDStatus200Response struct {
}

func (response PatchReminderUUIDStatus200Response) VisitPatchReminderUUIDStatusResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type StrictServerInterface interface {
	GetReminder(ctx context.Context, request GetReminderRequestObject) (GetReminderResponseObject, error)

	PostReminder(ctx context.Context, request PostReminderRequestObject) (PostReminderResponseObject, error)

	DeleteReminderUUID(ctx context.Context, request DeleteReminderUUIDRequestObject) (DeleteReminderUUIDResponseObject, error)

	PutReminderUUID(ctx context.Context, request PutReminderUUIDRequestObject) (PutReminderUUIDResponseObject, error)

	PatchReminderUUIDStatus(ctx context.Context, request PatchReminderUUIDStatusRequestObject) (PatchReminderUUIDStatusResponseObject, error)
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

func (sh *strictHandler) GetReminder(ctx echo.Context) error {
	var request GetReminderRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetReminder(ctx.Request().Context(), request.(GetReminderRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetReminder")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetReminderResponseObject); ok {
		return validResponse.VisitGetReminderResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostReminder(ctx echo.Context) error {
	var request PostReminderRequestObject

	var body PostReminderJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostReminder(ctx.Request().Context(), request.(PostReminderRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostReminder")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostReminderResponseObject); ok {
		return validResponse.VisitPostReminderResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) DeleteReminderUUID(ctx echo.Context, uUID Uuid) error {
	var request DeleteReminderUUIDRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteReminderUUID(ctx.Request().Context(), request.(DeleteReminderUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteReminderUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteReminderUUIDResponseObject); ok {
		return validResponse.VisitDeleteReminderUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PutReminderUUID(ctx echo.Context, uUID Uuid) error {
	var request PutReminderUUIDRequestObject

	request.UUID = uUID

	var body PutReminderUUIDJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutReminderUUID(ctx.Request().Context(), request.(PutReminderUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutReminderUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PutReminderUUIDResponseObject); ok {
		return validResponse.VisitPutReminderUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchReminderUUIDStatus(ctx echo.Context, uUID Uuid) error {
	var request PatchReminderUUIDStatusRequestObject

	request.UUID = uUID

	var body PatchReminderUUIDStatusJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchReminderUUIDStatus(ctx.Request().Context(), request.(PatchReminderUUIDStatusRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchReminderUUIDStatus")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchReminderUUIDStatusResponseObject); ok {
		return validResponse.VisitPatchReminderUUIDStatusResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
