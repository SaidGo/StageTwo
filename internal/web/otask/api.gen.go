package otask

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"example.com/local/Go2part/domain"
	"example.com/local/Go2part/dto"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

type ActivityDTO = dto.ActivityDTO

type CommentDTO = dto.CommentDTO

type NameRequest struct {
	Name string `json:"name" validate:"trim,name,min=0,max=100"`
}

type StatusRequest struct {
	Comment string `json:"comment" validate:"trim,min=0,max=300"`
	Status  int    `json:"status" validate:"gte=0,lte=20"`
}

type TaskCreateRequest struct {
	CoworkersBy   []string               `json:"coworkers_by" validate:"dive,email"`
	Description   string                 `json:"description" validate:"trim,max=5000"`
	Fields        map[string]interface{} `json:"fields"`
	FinishTo      *time.Time             `json:"finish_to,omitempty"`
	Icon          string                 `json:"icon" validate:"trim,max=50"`
	ImplementBy   string                 `json:"implement_by" validate:"omitempty,email"`
	ManagedBy     string                 `json:"managed_by" validate:"omitempty,email"`
	Name          string                 `json:"name" validate:"trim,name,min=3,max=200"`
	Path          []string               `json:"path" validate:"dive,uuid"`
	Priority      int                    `json:"priority" validate:"gte=0,lte=30"`
	ProjectUuid   openapi_types.UUID     `json:"project_uuid" validate:"uuid"`
	ResponsibleBy string                 `json:"responsible_by" validate:"omitempty,email"`
	Tags          []string               `json:"tags" validate:"dive,trim,name,max=40"`
	TaskEntities  []domain.TaskEntity    `json:"task_entities"`
}

type TaskDTO = dto.TaskDTO

type TaskDTOs = dto.TaskDTOs

type TaskPutRequest struct {
	Description *string                 `json:"description,omitempty" validate:"trim,max=5000"`
	Fields      *map[string]interface{} `json:"fields,omitempty"`
	FinishTo    *time.Time              `json:"finish_to,omitempty"`
	Icon        *string                 `json:"icon,omitempty" validate:"omitempty,trim,lte=20"`
	ManagedBy   *string                 `json:"managed_by,omitempty" validate:"omitempty,email"`
	Priority    *int                    `json:"priority,omitempty" validate:"gte=0,lte=30"`
	Tags        *[]string               `json:"tags,omitempty" validate:"dive,trim,name,max=40"`
}

type UploadDTO = dto.UploadDTO

type UserDTO = dto.UserDTO

type EntityUUID = openapi_types.UUID

type FileUUID = openapi_types.UUID

type Uuid = openapi_types.UUID

type GetTaskParams struct {
	Offset         *int               `form:"offset,omitempty" json:"offset,omitempty"`
	Limit          *int               `form:"limit,omitempty" json:"limit,omitempty"`
	IsMy           *bool              `form:"is_my,omitempty" json:"is_my,omitempty"`
	Status         *int               `form:"status,omitempty" json:"status,omitempty"`
	IsEpic         *bool              `form:"is_epic,omitempty" json:"is_epic,omitempty"`
	ProjectUuid    openapi_types.UUID `form:"project_uuid" json:"project_uuid"`
	FederationUuid openapi_types.UUID `form:"federation_uuid" json:"federation_uuid"`
	Participated   *[]string          `form:"participated,omitempty" json:"participated,omitempty"`
	Tags           *[]string          `form:"tags,omitempty" json:"tags,omitempty"`
	Path           *string            `form:"path,omitempty" json:"path,omitempty"`
	Name           *string            `form:"name,omitempty" json:"name,omitempty"`
	Fields         *string            `form:"fields,omitempty" json:"fields,omitempty"`
	Order          *string            `form:"order,omitempty" json:"order,omitempty"`
	By             *string            `form:"by,omitempty" json:"by,omitempty"`
	Format         *string            `form:"format,omitempty" json:"format,omitempty"`
}

type GetTaskUUIDActivityParams struct {
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
}

type PostTaskUUIDCommentMultipartBody struct {
	Comment   *string             `json:"comment,omitempty"`
	File      *openapi_types.File `json:"file,omitempty"`
	People    *[]string           `json:"people,omitempty"`
	ReplyUuid *openapi_types.UUID `json:"reply_uuid,omitempty"`
}

type PatchTaskUUIDCommentEntityUUIDMultipartBody struct {
	Comment   *string             `json:"comment,omitempty"`
	File      *openapi_types.File `json:"file,omitempty"`
	People    *[]string           `json:"people,omitempty"`
	ReplyUuid *openapi_types.UUID `json:"reply_uuid,omitempty"`
}

type PatchTaskUUIDParentJSONBody struct {
	Uuid *openapi_types.UUID `json:"uuid,omitempty" validate:"omitempty,uuid"`
}

type PatchTaskUUIDProjectJSONBody struct {
	Comment string             `json:"comment" validate:"trim,min=0,max=300"`
	Status  int                `json:"status" validate:"gte=0,lte=20"`
	Uuid    openapi_types.UUID `json:"uuid" validate:"uuid"`
}

type PatchTaskUUIDTeamJSONBody struct {
	CoworkersBy   *[]string `json:"coworkers_by,omitempty" validate:"omitempty,dive,email"`
	ImplementBy   *string   `json:"implement_by,omitempty" validate:"omitempty,optional_email"`
	ManagedBy     *string   `json:"managed_by,omitempty" validate:"omitempty,optional_email"`
	ResponsibleBy *string   `json:"responsible_by,omitempty" validate:"omitempty,optional_email"`
	WatchedBy     *[]string `json:"watched_by,omitempty" validate:"omitempty,dive,email"`
}

type PatchTaskUUIDUploadMultipartBody struct {
	File *openapi_types.File `json:"file,omitempty"`
}

type PostTaskUUIDUploadEntityUUIDRenameJSONBody struct {
	Name string `json:"name" validate:"trim,min=1,max=50"`
}

type PostTaskJSONRequestBody = TaskCreateRequest

type PutTaskUUIDJSONRequestBody = TaskPutRequest

type PostTaskUUIDCommentMultipartRequestBody PostTaskUUIDCommentMultipartBody

type PatchTaskUUIDCommentEntityUUIDMultipartRequestBody PatchTaskUUIDCommentEntityUUIDMultipartBody

type PatchTaskUUIDNameJSONRequestBody = NameRequest

type PatchTaskUUIDParentJSONRequestBody PatchTaskUUIDParentJSONBody

type PatchTaskUUIDProjectJSONRequestBody PatchTaskUUIDProjectJSONBody

type PatchTaskUUIDStatusJSONRequestBody = StatusRequest

type PatchTaskUUIDTeamJSONRequestBody PatchTaskUUIDTeamJSONBody

type PatchTaskUUIDUploadMultipartRequestBody PatchTaskUUIDUploadMultipartBody

type PostTaskUUIDUploadEntityUUIDRenameJSONRequestBody PostTaskUUIDUploadEntityUUIDRenameJSONBody

type ServerInterface interface {
	GetTask(ctx echo.Context, params GetTaskParams) error

	PostTask(ctx echo.Context) error

	DeleteTaskUUID(ctx echo.Context, uUID Uuid) error

	GetTaskUUID(ctx echo.Context, uUID Uuid) error

	PutTaskUUID(ctx echo.Context, uUID Uuid) error

	GetTaskUUIDActivity(ctx echo.Context, uUID Uuid, params GetTaskUUIDActivityParams) error

	GetTaskUUIDComment(ctx echo.Context, uUID Uuid) error

	PostTaskUUIDComment(ctx echo.Context, uUID Uuid) error

	DeleteTaskUUIDCommentEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error

	PatchTaskUUIDCommentEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error

	DeleteTaskUUIDCommentEntityUUIDFileFileUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID, fileUUID FileUUID) error

	PatchTaskUUIDCommentEntityUUIDLike(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error

	PatchTaskUUIDCommentEntityUUIDPin(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error

	PatchTaskUUIDName(ctx echo.Context, uUID Uuid) error

	PatchTaskUUIDParent(ctx echo.Context, uUID Uuid) error

	PatchTaskUUIDProject(ctx echo.Context, uUID Uuid) error

	PatchTaskUUIDStatus(ctx echo.Context, uUID Uuid) error

	DeleteTaskUUIDStopEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error

	PatchTaskUUIDTeam(ctx echo.Context, uUID Uuid) error

	GetTaskUUIDUpload(ctx echo.Context, uUID Uuid) error

	PatchTaskUUIDUpload(ctx echo.Context, uUID Uuid) error

	DeleteTaskUUIDUploadEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error

	GetTaskUUIDUploadEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error

	PostTaskUUIDUploadEntityUUIDRename(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error
}

type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

func (w *ServerInterfaceWrapper) GetTask(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	var params GetTaskParams

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "is_my", ctx.QueryParams(), &params.IsMy)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter is_my: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "is_epic", ctx.QueryParams(), &params.IsEpic)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter is_epic: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, true, "project_uuid", ctx.QueryParams(), &params.ProjectUuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter project_uuid: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, true, "federation_uuid", ctx.QueryParams(), &params.FederationUuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter federation_uuid: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "participated", ctx.QueryParams(), &params.Participated)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participated: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "path", ctx.QueryParams(), &params.Path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter path: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "fields", ctx.QueryParams(), &params.Fields)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fields: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "order", ctx.QueryParams(), &params.Order)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter order: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "by", ctx.QueryParams(), &params.By)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter by: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	err = w.Handler.GetTask(ctx, params)
	return err
}

func (w *ServerInterfaceWrapper) PostTask(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostTask(ctx)
	return err
}

func (w *ServerInterfaceWrapper) DeleteTaskUUID(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.DeleteTaskUUID(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) GetTaskUUID(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetTaskUUID(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PutTaskUUID(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PutTaskUUID(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) GetTaskUUIDActivity(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	var params GetTaskUUIDActivityParams

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	err = w.Handler.GetTaskUUIDActivity(ctx, uUID, params)
	return err
}

func (w *ServerInterfaceWrapper) GetTaskUUIDComment(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetTaskUUIDComment(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PostTaskUUIDComment(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostTaskUUIDComment(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) DeleteTaskUUIDCommentEntityUUID(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	var entityUUID EntityUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "entityUUID", runtime.ParamLocationPath, ctx.Param("entityUUID"), &entityUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter entityUUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.DeleteTaskUUIDCommentEntityUUID(ctx, uUID, entityUUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchTaskUUIDCommentEntityUUID(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	var entityUUID EntityUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "entityUUID", runtime.ParamLocationPath, ctx.Param("entityUUID"), &entityUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter entityUUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchTaskUUIDCommentEntityUUID(ctx, uUID, entityUUID)
	return err
}

func (w *ServerInterfaceWrapper) DeleteTaskUUIDCommentEntityUUIDFileFileUUID(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	var entityUUID EntityUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "entityUUID", runtime.ParamLocationPath, ctx.Param("entityUUID"), &entityUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter entityUUID: %s", err))
	}

	var fileUUID FileUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "fileUUID", runtime.ParamLocationPath, ctx.Param("fileUUID"), &fileUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fileUUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.DeleteTaskUUIDCommentEntityUUIDFileFileUUID(ctx, uUID, entityUUID, fileUUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchTaskUUIDCommentEntityUUIDLike(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	var entityUUID EntityUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "entityUUID", runtime.ParamLocationPath, ctx.Param("entityUUID"), &entityUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter entityUUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchTaskUUIDCommentEntityUUIDLike(ctx, uUID, entityUUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchTaskUUIDCommentEntityUUIDPin(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	var entityUUID EntityUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "entityUUID", runtime.ParamLocationPath, ctx.Param("entityUUID"), &entityUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter entityUUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchTaskUUIDCommentEntityUUIDPin(ctx, uUID, entityUUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchTaskUUIDName(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchTaskUUIDName(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchTaskUUIDParent(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchTaskUUIDParent(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchTaskUUIDProject(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchTaskUUIDProject(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchTaskUUIDStatus(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchTaskUUIDStatus(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) DeleteTaskUUIDStopEntityUUID(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	var entityUUID EntityUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "entityUUID", runtime.ParamLocationPath, ctx.Param("entityUUID"), &entityUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter entityUUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.DeleteTaskUUIDStopEntityUUID(ctx, uUID, entityUUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchTaskUUIDTeam(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchTaskUUIDTeam(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) GetTaskUUIDUpload(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetTaskUUIDUpload(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchTaskUUIDUpload(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchTaskUUIDUpload(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) DeleteTaskUUIDUploadEntityUUID(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	var entityUUID EntityUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "entityUUID", runtime.ParamLocationPath, ctx.Param("entityUUID"), &entityUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter entityUUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.DeleteTaskUUIDUploadEntityUUID(ctx, uUID, entityUUID)
	return err
}

func (w *ServerInterfaceWrapper) GetTaskUUIDUploadEntityUUID(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	var entityUUID EntityUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "entityUUID", runtime.ParamLocationPath, ctx.Param("entityUUID"), &entityUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter entityUUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetTaskUUIDUploadEntityUUID(ctx, uUID, entityUUID)
	return err
}

func (w *ServerInterfaceWrapper) PostTaskUUIDUploadEntityUUIDRename(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	var entityUUID EntityUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "entityUUID", runtime.ParamLocationPath, ctx.Param("entityUUID"), &entityUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter entityUUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostTaskUUIDUploadEntityUUIDRename(ctx, uUID, entityUUID)
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

	router.GET(baseURL+"/task", wrapper.GetTask)
	router.POST(baseURL+"/task", wrapper.PostTask)
	router.DELETE(baseURL+"/task/:UUID", wrapper.DeleteTaskUUID)
	router.GET(baseURL+"/task/:UUID", wrapper.GetTaskUUID)
	router.PUT(baseURL+"/task/:UUID", wrapper.PutTaskUUID)
	router.GET(baseURL+"/task/:UUID/activity", wrapper.GetTaskUUIDActivity)
	router.GET(baseURL+"/task/:UUID/comment", wrapper.GetTaskUUIDComment)
	router.POST(baseURL+"/task/:UUID/comment", wrapper.PostTaskUUIDComment)
	router.DELETE(baseURL+"/task/:UUID/comment/:entityUUID", wrapper.DeleteTaskUUIDCommentEntityUUID)
	router.PATCH(baseURL+"/task/:UUID/comment/:entityUUID", wrapper.PatchTaskUUIDCommentEntityUUID)
	router.DELETE(baseURL+"/task/:UUID/comment/:entityUUID/file/:fileUUID", wrapper.DeleteTaskUUIDCommentEntityUUIDFileFileUUID)
	router.PATCH(baseURL+"/task/:UUID/comment/:entityUUID/like", wrapper.PatchTaskUUIDCommentEntityUUIDLike)
	router.PATCH(baseURL+"/task/:UUID/comment/:entityUUID/pin", wrapper.PatchTaskUUIDCommentEntityUUIDPin)
	router.PATCH(baseURL+"/task/:UUID/name", wrapper.PatchTaskUUIDName)
	router.PATCH(baseURL+"/task/:UUID/parent", wrapper.PatchTaskUUIDParent)
	router.PATCH(baseURL+"/task/:UUID/project", wrapper.PatchTaskUUIDProject)
	router.PATCH(baseURL+"/task/:UUID/status", wrapper.PatchTaskUUIDStatus)
	router.DELETE(baseURL+"/task/:UUID/stop/:entityUUID", wrapper.DeleteTaskUUIDStopEntityUUID)
	router.PATCH(baseURL+"/task/:UUID/team", wrapper.PatchTaskUUIDTeam)
	router.GET(baseURL+"/task/:UUID/upload", wrapper.GetTaskUUIDUpload)
	router.PATCH(baseURL+"/task/:UUID/upload", wrapper.PatchTaskUUIDUpload)
	router.DELETE(baseURL+"/task/:UUID/upload/:entityUUID", wrapper.DeleteTaskUUIDUploadEntityUUID)
	router.GET(baseURL+"/task/:UUID/upload/:entityUUID", wrapper.GetTaskUUIDUploadEntityUUID)
	router.POST(baseURL+"/task/:UUID/upload/:entityUUID/rename", wrapper.PostTaskUUIDUploadEntityUUIDRename)

}

type GetTaskRequestObject struct {
	Params GetTaskParams
}

type GetTaskResponseObject interface {
	VisitGetTaskResponse(w http.ResponseWriter) error
}

type GetTask200ResponseHeaders struct {
	ContentDisposition string
	ContentType        string
	CacheControl       string
}

type GetTask200JSONResponse struct {
	Body struct {
		Count int        `json:"count"`
		Items []TaskDTOs `json:"items"`
		Total int64      `json:"total"`
	}
	Headers GetTask200ResponseHeaders
}

func (response GetTask200JSONResponse) VisitGetTaskResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", fmt.Sprint(response.Headers.ContentDisposition))
	w.Header().Set("Content-Type", fmt.Sprint(response.Headers.ContentType))
	w.Header().Set("cache-control", fmt.Sprint(response.Headers.CacheControl))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type GetTask200ApplicationxlsxResponse struct {
	Body          io.Reader
	Headers       GetTask200ResponseHeaders
	ContentLength int64
}

func (response GetTask200ApplicationxlsxResponse) VisitGetTaskResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/xlsx")
	if response.ContentLength != 0 {
		w.Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	w.Header().Set("Content-Disposition", fmt.Sprint(response.Headers.ContentDisposition))
	w.Header().Set("Content-Type", fmt.Sprint(response.Headers.ContentType))
	w.Header().Set("cache-control", fmt.Sprint(response.Headers.CacheControl))
	w.WriteHeader(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(w, response.Body)
	return err
}

type PostTaskRequestObject struct {
	Body *PostTaskJSONRequestBody
}

type PostTaskResponseObject interface {
	VisitPostTaskResponse(w http.ResponseWriter) error
}

type PostTask200JSONResponse struct {
	Id   int                `json:"id"`
	Uuid openapi_types.UUID `json:"uuid"`
}

func (response PostTask200JSONResponse) VisitPostTaskResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTaskUUIDRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type DeleteTaskUUIDResponseObject interface {
	VisitDeleteTaskUUIDResponse(w http.ResponseWriter) error
}

type DeleteTaskUUID200Response struct {
}

func (response DeleteTaskUUID200Response) VisitDeleteTaskUUIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type GetTaskUUIDRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type GetTaskUUIDResponseObject interface {
	VisitGetTaskUUIDResponse(w http.ResponseWriter) error
}

type GetTaskUUID200ResponseHeaders struct {
	CacheControl string
}

type GetTaskUUID200JSONResponse struct {
	Body    TaskDTO
	Headers GetTaskUUID200ResponseHeaders
}

func (response GetTaskUUID200JSONResponse) VisitGetTaskUUIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("cache-control", fmt.Sprint(response.Headers.CacheControl))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type PutTaskUUIDRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PutTaskUUIDJSONRequestBody
}

type PutTaskUUIDResponseObject interface {
	VisitPutTaskUUIDResponse(w http.ResponseWriter) error
}

type PutTaskUUID200Response struct {
}

func (response PutTaskUUID200Response) VisitPutTaskUUIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type GetTaskUUIDActivityRequestObject struct {
	UUID   Uuid `json:"UUID"`
	Params GetTaskUUIDActivityParams
}

type GetTaskUUIDActivityResponseObject interface {
	VisitGetTaskUUIDActivityResponse(w http.ResponseWriter) error
}

type GetTaskUUIDActivity200JSONResponse struct {
	Count int           `json:"count"`
	Items []ActivityDTO `json:"items"`
	Total int64         `json:"total"`
}

func (response GetTaskUUIDActivity200JSONResponse) VisitGetTaskUUIDActivityResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTaskUUIDCommentRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type GetTaskUUIDCommentResponseObject interface {
	VisitGetTaskUUIDCommentResponse(w http.ResponseWriter) error
}

type GetTaskUUIDComment200JSONResponse struct {
	Count int          `json:"count"`
	Items []CommentDTO `json:"items"`
}

func (response GetTaskUUIDComment200JSONResponse) VisitGetTaskUUIDCommentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTaskUUIDCommentRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *multipart.Reader
}

type PostTaskUUIDCommentResponseObject interface {
	VisitPostTaskUUIDCommentResponse(w http.ResponseWriter) error
}

type PostTaskUUIDComment200JSONResponse struct {
	Comment      string             `json:"comment"`
	People       *[]UserDTO         `json:"people,omitempty"`
	ReplyMessage *string            `json:"reply_message,omitempty"`
	Uploads      *[]UploadDTO       `json:"uploads,omitempty"`
	Uuid         openapi_types.UUID `json:"uuid"`
}

func (response PostTaskUUIDComment200JSONResponse) VisitPostTaskUUIDCommentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTaskUUIDCommentEntityUUIDRequestObject struct {
	UUID       Uuid       `json:"UUID"`
	EntityUUID EntityUUID `json:"entityUUID"`
}

type DeleteTaskUUIDCommentEntityUUIDResponseObject interface {
	VisitDeleteTaskUUIDCommentEntityUUIDResponse(w http.ResponseWriter) error
}

type DeleteTaskUUIDCommentEntityUUID200Response struct {
}

func (response DeleteTaskUUIDCommentEntityUUID200Response) VisitDeleteTaskUUIDCommentEntityUUIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchTaskUUIDCommentEntityUUIDRequestObject struct {
	UUID       Uuid       `json:"UUID"`
	EntityUUID EntityUUID `json:"entityUUID"`
	Body       *multipart.Reader
}

type PatchTaskUUIDCommentEntityUUIDResponseObject interface {
	VisitPatchTaskUUIDCommentEntityUUIDResponse(w http.ResponseWriter) error
}

type PatchTaskUUIDCommentEntityUUID200JSONResponse struct {
	Comment      string             `json:"comment"`
	People       *[]UserDTO         `json:"people,omitempty"`
	ReplyMessage *string            `json:"reply_message,omitempty"`
	Uploads      *[]UploadDTO       `json:"uploads,omitempty"`
	Uuid         openapi_types.UUID `json:"uuid"`
}

func (response PatchTaskUUIDCommentEntityUUID200JSONResponse) VisitPatchTaskUUIDCommentEntityUUIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTaskUUIDCommentEntityUUIDFileFileUUIDRequestObject struct {
	UUID       Uuid       `json:"UUID"`
	EntityUUID EntityUUID `json:"entityUUID"`
	FileUUID   FileUUID   `json:"fileUUID"`
}

type DeleteTaskUUIDCommentEntityUUIDFileFileUUIDResponseObject interface {
	VisitDeleteTaskUUIDCommentEntityUUIDFileFileUUIDResponse(w http.ResponseWriter) error
}

type DeleteTaskUUIDCommentEntityUUIDFileFileUUID200Response struct {
}

func (response DeleteTaskUUIDCommentEntityUUIDFileFileUUID200Response) VisitDeleteTaskUUIDCommentEntityUUIDFileFileUUIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchTaskUUIDCommentEntityUUIDLikeRequestObject struct {
	UUID       Uuid       `json:"UUID"`
	EntityUUID EntityUUID `json:"entityUUID"`
}

type PatchTaskUUIDCommentEntityUUIDLikeResponseObject interface {
	VisitPatchTaskUUIDCommentEntityUUIDLikeResponse(w http.ResponseWriter) error
}

type PatchTaskUUIDCommentEntityUUIDLike200JSONResponse struct {
	Liked bool      `json:"liked"`
	Likes []UserDTO `json:"likes"`
}

func (response PatchTaskUUIDCommentEntityUUIDLike200JSONResponse) VisitPatchTaskUUIDCommentEntityUUIDLikeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchTaskUUIDCommentEntityUUIDPinRequestObject struct {
	UUID       Uuid       `json:"UUID"`
	EntityUUID EntityUUID `json:"entityUUID"`
}

type PatchTaskUUIDCommentEntityUUIDPinResponseObject interface {
	VisitPatchTaskUUIDCommentEntityUUIDPinResponse(w http.ResponseWriter) error
}

type PatchTaskUUIDCommentEntityUUIDPin200Response struct {
}

func (response PatchTaskUUIDCommentEntityUUIDPin200Response) VisitPatchTaskUUIDCommentEntityUUIDPinResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchTaskUUIDNameRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PatchTaskUUIDNameJSONRequestBody
}

type PatchTaskUUIDNameResponseObject interface {
	VisitPatchTaskUUIDNameResponse(w http.ResponseWriter) error
}

type PatchTaskUUIDName200Response struct {
}

func (response PatchTaskUUIDName200Response) VisitPatchTaskUUIDNameResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchTaskUUIDName400Response struct {
}

func (response PatchTaskUUIDName400Response) VisitPatchTaskUUIDNameResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PatchTaskUUIDName401Response struct {
}

func (response PatchTaskUUIDName401Response) VisitPatchTaskUUIDNameResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type PatchTaskUUIDParentRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PatchTaskUUIDParentJSONRequestBody
}

type PatchTaskUUIDParentResponseObject interface {
	VisitPatchTaskUUIDParentResponse(w http.ResponseWriter) error
}

type PatchTaskUUIDParent200Response struct {
}

func (response PatchTaskUUIDParent200Response) VisitPatchTaskUUIDParentResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchTaskUUIDProjectRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PatchTaskUUIDProjectJSONRequestBody
}

type PatchTaskUUIDProjectResponseObject interface {
	VisitPatchTaskUUIDProjectResponse(w http.ResponseWriter) error
}

type PatchTaskUUIDProject200Response struct {
}

func (response PatchTaskUUIDProject200Response) VisitPatchTaskUUIDProjectResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchTaskUUIDStatusRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PatchTaskUUIDStatusJSONRequestBody
}

type PatchTaskUUIDStatusResponseObject interface {
	VisitPatchTaskUUIDStatusResponse(w http.ResponseWriter) error
}

type PatchTaskUUIDStatus200JSONResponse struct {
	Path     []string           `json:"path"`
	StopUuid openapi_types.UUID `json:"stop_uuid"`
}

func (response PatchTaskUUIDStatus200JSONResponse) VisitPatchTaskUUIDStatusResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTaskUUIDStopEntityUUIDRequestObject struct {
	UUID       Uuid       `json:"UUID"`
	EntityUUID EntityUUID `json:"entityUUID"`
}

type DeleteTaskUUIDStopEntityUUIDResponseObject interface {
	VisitDeleteTaskUUIDStopEntityUUIDResponse(w http.ResponseWriter) error
}

type DeleteTaskUUIDStopEntityUUID200Response struct {
}

func (response DeleteTaskUUIDStopEntityUUID200Response) VisitDeleteTaskUUIDStopEntityUUIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchTaskUUIDTeamRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PatchTaskUUIDTeamJSONRequestBody
}

type PatchTaskUUIDTeamResponseObject interface {
	VisitPatchTaskUUIDTeamResponse(w http.ResponseWriter) error
}

type PatchTaskUUIDTeam200JSONResponse struct {
	CoworkersBy   []UserDTO `json:"coworkers_by"`
	ImplementBy   *UserDTO  `json:"implement_by,omitempty"`
	ManagedBy     *UserDTO  `json:"managed_by,omitempty"`
	ResponsibleBy *UserDTO  `json:"responsible_by,omitempty"`
	WatchedBy     []UserDTO `json:"watched_by"`
}

func (response PatchTaskUUIDTeam200JSONResponse) VisitPatchTaskUUIDTeamResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTaskUUIDUploadRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type GetTaskUUIDUploadResponseObject interface {
	VisitGetTaskUUIDUploadResponse(w http.ResponseWriter) error
}

type GetTaskUUIDUpload200JSONResponse struct {
	Count int         `json:"count"`
	Items []UploadDTO `json:"items"`
}

func (response GetTaskUUIDUpload200JSONResponse) VisitGetTaskUUIDUploadResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchTaskUUIDUploadRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *multipart.Reader
}

type PatchTaskUUIDUploadResponseObject interface {
	VisitPatchTaskUUIDUploadResponse(w http.ResponseWriter) error
}

type PatchTaskUUIDUpload200JSONResponse UploadDTO

func (response PatchTaskUUIDUpload200JSONResponse) VisitPatchTaskUUIDUploadResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTaskUUIDUploadEntityUUIDRequestObject struct {
	UUID       Uuid       `json:"UUID"`
	EntityUUID EntityUUID `json:"entityUUID"`
}

type DeleteTaskUUIDUploadEntityUUIDResponseObject interface {
	VisitDeleteTaskUUIDUploadEntityUUIDResponse(w http.ResponseWriter) error
}

type DeleteTaskUUIDUploadEntityUUID200Response struct {
}

func (response DeleteTaskUUIDUploadEntityUUID200Response) VisitDeleteTaskUUIDUploadEntityUUIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type GetTaskUUIDUploadEntityUUIDRequestObject struct {
	UUID       Uuid       `json:"UUID"`
	EntityUUID EntityUUID `json:"entityUUID"`
}

type GetTaskUUIDUploadEntityUUIDResponseObject interface {
	VisitGetTaskUUIDUploadEntityUUIDResponse(w http.ResponseWriter) error
}

type GetTaskUUIDUploadEntityUUID302ResponseHeaders struct {
	Location string
}

type GetTaskUUIDUploadEntityUUID302Response struct {
	Headers GetTaskUUIDUploadEntityUUID302ResponseHeaders
}

func (response GetTaskUUIDUploadEntityUUID302Response) VisitGetTaskUUIDUploadEntityUUIDResponse(w http.ResponseWriter) error {
	w.Header().Set("location", fmt.Sprint(response.Headers.Location))
	w.WriteHeader(302)
	return nil
}

type PostTaskUUIDUploadEntityUUIDRenameRequestObject struct {
	UUID       Uuid       `json:"UUID"`
	EntityUUID EntityUUID `json:"entityUUID"`
	Body       *PostTaskUUIDUploadEntityUUIDRenameJSONRequestBody
}

type PostTaskUUIDUploadEntityUUIDRenameResponseObject interface {
	VisitPostTaskUUIDUploadEntityUUIDRenameResponse(w http.ResponseWriter) error
}

type PostTaskUUIDUploadEntityUUIDRename200Response struct {
}

func (response PostTaskUUIDUploadEntityUUIDRename200Response) VisitPostTaskUUIDUploadEntityUUIDRenameResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type StrictServerInterface interface {
	GetTask(ctx context.Context, request GetTaskRequestObject) (GetTaskResponseObject, error)

	PostTask(ctx context.Context, request PostTaskRequestObject) (PostTaskResponseObject, error)

	DeleteTaskUUID(ctx context.Context, request DeleteTaskUUIDRequestObject) (DeleteTaskUUIDResponseObject, error)

	GetTaskUUID(ctx context.Context, request GetTaskUUIDRequestObject) (GetTaskUUIDResponseObject, error)

	PutTaskUUID(ctx context.Context, request PutTaskUUIDRequestObject) (PutTaskUUIDResponseObject, error)

	GetTaskUUIDActivity(ctx context.Context, request GetTaskUUIDActivityRequestObject) (GetTaskUUIDActivityResponseObject, error)

	GetTaskUUIDComment(ctx context.Context, request GetTaskUUIDCommentRequestObject) (GetTaskUUIDCommentResponseObject, error)

	PostTaskUUIDComment(ctx context.Context, request PostTaskUUIDCommentRequestObject) (PostTaskUUIDCommentResponseObject, error)

	DeleteTaskUUIDCommentEntityUUID(ctx context.Context, request DeleteTaskUUIDCommentEntityUUIDRequestObject) (DeleteTaskUUIDCommentEntityUUIDResponseObject, error)

	PatchTaskUUIDCommentEntityUUID(ctx context.Context, request PatchTaskUUIDCommentEntityUUIDRequestObject) (PatchTaskUUIDCommentEntityUUIDResponseObject, error)

	DeleteTaskUUIDCommentEntityUUIDFileFileUUID(ctx context.Context, request DeleteTaskUUIDCommentEntityUUIDFileFileUUIDRequestObject) (DeleteTaskUUIDCommentEntityUUIDFileFileUUIDResponseObject, error)

	PatchTaskUUIDCommentEntityUUIDLike(ctx context.Context, request PatchTaskUUIDCommentEntityUUIDLikeRequestObject) (PatchTaskUUIDCommentEntityUUIDLikeResponseObject, error)

	PatchTaskUUIDCommentEntityUUIDPin(ctx context.Context, request PatchTaskUUIDCommentEntityUUIDPinRequestObject) (PatchTaskUUIDCommentEntityUUIDPinResponseObject, error)

	PatchTaskUUIDName(ctx context.Context, request PatchTaskUUIDNameRequestObject) (PatchTaskUUIDNameResponseObject, error)

	PatchTaskUUIDParent(ctx context.Context, request PatchTaskUUIDParentRequestObject) (PatchTaskUUIDParentResponseObject, error)

	PatchTaskUUIDProject(ctx context.Context, request PatchTaskUUIDProjectRequestObject) (PatchTaskUUIDProjectResponseObject, error)

	PatchTaskUUIDStatus(ctx context.Context, request PatchTaskUUIDStatusRequestObject) (PatchTaskUUIDStatusResponseObject, error)

	DeleteTaskUUIDStopEntityUUID(ctx context.Context, request DeleteTaskUUIDStopEntityUUIDRequestObject) (DeleteTaskUUIDStopEntityUUIDResponseObject, error)

	PatchTaskUUIDTeam(ctx context.Context, request PatchTaskUUIDTeamRequestObject) (PatchTaskUUIDTeamResponseObject, error)

	GetTaskUUIDUpload(ctx context.Context, request GetTaskUUIDUploadRequestObject) (GetTaskUUIDUploadResponseObject, error)

	PatchTaskUUIDUpload(ctx context.Context, request PatchTaskUUIDUploadRequestObject) (PatchTaskUUIDUploadResponseObject, error)

	DeleteTaskUUIDUploadEntityUUID(ctx context.Context, request DeleteTaskUUIDUploadEntityUUIDRequestObject) (DeleteTaskUUIDUploadEntityUUIDResponseObject, error)

	GetTaskUUIDUploadEntityUUID(ctx context.Context, request GetTaskUUIDUploadEntityUUIDRequestObject) (GetTaskUUIDUploadEntityUUIDResponseObject, error)

	PostTaskUUIDUploadEntityUUIDRename(ctx context.Context, request PostTaskUUIDUploadEntityUUIDRenameRequestObject) (PostTaskUUIDUploadEntityUUIDRenameResponseObject, error)
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

func (sh *strictHandler) GetTask(ctx echo.Context, params GetTaskParams) error {
	var request GetTaskRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTask(ctx.Request().Context(), request.(GetTaskRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTask")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTaskResponseObject); ok {
		return validResponse.VisitGetTaskResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostTask(ctx echo.Context) error {
	var request PostTaskRequestObject

	var body PostTaskJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTask(ctx.Request().Context(), request.(PostTaskRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTask")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTaskResponseObject); ok {
		return validResponse.VisitPostTaskResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) DeleteTaskUUID(ctx echo.Context, uUID Uuid) error {
	var request DeleteTaskUUIDRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTaskUUID(ctx.Request().Context(), request.(DeleteTaskUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTaskUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTaskUUIDResponseObject); ok {
		return validResponse.VisitDeleteTaskUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) GetTaskUUID(ctx echo.Context, uUID Uuid) error {
	var request GetTaskUUIDRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTaskUUID(ctx.Request().Context(), request.(GetTaskUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTaskUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTaskUUIDResponseObject); ok {
		return validResponse.VisitGetTaskUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PutTaskUUID(ctx echo.Context, uUID Uuid) error {
	var request PutTaskUUIDRequestObject

	request.UUID = uUID

	var body PutTaskUUIDJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutTaskUUID(ctx.Request().Context(), request.(PutTaskUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutTaskUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PutTaskUUIDResponseObject); ok {
		return validResponse.VisitPutTaskUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) GetTaskUUIDActivity(ctx echo.Context, uUID Uuid, params GetTaskUUIDActivityParams) error {
	var request GetTaskUUIDActivityRequestObject

	request.UUID = uUID
	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTaskUUIDActivity(ctx.Request().Context(), request.(GetTaskUUIDActivityRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTaskUUIDActivity")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTaskUUIDActivityResponseObject); ok {
		return validResponse.VisitGetTaskUUIDActivityResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) GetTaskUUIDComment(ctx echo.Context, uUID Uuid) error {
	var request GetTaskUUIDCommentRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTaskUUIDComment(ctx.Request().Context(), request.(GetTaskUUIDCommentRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTaskUUIDComment")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTaskUUIDCommentResponseObject); ok {
		return validResponse.VisitGetTaskUUIDCommentResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostTaskUUIDComment(ctx echo.Context, uUID Uuid) error {
	var request PostTaskUUIDCommentRequestObject

	request.UUID = uUID

	if reader, err := ctx.Request().MultipartReader(); err != nil {
		return err
	} else {
		request.Body = reader
	}

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTaskUUIDComment(ctx.Request().Context(), request.(PostTaskUUIDCommentRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTaskUUIDComment")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTaskUUIDCommentResponseObject); ok {
		return validResponse.VisitPostTaskUUIDCommentResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) DeleteTaskUUIDCommentEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error {
	var request DeleteTaskUUIDCommentEntityUUIDRequestObject

	request.UUID = uUID
	request.EntityUUID = entityUUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTaskUUIDCommentEntityUUID(ctx.Request().Context(), request.(DeleteTaskUUIDCommentEntityUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTaskUUIDCommentEntityUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTaskUUIDCommentEntityUUIDResponseObject); ok {
		return validResponse.VisitDeleteTaskUUIDCommentEntityUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchTaskUUIDCommentEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error {
	var request PatchTaskUUIDCommentEntityUUIDRequestObject

	request.UUID = uUID
	request.EntityUUID = entityUUID

	if reader, err := ctx.Request().MultipartReader(); err != nil {
		return err
	} else {
		request.Body = reader
	}

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTaskUUIDCommentEntityUUID(ctx.Request().Context(), request.(PatchTaskUUIDCommentEntityUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTaskUUIDCommentEntityUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTaskUUIDCommentEntityUUIDResponseObject); ok {
		return validResponse.VisitPatchTaskUUIDCommentEntityUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) DeleteTaskUUIDCommentEntityUUIDFileFileUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID, fileUUID FileUUID) error {
	var request DeleteTaskUUIDCommentEntityUUIDFileFileUUIDRequestObject

	request.UUID = uUID
	request.EntityUUID = entityUUID
	request.FileUUID = fileUUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTaskUUIDCommentEntityUUIDFileFileUUID(ctx.Request().Context(), request.(DeleteTaskUUIDCommentEntityUUIDFileFileUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTaskUUIDCommentEntityUUIDFileFileUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTaskUUIDCommentEntityUUIDFileFileUUIDResponseObject); ok {
		return validResponse.VisitDeleteTaskUUIDCommentEntityUUIDFileFileUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchTaskUUIDCommentEntityUUIDLike(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error {
	var request PatchTaskUUIDCommentEntityUUIDLikeRequestObject

	request.UUID = uUID
	request.EntityUUID = entityUUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTaskUUIDCommentEntityUUIDLike(ctx.Request().Context(), request.(PatchTaskUUIDCommentEntityUUIDLikeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTaskUUIDCommentEntityUUIDLike")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTaskUUIDCommentEntityUUIDLikeResponseObject); ok {
		return validResponse.VisitPatchTaskUUIDCommentEntityUUIDLikeResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchTaskUUIDCommentEntityUUIDPin(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error {
	var request PatchTaskUUIDCommentEntityUUIDPinRequestObject

	request.UUID = uUID
	request.EntityUUID = entityUUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTaskUUIDCommentEntityUUIDPin(ctx.Request().Context(), request.(PatchTaskUUIDCommentEntityUUIDPinRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTaskUUIDCommentEntityUUIDPin")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTaskUUIDCommentEntityUUIDPinResponseObject); ok {
		return validResponse.VisitPatchTaskUUIDCommentEntityUUIDPinResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchTaskUUIDName(ctx echo.Context, uUID Uuid) error {
	var request PatchTaskUUIDNameRequestObject

	request.UUID = uUID

	var body PatchTaskUUIDNameJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTaskUUIDName(ctx.Request().Context(), request.(PatchTaskUUIDNameRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTaskUUIDName")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTaskUUIDNameResponseObject); ok {
		return validResponse.VisitPatchTaskUUIDNameResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchTaskUUIDParent(ctx echo.Context, uUID Uuid) error {
	var request PatchTaskUUIDParentRequestObject

	request.UUID = uUID

	var body PatchTaskUUIDParentJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTaskUUIDParent(ctx.Request().Context(), request.(PatchTaskUUIDParentRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTaskUUIDParent")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTaskUUIDParentResponseObject); ok {
		return validResponse.VisitPatchTaskUUIDParentResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchTaskUUIDProject(ctx echo.Context, uUID Uuid) error {
	var request PatchTaskUUIDProjectRequestObject

	request.UUID = uUID

	var body PatchTaskUUIDProjectJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTaskUUIDProject(ctx.Request().Context(), request.(PatchTaskUUIDProjectRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTaskUUIDProject")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTaskUUIDProjectResponseObject); ok {
		return validResponse.VisitPatchTaskUUIDProjectResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchTaskUUIDStatus(ctx echo.Context, uUID Uuid) error {
	var request PatchTaskUUIDStatusRequestObject

	request.UUID = uUID

	var body PatchTaskUUIDStatusJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTaskUUIDStatus(ctx.Request().Context(), request.(PatchTaskUUIDStatusRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTaskUUIDStatus")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTaskUUIDStatusResponseObject); ok {
		return validResponse.VisitPatchTaskUUIDStatusResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) DeleteTaskUUIDStopEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error {
	var request DeleteTaskUUIDStopEntityUUIDRequestObject

	request.UUID = uUID
	request.EntityUUID = entityUUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTaskUUIDStopEntityUUID(ctx.Request().Context(), request.(DeleteTaskUUIDStopEntityUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTaskUUIDStopEntityUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTaskUUIDStopEntityUUIDResponseObject); ok {
		return validResponse.VisitDeleteTaskUUIDStopEntityUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchTaskUUIDTeam(ctx echo.Context, uUID Uuid) error {
	var request PatchTaskUUIDTeamRequestObject

	request.UUID = uUID

	var body PatchTaskUUIDTeamJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTaskUUIDTeam(ctx.Request().Context(), request.(PatchTaskUUIDTeamRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTaskUUIDTeam")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTaskUUIDTeamResponseObject); ok {
		return validResponse.VisitPatchTaskUUIDTeamResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) GetTaskUUIDUpload(ctx echo.Context, uUID Uuid) error {
	var request GetTaskUUIDUploadRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTaskUUIDUpload(ctx.Request().Context(), request.(GetTaskUUIDUploadRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTaskUUIDUpload")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTaskUUIDUploadResponseObject); ok {
		return validResponse.VisitGetTaskUUIDUploadResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchTaskUUIDUpload(ctx echo.Context, uUID Uuid) error {
	var request PatchTaskUUIDUploadRequestObject

	request.UUID = uUID

	if reader, err := ctx.Request().MultipartReader(); err != nil {
		return err
	} else {
		request.Body = reader
	}

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTaskUUIDUpload(ctx.Request().Context(), request.(PatchTaskUUIDUploadRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTaskUUIDUpload")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTaskUUIDUploadResponseObject); ok {
		return validResponse.VisitPatchTaskUUIDUploadResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) DeleteTaskUUIDUploadEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error {
	var request DeleteTaskUUIDUploadEntityUUIDRequestObject

	request.UUID = uUID
	request.EntityUUID = entityUUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTaskUUIDUploadEntityUUID(ctx.Request().Context(), request.(DeleteTaskUUIDUploadEntityUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTaskUUIDUploadEntityUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTaskUUIDUploadEntityUUIDResponseObject); ok {
		return validResponse.VisitDeleteTaskUUIDUploadEntityUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) GetTaskUUIDUploadEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error {
	var request GetTaskUUIDUploadEntityUUIDRequestObject

	request.UUID = uUID
	request.EntityUUID = entityUUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTaskUUIDUploadEntityUUID(ctx.Request().Context(), request.(GetTaskUUIDUploadEntityUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTaskUUIDUploadEntityUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTaskUUIDUploadEntityUUIDResponseObject); ok {
		return validResponse.VisitGetTaskUUIDUploadEntityUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostTaskUUIDUploadEntityUUIDRename(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error {
	var request PostTaskUUIDUploadEntityUUIDRenameRequestObject

	request.UUID = uUID
	request.EntityUUID = entityUUID

	var body PostTaskUUIDUploadEntityUUIDRenameJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTaskUUIDUploadEntityUUIDRename(ctx.Request().Context(), request.(PostTaskUUIDUploadEntityUUIDRenameRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTaskUUIDUploadEntityUUIDRename")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTaskUUIDUploadEntityUUIDRenameResponseObject); ok {
		return validResponse.VisitPostTaskUUIDUploadEntityUUIDRenameResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
