package oprofile

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
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

const (
	PostProfileDislikeJSONBodyTypeCompany    PostProfileDislikeJSONBodyType = "company"
	PostProfileDislikeJSONBodyTypeFederation PostProfileDislikeJSONBodyType = "federation"
	PostProfileDislikeJSONBodyTypeProject    PostProfileDislikeJSONBodyType = "project"
	PostProfileDislikeJSONBodyTypeTask       PostProfileDislikeJSONBodyType = "task"
)

const (
	PostProfileLikeJSONBodyTypeCompany    PostProfileLikeJSONBodyType = "company"
	PostProfileLikeJSONBodyTypeFederation PostProfileLikeJSONBodyType = "federation"
	PostProfileLikeJSONBodyTypeProject    PostProfileLikeJSONBodyType = "project"
	PostProfileLikeJSONBodyTypeTask       PostProfileLikeJSONBodyType = "task"
)

type CompanyDTO = dto.CompanyDTO

type CompanyDTOs = dto.CompanyDTOs

type FederationDTO = dto.FederationDTO

type FederationDTOs = dto.FederationDTOs

type InviteDTO = dto.InviteDTO

type NotificationReminderDTO = dto.NotificationReminderDTO

type NotificationTaskDTO = dto.NotificationTaskDTO

type ProfileChangePasswordRequest struct {
	NewPassword string `json:"new_password" validate:"trim,min=6,max=30"`
	Password    string `json:"password" validate:"trim,min=6,max=30"`
}

type ProfileDTO = dto.ProfileDTO

type ProfileLoginAsRequest struct {
	Email string `json:"email" validate:"trim,min=5,max=200"`
}

type ProfileLoginRequest struct {
	Email      string `json:"email" validate:"trim,min=5,max=200"`
	Password   string `json:"password" validate:"trim,min=6,max=30"`
	RememberMe bool   `json:"remember_me"`
}

type ProfileLoginResponse struct {
	AccessToken  string             `json:"access_token"`
	RefreshToken string             `json:"refresh_token"`
	Uuid         openapi_types.UUID `json:"uuid"`
	ValidUntil   time.Time          `json:"valid_until"`
}

type ProfilePhotoDTO = dto.ProfilePhotoDTO

type ProfileRegisterRequest struct {
	Email    string `json:"email" validate:"trim,min=5,max=200"`
	Lname    string `json:"lname" validate:"trim,name,min=0,max=30"`
	Name     string `json:"name" validate:"trim,name,min=1,max=30"`
	Password string `json:"password" validate:"trim,min=6,max=30"`
	Phone    int    `json:"phone" validate:"min=10000000000,max=9999999999999"`
	Pname    string `json:"pname" validate:"trim,name,min=0,max=30"`
}

type ProfileResetRequest struct {
	Code     string `json:"code" validate:"trim,min=20,max=100"`
	Password string `json:"password" validate:"trim,min=6,max=30"`
}

type ProfileResetSendRequest struct {
	Email string `json:"email" validate:"trim,email"`
}

type ProfileSentValidateRequest struct {
	Email string `json:"email" validate:"email"`
}

type ProfileValidateRequest struct {
	Code string `json:"code" validate:"trim,min=20,max=20"`
}

type ProfileValidateSimpleRequest struct {
	Code  int    `json:"code" validate:"gte=100000,max=999999"`
	Email string `json:"email" validate:"email"`
}

type ProjectDTOs = dto.ProjectDTOs

type UUIDResponse struct {
	Uuid openapi_types.UUID `json:"uuid"`
}

type UserDTO = dto.UserDTO

type EntityUUID = openapi_types.UUID

type Uuid = openapi_types.UUID

type PatchProfileColorJSONBody struct {
	Color string `json:"color" validate:"color"`
}

type PostProfileDislikeJSONBody struct {
	Type PostProfileDislikeJSONBodyType `json:"type"`
	Uuid openapi_types.UUID             `json:"uuid"`
}

type PostProfileDislikeJSONBodyType string

type PatchProfileFioJSONBody struct {
	Lname *string `json:"lname,omitempty" validate:"trim,name,omitempty,max=50"`
	Name  *string `json:"name,omitempty" validate:"trim,name,omitempty,min=3,max=50"`
	Pname *string `json:"pname,omitempty" validate:"trim,name,omitempty,max=50"`
}

type PostProfileLikeJSONBody struct {
	Type PostProfileLikeJSONBodyType `json:"type"`
	Uuid openapi_types.UUID          `json:"uuid"`
}

type PostProfileLikeJSONBodyType string

type PatchProfilePhoneJSONBody struct {
	Phone int `json:"phone" validate:"trim,min=10000000000,max=9999999999999"`
}

type PatchProfilePhotoMultipartBody struct {
	File *openapi_types.File `json:"file,omitempty"`
}

type PatchProfilePreferencesJSONBody struct {
	Timezone *string `json:"timezone,omitempty"`
}

type PostProfileJSONRequestBody = ProfileRegisterRequest

type PatchProfileColorJSONRequestBody PatchProfileColorJSONBody

type PostProfileDislikeJSONRequestBody PostProfileDislikeJSONBody

type PatchProfileFioJSONRequestBody PatchProfileFioJSONBody

type PostProfileLikeJSONRequestBody PostProfileLikeJSONBody

type PostProfileLoginJSONRequestBody = ProfileLoginRequest

type PostProfileLoginAsJSONRequestBody = ProfileLoginAsRequest

type PatchProfilePasswordJSONRequestBody = ProfileChangePasswordRequest

type PatchProfilePhoneJSONRequestBody PatchProfilePhoneJSONBody

type PatchProfilePhotoMultipartRequestBody PatchProfilePhotoMultipartBody

type PatchProfilePreferencesJSONRequestBody PatchProfilePreferencesJSONBody

type PostProfileResetJSONRequestBody = ProfileResetRequest

type PostProfileResetSendJSONRequestBody = ProfileResetSendRequest

type PostProfileValidateJSONRequestBody = ProfileValidateRequest

type PostProfileValidateSimpleJSONRequestBody = ProfileValidateSimpleRequest

type PostProfileValidateSimpleSendJSONRequestBody = ProfileSentValidateRequest

type PostProfileValidateSendJSONRequestBody = ProfileSentValidateRequest

type ServerInterface interface {
	DeleteProfile(ctx echo.Context) error

	GetProfile(ctx echo.Context) error

	PostProfile(ctx echo.Context) error

	PatchProfileColor(ctx echo.Context) error

	PostProfileDislike(ctx echo.Context) error

	PatchProfileFio(ctx echo.Context) error

	GetProfileInvite(ctx echo.Context) error

	PatchProfileInviteUUIDAccept(ctx echo.Context, uUID Uuid) error

	PatchProfileInviteUUIDDecline(ctx echo.Context, uUID Uuid) error

	PostProfileLike(ctx echo.Context) error

	GetProfileLikes(ctx echo.Context) error

	PostProfileLogin(ctx echo.Context) error

	PostProfileLoginAs(ctx echo.Context) error

	GetProfileLogout(ctx echo.Context) error

	DeleteProfileNotifications(ctx echo.Context) error

	GetProfileNotifications(ctx echo.Context) error

	PostProfileNotificationsTaskUUIDHide(ctx echo.Context, uUID Uuid) error

	DeleteProfileNotificationsTaskUUIDStar(ctx echo.Context, uUID Uuid) error

	PostProfileNotificationsTaskUUIDStar(ctx echo.Context, uUID Uuid) error

	PatchProfilePassword(ctx echo.Context) error

	PatchProfilePhone(ctx echo.Context) error

	DeleteProfilePhoto(ctx echo.Context) error

	PatchProfilePhoto(ctx echo.Context) error

	PatchProfilePreferences(ctx echo.Context) error

	PostProfileReset(ctx echo.Context) error

	PostProfileResetSend(ctx echo.Context) error

	PostProfileValidate(ctx echo.Context) error

	PostProfileValidateSimple(ctx echo.Context) error

	PostProfileValidateSimpleSend(ctx echo.Context) error

	PostProfileValidateSend(ctx echo.Context) error
}

type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

func (w *ServerInterfaceWrapper) DeleteProfile(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.DeleteProfile(ctx)
	return err
}

func (w *ServerInterfaceWrapper) GetProfile(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetProfile(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostProfile(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfile(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PatchProfileColor(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchProfileColor(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileDislike(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileDislike(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PatchProfileFio(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchProfileFio(ctx)
	return err
}

func (w *ServerInterfaceWrapper) GetProfileInvite(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetProfileInvite(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PatchProfileInviteUUIDAccept(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchProfileInviteUUIDAccept(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchProfileInviteUUIDDecline(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchProfileInviteUUIDDecline(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileLike(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileLike(ctx)
	return err
}

func (w *ServerInterfaceWrapper) GetProfileLikes(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetProfileLikes(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileLogin(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileLogin(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileLoginAs(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileLoginAs(ctx)
	return err
}

func (w *ServerInterfaceWrapper) GetProfileLogout(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetProfileLogout(ctx)
	return err
}

func (w *ServerInterfaceWrapper) DeleteProfileNotifications(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.DeleteProfileNotifications(ctx)
	return err
}

func (w *ServerInterfaceWrapper) GetProfileNotifications(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.GetProfileNotifications(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileNotificationsTaskUUIDHide(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileNotificationsTaskUUIDHide(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) DeleteProfileNotificationsTaskUUIDStar(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.DeleteProfileNotificationsTaskUUIDStar(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileNotificationsTaskUUIDStar(ctx echo.Context) error {
	var err error

	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileNotificationsTaskUUIDStar(ctx, uUID)
	return err
}

func (w *ServerInterfaceWrapper) PatchProfilePassword(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchProfilePassword(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PatchProfilePhone(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchProfilePhone(ctx)
	return err
}

func (w *ServerInterfaceWrapper) DeleteProfilePhoto(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.DeleteProfilePhoto(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PatchProfilePhoto(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchProfilePhoto(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PatchProfilePreferences(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PatchProfilePreferences(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileReset(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileReset(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileResetSend(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileResetSend(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileValidate(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileValidate(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileValidateSimple(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileValidateSimple(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileValidateSimpleSend(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileValidateSimpleSend(ctx)
	return err
}

func (w *ServerInterfaceWrapper) PostProfileValidateSend(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	err = w.Handler.PostProfileValidateSend(ctx)
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

	router.DELETE(baseURL+"/profile", wrapper.DeleteProfile)
	router.GET(baseURL+"/profile", wrapper.GetProfile)
	router.POST(baseURL+"/profile", wrapper.PostProfile)
	router.PATCH(baseURL+"/profile/color", wrapper.PatchProfileColor)
	router.POST(baseURL+"/profile/dislike", wrapper.PostProfileDislike)
	router.PATCH(baseURL+"/profile/fio", wrapper.PatchProfileFio)
	router.GET(baseURL+"/profile/invite", wrapper.GetProfileInvite)
	router.PATCH(baseURL+"/profile/invite/:UUID/accept", wrapper.PatchProfileInviteUUIDAccept)
	router.PATCH(baseURL+"/profile/invite/:UUID/decline", wrapper.PatchProfileInviteUUIDDecline)
	router.POST(baseURL+"/profile/like", wrapper.PostProfileLike)
	router.GET(baseURL+"/profile/likes", wrapper.GetProfileLikes)
	router.POST(baseURL+"/profile/login", wrapper.PostProfileLogin)
	router.POST(baseURL+"/profile/login_as", wrapper.PostProfileLoginAs)
	router.GET(baseURL+"/profile/logout", wrapper.GetProfileLogout)
	router.DELETE(baseURL+"/profile/notifications", wrapper.DeleteProfileNotifications)
	router.GET(baseURL+"/profile/notifications", wrapper.GetProfileNotifications)
	router.POST(baseURL+"/profile/notifications/task/:UUID/hide", wrapper.PostProfileNotificationsTaskUUIDHide)
	router.DELETE(baseURL+"/profile/notifications/task/:UUID/star", wrapper.DeleteProfileNotificationsTaskUUIDStar)
	router.POST(baseURL+"/profile/notifications/task/:UUID/star", wrapper.PostProfileNotificationsTaskUUIDStar)
	router.PATCH(baseURL+"/profile/password", wrapper.PatchProfilePassword)
	router.PATCH(baseURL+"/profile/phone", wrapper.PatchProfilePhone)
	router.DELETE(baseURL+"/profile/photo", wrapper.DeleteProfilePhoto)
	router.PATCH(baseURL+"/profile/photo", wrapper.PatchProfilePhoto)
	router.PATCH(baseURL+"/profile/preferences", wrapper.PatchProfilePreferences)
	router.POST(baseURL+"/profile/reset", wrapper.PostProfileReset)
	router.POST(baseURL+"/profile/reset/send", wrapper.PostProfileResetSend)
	router.POST(baseURL+"/profile/validate", wrapper.PostProfileValidate)
	router.POST(baseURL+"/profile/validate-simple", wrapper.PostProfileValidateSimple)
	router.POST(baseURL+"/profile/validate-simple/send", wrapper.PostProfileValidateSimpleSend)
	router.POST(baseURL+"/profile/validate/send", wrapper.PostProfileValidateSend)

}

type DeleteProfileRequestObject struct {
}

type DeleteProfileResponseObject interface {
	VisitDeleteProfileResponse(w http.ResponseWriter) error
}

type DeleteProfile200ResponseHeaders struct {
	SetCookie string
}

type DeleteProfile200Response struct {
	Headers DeleteProfile200ResponseHeaders
}

func (response DeleteProfile200Response) VisitDeleteProfileResponse(w http.ResponseWriter) error {
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)
	return nil
}

type GetProfileRequestObject struct {
}

type GetProfileResponseObject interface {
	VisitGetProfileResponse(w http.ResponseWriter) error
}

type GetProfile200JSONResponse ProfileDTO

func (response GetProfile200JSONResponse) VisitGetProfileResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostProfileRequestObject struct {
	Body *PostProfileJSONRequestBody
}

type PostProfileResponseObject interface {
	VisitPostProfileResponse(w http.ResponseWriter) error
}

type PostProfile200ResponseHeaders struct {
	Hint string
}

type PostProfile200JSONResponse struct {
	Body    UUIDResponse
	Headers PostProfile200ResponseHeaders
}

func (response PostProfile200JSONResponse) VisitPostProfileResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Hint", fmt.Sprint(response.Headers.Hint))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type PatchProfileColorRequestObject struct {
	Body *PatchProfileColorJSONRequestBody
}

type PatchProfileColorResponseObject interface {
	VisitPatchProfileColorResponse(w http.ResponseWriter) error
}

type PatchProfileColor200Response struct {
}

func (response PatchProfileColor200Response) VisitPatchProfileColorResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostProfileDislikeRequestObject struct {
	Body *PostProfileDislikeJSONRequestBody
}

type PostProfileDislikeResponseObject interface {
	VisitPostProfileDislikeResponse(w http.ResponseWriter) error
}

type PostProfileDislike200Response struct {
}

func (response PostProfileDislike200Response) VisitPostProfileDislikeResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchProfileFioRequestObject struct {
	Body *PatchProfileFioJSONRequestBody
}

type PatchProfileFioResponseObject interface {
	VisitPatchProfileFioResponse(w http.ResponseWriter) error
}

type PatchProfileFio200Response struct {
}

func (response PatchProfileFio200Response) VisitPatchProfileFioResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type GetProfileInviteRequestObject struct {
}

type GetProfileInviteResponseObject interface {
	VisitGetProfileInviteResponse(w http.ResponseWriter) error
}

type GetProfileInvite200JSONResponse struct {
	Count int         `json:"count"`
	Items []InviteDTO `json:"items"`
}

func (response GetProfileInvite200JSONResponse) VisitGetProfileInviteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchProfileInviteUUIDAcceptRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type PatchProfileInviteUUIDAcceptResponseObject interface {
	VisitPatchProfileInviteUUIDAcceptResponse(w http.ResponseWriter) error
}

type PatchProfileInviteUUIDAccept200Response struct {
}

func (response PatchProfileInviteUUIDAccept200Response) VisitPatchProfileInviteUUIDAcceptResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchProfileInviteUUIDDeclineRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type PatchProfileInviteUUIDDeclineResponseObject interface {
	VisitPatchProfileInviteUUIDDeclineResponse(w http.ResponseWriter) error
}

type PatchProfileInviteUUIDDecline200Response struct {
}

func (response PatchProfileInviteUUIDDecline200Response) VisitPatchProfileInviteUUIDDeclineResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostProfileLikeRequestObject struct {
	Body *PostProfileLikeJSONRequestBody
}

type PostProfileLikeResponseObject interface {
	VisitPostProfileLikeResponse(w http.ResponseWriter) error
}

type PostProfileLike200Response struct {
}

func (response PostProfileLike200Response) VisitPostProfileLikeResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type GetProfileLikesRequestObject struct {
}

type GetProfileLikesResponseObject interface {
	VisitGetProfileLikesResponse(w http.ResponseWriter) error
}

type GetProfileLikes200JSONResponse struct {
	Companies   []openapi_types.UUID `json:"companies"`
	Federations []openapi_types.UUID `json:"federations"`
	Projects    []openapi_types.UUID `json:"projects"`
	Tasks       []openapi_types.UUID `json:"tasks"`
}

func (response GetProfileLikes200JSONResponse) VisitGetProfileLikesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostProfileLoginRequestObject struct {
	Body *PostProfileLoginJSONRequestBody
}

type PostProfileLoginResponseObject interface {
	VisitPostProfileLoginResponse(w http.ResponseWriter) error
}

type PostProfileLogin200ResponseHeaders struct {
	SetCookie string
}

type PostProfileLogin200JSONResponse struct {
	Body    ProfileLoginResponse
	Headers PostProfileLogin200ResponseHeaders
}

func (response PostProfileLogin200JSONResponse) VisitPostProfileLoginResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostProfileLoginAsRequestObject struct {
	Body *PostProfileLoginAsJSONRequestBody
}

type PostProfileLoginAsResponseObject interface {
	VisitPostProfileLoginAsResponse(w http.ResponseWriter) error
}

type PostProfileLoginAs200ResponseHeaders struct {
	SetCookie string
}

type PostProfileLoginAs200JSONResponse struct {
	Body    ProfileLoginResponse
	Headers PostProfileLoginAs200ResponseHeaders
}

func (response PostProfileLoginAs200JSONResponse) VisitPostProfileLoginAsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type GetProfileLogoutRequestObject struct {
}

type GetProfileLogoutResponseObject interface {
	VisitGetProfileLogoutResponse(w http.ResponseWriter) error
}

type GetProfileLogout200ResponseHeaders struct {
	SetCookie string
}

type GetProfileLogout200Response struct {
	Headers GetProfileLogout200ResponseHeaders
}

func (response GetProfileLogout200Response) VisitGetProfileLogoutResponse(w http.ResponseWriter) error {
	w.Header().Set("Set-Cookie", fmt.Sprint(response.Headers.SetCookie))
	w.WriteHeader(200)
	return nil
}

type DeleteProfileNotificationsRequestObject struct {
}

type DeleteProfileNotificationsResponseObject interface {
	VisitDeleteProfileNotificationsResponse(w http.ResponseWriter) error
}

type DeleteProfileNotifications200Response struct {
}

func (response DeleteProfileNotifications200Response) VisitDeleteProfileNotificationsResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type GetProfileNotificationsRequestObject struct {
}

type GetProfileNotificationsResponseObject interface {
	VisitGetProfileNotificationsResponse(w http.ResponseWriter) error
}

type GetProfileNotifications200JSONResponse struct {
	Count int           `json:"count"`
	Items []interface{} `json:"items"`
}

func (response GetProfileNotifications200JSONResponse) VisitGetProfileNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostProfileNotificationsTaskUUIDHideRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type PostProfileNotificationsTaskUUIDHideResponseObject interface {
	VisitPostProfileNotificationsTaskUUIDHideResponse(w http.ResponseWriter) error
}

type PostProfileNotificationsTaskUUIDHide200Response struct {
}

func (response PostProfileNotificationsTaskUUIDHide200Response) VisitPostProfileNotificationsTaskUUIDHideResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteProfileNotificationsTaskUUIDStarRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type DeleteProfileNotificationsTaskUUIDStarResponseObject interface {
	VisitDeleteProfileNotificationsTaskUUIDStarResponse(w http.ResponseWriter) error
}

type DeleteProfileNotificationsTaskUUIDStar200Response struct {
}

func (response DeleteProfileNotificationsTaskUUIDStar200Response) VisitDeleteProfileNotificationsTaskUUIDStarResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostProfileNotificationsTaskUUIDStarRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type PostProfileNotificationsTaskUUIDStarResponseObject interface {
	VisitPostProfileNotificationsTaskUUIDStarResponse(w http.ResponseWriter) error
}

type PostProfileNotificationsTaskUUIDStar200Response struct {
}

func (response PostProfileNotificationsTaskUUIDStar200Response) VisitPostProfileNotificationsTaskUUIDStarResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchProfilePasswordRequestObject struct {
	Body *PatchProfilePasswordJSONRequestBody
}

type PatchProfilePasswordResponseObject interface {
	VisitPatchProfilePasswordResponse(w http.ResponseWriter) error
}

type PatchProfilePassword200Response struct {
}

func (response PatchProfilePassword200Response) VisitPatchProfilePasswordResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchProfilePassword401Response struct {
}

func (response PatchProfilePassword401Response) VisitPatchProfilePasswordResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type PatchProfilePhoneRequestObject struct {
	Body *PatchProfilePhoneJSONRequestBody
}

type PatchProfilePhoneResponseObject interface {
	VisitPatchProfilePhoneResponse(w http.ResponseWriter) error
}

type PatchProfilePhone200Response struct {
}

func (response PatchProfilePhone200Response) VisitPatchProfilePhoneResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteProfilePhotoRequestObject struct {
}

type DeleteProfilePhotoResponseObject interface {
	VisitDeleteProfilePhotoResponse(w http.ResponseWriter) error
}

type DeleteProfilePhoto200Response struct {
}

func (response DeleteProfilePhoto200Response) VisitDeleteProfilePhotoResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteProfilePhoto401Response struct {
}

func (response DeleteProfilePhoto401Response) VisitDeleteProfilePhotoResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type PatchProfilePhotoRequestObject struct {
	Body *multipart.Reader
}

type PatchProfilePhotoResponseObject interface {
	VisitPatchProfilePhotoResponse(w http.ResponseWriter) error
}

type PatchProfilePhoto200JSONResponse ProfilePhotoDTO

func (response PatchProfilePhoto200JSONResponse) VisitPatchProfilePhotoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchProfilePreferencesRequestObject struct {
	Body *PatchProfilePreferencesJSONRequestBody
}

type PatchProfilePreferencesResponseObject interface {
	VisitPatchProfilePreferencesResponse(w http.ResponseWriter) error
}

type PatchProfilePreferences200Response struct {
}

func (response PatchProfilePreferences200Response) VisitPatchProfilePreferencesResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostProfileResetRequestObject struct {
	Body *PostProfileResetJSONRequestBody
}

type PostProfileResetResponseObject interface {
	VisitPostProfileResetResponse(w http.ResponseWriter) error
}

type PostProfileReset200Response struct {
}

func (response PostProfileReset200Response) VisitPostProfileResetResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostProfileResetSendRequestObject struct {
	Body *PostProfileResetSendJSONRequestBody
}

type PostProfileResetSendResponseObject interface {
	VisitPostProfileResetSendResponse(w http.ResponseWriter) error
}

type PostProfileResetSend200ResponseHeaders struct {
	Hint string
}

type PostProfileResetSend200Response struct {
	Headers PostProfileResetSend200ResponseHeaders
}

func (response PostProfileResetSend200Response) VisitPostProfileResetSendResponse(w http.ResponseWriter) error {
	w.Header().Set("Hint", fmt.Sprint(response.Headers.Hint))
	w.WriteHeader(200)
	return nil
}

type PostProfileValidateRequestObject struct {
	Body *PostProfileValidateJSONRequestBody
}

type PostProfileValidateResponseObject interface {
	VisitPostProfileValidateResponse(w http.ResponseWriter) error
}

type PostProfileValidate200Response struct {
}

func (response PostProfileValidate200Response) VisitPostProfileValidateResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostProfileValidateSimpleRequestObject struct {
	Body *PostProfileValidateSimpleJSONRequestBody
}

type PostProfileValidateSimpleResponseObject interface {
	VisitPostProfileValidateSimpleResponse(w http.ResponseWriter) error
}

type PostProfileValidateSimple200Response struct {
}

func (response PostProfileValidateSimple200Response) VisitPostProfileValidateSimpleResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostProfileValidateSimpleSendRequestObject struct {
	Body *PostProfileValidateSimpleSendJSONRequestBody
}

type PostProfileValidateSimpleSendResponseObject interface {
	VisitPostProfileValidateSimpleSendResponse(w http.ResponseWriter) error
}

type PostProfileValidateSimpleSend200ResponseHeaders struct {
	Hint string
}

type PostProfileValidateSimpleSend200Response struct {
	Headers PostProfileValidateSimpleSend200ResponseHeaders
}

func (response PostProfileValidateSimpleSend200Response) VisitPostProfileValidateSimpleSendResponse(w http.ResponseWriter) error {
	w.Header().Set("Hint", fmt.Sprint(response.Headers.Hint))
	w.WriteHeader(200)
	return nil
}

type PostProfileValidateSendRequestObject struct {
	Body *PostProfileValidateSendJSONRequestBody
}

type PostProfileValidateSendResponseObject interface {
	VisitPostProfileValidateSendResponse(w http.ResponseWriter) error
}

type PostProfileValidateSend200ResponseHeaders struct {
	Hint string
}

type PostProfileValidateSend200Response struct {
	Headers PostProfileValidateSend200ResponseHeaders
}

func (response PostProfileValidateSend200Response) VisitPostProfileValidateSendResponse(w http.ResponseWriter) error {
	w.Header().Set("Hint", fmt.Sprint(response.Headers.Hint))
	w.WriteHeader(200)
	return nil
}

type StrictServerInterface interface {
	DeleteProfile(ctx context.Context, request DeleteProfileRequestObject) (DeleteProfileResponseObject, error)

	GetProfile(ctx context.Context, request GetProfileRequestObject) (GetProfileResponseObject, error)

	PostProfile(ctx context.Context, request PostProfileRequestObject) (PostProfileResponseObject, error)

	PatchProfileColor(ctx context.Context, request PatchProfileColorRequestObject) (PatchProfileColorResponseObject, error)

	PostProfileDislike(ctx context.Context, request PostProfileDislikeRequestObject) (PostProfileDislikeResponseObject, error)

	PatchProfileFio(ctx context.Context, request PatchProfileFioRequestObject) (PatchProfileFioResponseObject, error)

	GetProfileInvite(ctx context.Context, request GetProfileInviteRequestObject) (GetProfileInviteResponseObject, error)

	PatchProfileInviteUUIDAccept(ctx context.Context, request PatchProfileInviteUUIDAcceptRequestObject) (PatchProfileInviteUUIDAcceptResponseObject, error)

	PatchProfileInviteUUIDDecline(ctx context.Context, request PatchProfileInviteUUIDDeclineRequestObject) (PatchProfileInviteUUIDDeclineResponseObject, error)

	PostProfileLike(ctx context.Context, request PostProfileLikeRequestObject) (PostProfileLikeResponseObject, error)

	GetProfileLikes(ctx context.Context, request GetProfileLikesRequestObject) (GetProfileLikesResponseObject, error)

	PostProfileLogin(ctx context.Context, request PostProfileLoginRequestObject) (PostProfileLoginResponseObject, error)

	PostProfileLoginAs(ctx context.Context, request PostProfileLoginAsRequestObject) (PostProfileLoginAsResponseObject, error)

	GetProfileLogout(ctx context.Context, request GetProfileLogoutRequestObject) (GetProfileLogoutResponseObject, error)

	DeleteProfileNotifications(ctx context.Context, request DeleteProfileNotificationsRequestObject) (DeleteProfileNotificationsResponseObject, error)

	GetProfileNotifications(ctx context.Context, request GetProfileNotificationsRequestObject) (GetProfileNotificationsResponseObject, error)

	PostProfileNotificationsTaskUUIDHide(ctx context.Context, request PostProfileNotificationsTaskUUIDHideRequestObject) (PostProfileNotificationsTaskUUIDHideResponseObject, error)

	DeleteProfileNotificationsTaskUUIDStar(ctx context.Context, request DeleteProfileNotificationsTaskUUIDStarRequestObject) (DeleteProfileNotificationsTaskUUIDStarResponseObject, error)

	PostProfileNotificationsTaskUUIDStar(ctx context.Context, request PostProfileNotificationsTaskUUIDStarRequestObject) (PostProfileNotificationsTaskUUIDStarResponseObject, error)

	PatchProfilePassword(ctx context.Context, request PatchProfilePasswordRequestObject) (PatchProfilePasswordResponseObject, error)

	PatchProfilePhone(ctx context.Context, request PatchProfilePhoneRequestObject) (PatchProfilePhoneResponseObject, error)

	DeleteProfilePhoto(ctx context.Context, request DeleteProfilePhotoRequestObject) (DeleteProfilePhotoResponseObject, error)

	PatchProfilePhoto(ctx context.Context, request PatchProfilePhotoRequestObject) (PatchProfilePhotoResponseObject, error)

	PatchProfilePreferences(ctx context.Context, request PatchProfilePreferencesRequestObject) (PatchProfilePreferencesResponseObject, error)

	PostProfileReset(ctx context.Context, request PostProfileResetRequestObject) (PostProfileResetResponseObject, error)

	PostProfileResetSend(ctx context.Context, request PostProfileResetSendRequestObject) (PostProfileResetSendResponseObject, error)

	PostProfileValidate(ctx context.Context, request PostProfileValidateRequestObject) (PostProfileValidateResponseObject, error)

	PostProfileValidateSimple(ctx context.Context, request PostProfileValidateSimpleRequestObject) (PostProfileValidateSimpleResponseObject, error)

	PostProfileValidateSimpleSend(ctx context.Context, request PostProfileValidateSimpleSendRequestObject) (PostProfileValidateSimpleSendResponseObject, error)

	PostProfileValidateSend(ctx context.Context, request PostProfileValidateSendRequestObject) (PostProfileValidateSendResponseObject, error)
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

func (sh *strictHandler) DeleteProfile(ctx echo.Context) error {
	var request DeleteProfileRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteProfile(ctx.Request().Context(), request.(DeleteProfileRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteProfile")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteProfileResponseObject); ok {
		return validResponse.VisitDeleteProfileResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) GetProfile(ctx echo.Context) error {
	var request GetProfileRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProfile(ctx.Request().Context(), request.(GetProfileRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProfile")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetProfileResponseObject); ok {
		return validResponse.VisitGetProfileResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfile(ctx echo.Context) error {
	var request PostProfileRequestObject

	var body PostProfileJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfile(ctx.Request().Context(), request.(PostProfileRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfile")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileResponseObject); ok {
		return validResponse.VisitPostProfileResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchProfileColor(ctx echo.Context) error {
	var request PatchProfileColorRequestObject

	var body PatchProfileColorJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchProfileColor(ctx.Request().Context(), request.(PatchProfileColorRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchProfileColor")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchProfileColorResponseObject); ok {
		return validResponse.VisitPatchProfileColorResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileDislike(ctx echo.Context) error {
	var request PostProfileDislikeRequestObject

	var body PostProfileDislikeJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileDislike(ctx.Request().Context(), request.(PostProfileDislikeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileDislike")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileDislikeResponseObject); ok {
		return validResponse.VisitPostProfileDislikeResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchProfileFio(ctx echo.Context) error {
	var request PatchProfileFioRequestObject

	var body PatchProfileFioJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchProfileFio(ctx.Request().Context(), request.(PatchProfileFioRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchProfileFio")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchProfileFioResponseObject); ok {
		return validResponse.VisitPatchProfileFioResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) GetProfileInvite(ctx echo.Context) error {
	var request GetProfileInviteRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProfileInvite(ctx.Request().Context(), request.(GetProfileInviteRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProfileInvite")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetProfileInviteResponseObject); ok {
		return validResponse.VisitGetProfileInviteResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchProfileInviteUUIDAccept(ctx echo.Context, uUID Uuid) error {
	var request PatchProfileInviteUUIDAcceptRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchProfileInviteUUIDAccept(ctx.Request().Context(), request.(PatchProfileInviteUUIDAcceptRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchProfileInviteUUIDAccept")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchProfileInviteUUIDAcceptResponseObject); ok {
		return validResponse.VisitPatchProfileInviteUUIDAcceptResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchProfileInviteUUIDDecline(ctx echo.Context, uUID Uuid) error {
	var request PatchProfileInviteUUIDDeclineRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchProfileInviteUUIDDecline(ctx.Request().Context(), request.(PatchProfileInviteUUIDDeclineRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchProfileInviteUUIDDecline")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchProfileInviteUUIDDeclineResponseObject); ok {
		return validResponse.VisitPatchProfileInviteUUIDDeclineResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileLike(ctx echo.Context) error {
	var request PostProfileLikeRequestObject

	var body PostProfileLikeJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileLike(ctx.Request().Context(), request.(PostProfileLikeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileLike")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileLikeResponseObject); ok {
		return validResponse.VisitPostProfileLikeResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) GetProfileLikes(ctx echo.Context) error {
	var request GetProfileLikesRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProfileLikes(ctx.Request().Context(), request.(GetProfileLikesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProfileLikes")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetProfileLikesResponseObject); ok {
		return validResponse.VisitGetProfileLikesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileLogin(ctx echo.Context) error {
	var request PostProfileLoginRequestObject

	var body PostProfileLoginJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileLogin(ctx.Request().Context(), request.(PostProfileLoginRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileLogin")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileLoginResponseObject); ok {
		return validResponse.VisitPostProfileLoginResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileLoginAs(ctx echo.Context) error {
	var request PostProfileLoginAsRequestObject

	var body PostProfileLoginAsJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileLoginAs(ctx.Request().Context(), request.(PostProfileLoginAsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileLoginAs")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileLoginAsResponseObject); ok {
		return validResponse.VisitPostProfileLoginAsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) GetProfileLogout(ctx echo.Context) error {
	var request GetProfileLogoutRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProfileLogout(ctx.Request().Context(), request.(GetProfileLogoutRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProfileLogout")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetProfileLogoutResponseObject); ok {
		return validResponse.VisitGetProfileLogoutResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) DeleteProfileNotifications(ctx echo.Context) error {
	var request DeleteProfileNotificationsRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteProfileNotifications(ctx.Request().Context(), request.(DeleteProfileNotificationsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteProfileNotifications")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteProfileNotificationsResponseObject); ok {
		return validResponse.VisitDeleteProfileNotificationsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) GetProfileNotifications(ctx echo.Context) error {
	var request GetProfileNotificationsRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProfileNotifications(ctx.Request().Context(), request.(GetProfileNotificationsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProfileNotifications")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetProfileNotificationsResponseObject); ok {
		return validResponse.VisitGetProfileNotificationsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileNotificationsTaskUUIDHide(ctx echo.Context, uUID Uuid) error {
	var request PostProfileNotificationsTaskUUIDHideRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileNotificationsTaskUUIDHide(ctx.Request().Context(), request.(PostProfileNotificationsTaskUUIDHideRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileNotificationsTaskUUIDHide")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileNotificationsTaskUUIDHideResponseObject); ok {
		return validResponse.VisitPostProfileNotificationsTaskUUIDHideResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) DeleteProfileNotificationsTaskUUIDStar(ctx echo.Context, uUID Uuid) error {
	var request DeleteProfileNotificationsTaskUUIDStarRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteProfileNotificationsTaskUUIDStar(ctx.Request().Context(), request.(DeleteProfileNotificationsTaskUUIDStarRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteProfileNotificationsTaskUUIDStar")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteProfileNotificationsTaskUUIDStarResponseObject); ok {
		return validResponse.VisitDeleteProfileNotificationsTaskUUIDStarResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileNotificationsTaskUUIDStar(ctx echo.Context, uUID Uuid) error {
	var request PostProfileNotificationsTaskUUIDStarRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileNotificationsTaskUUIDStar(ctx.Request().Context(), request.(PostProfileNotificationsTaskUUIDStarRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileNotificationsTaskUUIDStar")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileNotificationsTaskUUIDStarResponseObject); ok {
		return validResponse.VisitPostProfileNotificationsTaskUUIDStarResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchProfilePassword(ctx echo.Context) error {
	var request PatchProfilePasswordRequestObject

	var body PatchProfilePasswordJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchProfilePassword(ctx.Request().Context(), request.(PatchProfilePasswordRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchProfilePassword")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchProfilePasswordResponseObject); ok {
		return validResponse.VisitPatchProfilePasswordResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchProfilePhone(ctx echo.Context) error {
	var request PatchProfilePhoneRequestObject

	var body PatchProfilePhoneJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchProfilePhone(ctx.Request().Context(), request.(PatchProfilePhoneRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchProfilePhone")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchProfilePhoneResponseObject); ok {
		return validResponse.VisitPatchProfilePhoneResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) DeleteProfilePhoto(ctx echo.Context) error {
	var request DeleteProfilePhotoRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteProfilePhoto(ctx.Request().Context(), request.(DeleteProfilePhotoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteProfilePhoto")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteProfilePhotoResponseObject); ok {
		return validResponse.VisitDeleteProfilePhotoResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchProfilePhoto(ctx echo.Context) error {
	var request PatchProfilePhotoRequestObject

	if reader, err := ctx.Request().MultipartReader(); err != nil {
		return err
	} else {
		request.Body = reader
	}

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchProfilePhoto(ctx.Request().Context(), request.(PatchProfilePhotoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchProfilePhoto")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchProfilePhotoResponseObject); ok {
		return validResponse.VisitPatchProfilePhotoResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PatchProfilePreferences(ctx echo.Context) error {
	var request PatchProfilePreferencesRequestObject

	var body PatchProfilePreferencesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchProfilePreferences(ctx.Request().Context(), request.(PatchProfilePreferencesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchProfilePreferences")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchProfilePreferencesResponseObject); ok {
		return validResponse.VisitPatchProfilePreferencesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileReset(ctx echo.Context) error {
	var request PostProfileResetRequestObject

	var body PostProfileResetJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileReset(ctx.Request().Context(), request.(PostProfileResetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileReset")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileResetResponseObject); ok {
		return validResponse.VisitPostProfileResetResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileResetSend(ctx echo.Context) error {
	var request PostProfileResetSendRequestObject

	var body PostProfileResetSendJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileResetSend(ctx.Request().Context(), request.(PostProfileResetSendRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileResetSend")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileResetSendResponseObject); ok {
		return validResponse.VisitPostProfileResetSendResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileValidate(ctx echo.Context) error {
	var request PostProfileValidateRequestObject

	var body PostProfileValidateJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileValidate(ctx.Request().Context(), request.(PostProfileValidateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileValidate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileValidateResponseObject); ok {
		return validResponse.VisitPostProfileValidateResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileValidateSimple(ctx echo.Context) error {
	var request PostProfileValidateSimpleRequestObject

	var body PostProfileValidateSimpleJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileValidateSimple(ctx.Request().Context(), request.(PostProfileValidateSimpleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileValidateSimple")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileValidateSimpleResponseObject); ok {
		return validResponse.VisitPostProfileValidateSimpleResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileValidateSimpleSend(ctx echo.Context) error {
	var request PostProfileValidateSimpleSendRequestObject

	var body PostProfileValidateSimpleSendJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileValidateSimpleSend(ctx.Request().Context(), request.(PostProfileValidateSimpleSendRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileValidateSimpleSend")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileValidateSimpleSendResponseObject); ok {
		return validResponse.VisitPostProfileValidateSimpleSendResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

func (sh *strictHandler) PostProfileValidateSend(ctx echo.Context) error {
	var request PostProfileValidateSendRequestObject

	var body PostProfileValidateSendJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostProfileValidateSend(ctx.Request().Context(), request.(PostProfileValidateSendRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProfileValidateSend")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostProfileValidateSendResponseObject); ok {
		return validResponse.VisitPostProfileValidateSendResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
