// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package http

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Defines values for EnvironmentType.
const (
	Development EnvironmentType = "development"
	Production  EnvironmentType = "production"
)

// AddEndpointRequest defines model for AddEndpointRequest.
type AddEndpointRequest struct {
	Description  *string   `json:"description,omitempty"`
	EventTypeIds *[]string `json:"event_type_ids,omitempty"`
	Url          string    `json:"url"`
}

// ApiKey defines model for ApiKey.
type ApiKey struct {
	CreatedAt       time.Time  `json:"created_at"`
	EnvironmentId   string     `json:"environment_id"`
	ExpiresAt       *time.Time `json:"expires_at,omitempty"`
	Id              string     `json:"id"`
	MaskedSecretKey string     `json:"masked_secret_key"`
	Name            string     `json:"name"`
	OrganizationId  string     `json:"organization_id"`
}

// Application defines model for Application.
type Application struct {
	CreatedAt     time.Time `json:"created_at"`
	EnvironmentId string    `json:"environment_id"`
	Id            string    `json:"id"`
	Name          string    `json:"name"`
}

// CreateApiKeyPayload defines model for CreateApiKeyPayload.
type CreateApiKeyPayload struct {
	UnmaskedApiKey string `json:"unmasked_api_key"`
}

// CreateApiKeyRequest defines model for CreateApiKeyRequest.
type CreateApiKeyRequest struct {
	EnvironmentId string     `json:"environment_id"`
	ExpiresAt     *time.Time `json:"expires_at"`
	Name          string     `json:"name"`
}

// CreateApplicationPayload defines model for CreateApplicationPayload.
type CreateApplicationPayload struct {
	Id string `json:"id"`
}

// CreateApplicationRequest defines model for CreateApplicationRequest.
type CreateApplicationRequest struct {
	Name string `json:"name"`
}

// CreateEventTypePayload defines model for CreateEventTypePayload.
type CreateEventTypePayload struct {
	Id string `json:"id"`
}

// CreateEventTypeRequest defines model for CreateEventTypeRequest.
type CreateEventTypeRequest struct {
	Description   *string `json:"description,omitempty"`
	Name          string  `json:"name"`
	Schema        *string `json:"schema,omitempty"`
	SchemaExample *string `json:"schema_example,omitempty"`
}

// Environment defines model for Environment.
type Environment struct {
	ArchivedAt     *time.Time      `json:"archived_at,omitempty"`
	CreatedAt      time.Time       `json:"created_at"`
	Id             string          `json:"id"`
	Name           string          `json:"name"`
	OrganizationId string          `json:"organization_id"`
	Type           EnvironmentType `json:"type"`
}

// EnvironmentType defines model for Environment.Type.
type EnvironmentType string

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Error Error custom error code such as 'email_in_use'
	Error string `json:"error"`

	// Message A description about the error
	Message string `json:"message"`
}

// EventType defines model for EventType.
type EventType struct {
	ArchivedAt    *time.Time `json:"archived_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	Description   string     `json:"description"`
	Id            string     `json:"id"`
	Name          string     `json:"name"`
	Schema        string     `json:"schema"`
	SchemaExample string     `json:"schema_example"`
}

// GetAllApiKeysPayload defines model for GetAllApiKeysPayload.
type GetAllApiKeysPayload struct {
	Data []ApiKey `json:"data"`
}

// GetAllEnvironmentsPayload defines model for GetAllEnvironmentsPayload.
type GetAllEnvironmentsPayload struct {
	Data []Environment `json:"data"`
}

// GetApplicationByIdPayload defines model for GetApplicationByIdPayload.
type GetApplicationByIdPayload struct {
	Data Application `json:"data"`
}

// GetApplicationsPayload defines model for GetApplicationsPayload.
type GetApplicationsPayload struct {
	Data       []Application `json:"data"`
	Pagination Pagination    `json:"pagination"`
}

// GetEventTypeByIdPayload defines model for GetEventTypeByIdPayload.
type GetEventTypeByIdPayload struct {
	Data EventType `json:"data"`
}

// GetEventTypesPayload defines model for GetEventTypesPayload.
type GetEventTypesPayload struct {
	Data       []EventType `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// Pagination defines model for Pagination.
type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
	TotalPages  int `json:"total_pages"`
}

// SendMessageRequest defines model for SendMessageRequest.
type SendMessageRequest struct {
	EventTypeId *string `json:"event_type_id,omitempty"`
	Payload     string  `json:"payload"`
}

// SignInPayload defines model for SignInPayload.
type SignInPayload struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	Id        string `json:"id"`
	LastName  string `json:"last_name"`
	Token     string `json:"token"`
}

// SigninRequest defines model for SigninRequest.
type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignupRequest defines model for SignupRequest.
type SignupRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

// ApplicationId defines model for applicationId.
type ApplicationId = string

// EnvironmentId defines model for environmentId.
type EnvironmentId = string

// PaginationParamLimit defines model for paginationParamLimit.
type PaginationParamLimit = int

// PaginationParamPage defines model for paginationParamPage.
type PaginationParamPage = int

// DefaultError defines model for DefaultError.
type DefaultError = ErrorResponse

// NotFoundError defines model for NotFoundError.
type NotFoundError = ErrorResponse

// GetAllApiKeysParams defines parameters for GetAllApiKeys.
type GetAllApiKeysParams struct {
	EnvironmentId string `form:"environment_id" json:"environment_id"`
}

// GetApplicationsParams defines parameters for GetApplications.
type GetApplicationsParams struct {
	EnvironmentID EnvironmentId `form:"environmentID" json:"environmentID"`

	// Limit The number of items per page
	Limit *PaginationParamLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// Page The page number
	Page *PaginationParamPage `form:"page,omitempty" json:"page,omitempty"`
}

// GetEventTypesParams defines parameters for GetEventTypes.
type GetEventTypesParams struct {
	// Limit The number of items per page
	Limit *PaginationParamLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// Page The page number
	Page *PaginationParamPage `form:"page,omitempty" json:"page,omitempty"`
}

// CreateApiKeyJSONRequestBody defines body for CreateApiKey for application/json ContentType.
type CreateApiKeyJSONRequestBody = CreateApiKeyRequest

// CreateApplicationJSONRequestBody defines body for CreateApplication for application/json ContentType.
type CreateApplicationJSONRequestBody = CreateApplicationRequest

// AddEndpointJSONRequestBody defines body for AddEndpoint for application/json ContentType.
type AddEndpointJSONRequestBody = AddEndpointRequest

// SendMessageJSONRequestBody defines body for SendMessage for application/json ContentType.
type SendMessageJSONRequestBody = SendMessageRequest

// CreateEventTypeJSONRequestBody defines body for CreateEventType for application/json ContentType.
type CreateEventTypeJSONRequestBody = CreateEventTypeRequest

// SignInJSONRequestBody defines body for SignIn for application/json ContentType.
type SignInJSONRequestBody = SigninRequest

// SignUpJSONRequestBody defines body for SignUp for application/json ContentType.
type SignUpJSONRequestBody = SignupRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all api keys
	// (GET /api-keys)
	GetAllApiKeys(ctx echo.Context, params GetAllApiKeysParams) error
	// Create a new api key
	// (POST /api-keys)
	CreateApiKey(ctx echo.Context) error
	// Destroy a new api key given an id
	// (DELETE /api-keys/{apiKeyId})
	DestroyApiKey(ctx echo.Context, apiKeyId string) error
	// Returns a list of applications based on the org_id and environment_id
	// (GET /applications)
	GetApplications(ctx echo.Context, params GetApplicationsParams) error
	// Creates a new application
	// (POST /applications)
	CreateApplication(ctx echo.Context) error
	// Returns an application
	// (GET /applications/{applicationID})
	GetApplicationById(ctx echo.Context, applicationID ApplicationId) error
	// Add an endpoint to an application
	// (POST /applications/{applicationID}/endpoints)
	AddEndpoint(ctx echo.Context, applicationID string) error
	// Send a message to an application
	// (POST /applications/{applicationID}/messages)
	SendMessage(ctx echo.Context, applicationID ApplicationId) error
	// Get all environments
	// (GET /environments)
	GetEnvironments(ctx echo.Context) error
	// Get event types by org_id
	// (GET /event-types)
	GetEventTypes(ctx echo.Context, params GetEventTypesParams) error
	// Creates a new event type
	// (POST /event-types)
	CreateEventType(ctx echo.Context) error
	// Get event type by id and org_id
	// (GET /event-types/{eventTypeID})
	GetEventTypeById(ctx echo.Context, eventTypeID string) error

	// (GET /health)
	HealthCheck(ctx echo.Context) error
	// Authenticates a user
	// (POST /signin)
	SignIn(ctx echo.Context) error
	// Creates a new organization and it's first member
	// (POST /signup)
	SignUp(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAllApiKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllApiKeys(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAllApiKeysParams
	// ------------- Required query parameter "environment_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "environment_id", ctx.QueryParams(), &params.EnvironmentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter environment_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAllApiKeys(ctx, params)
	return err
}

// CreateApiKey converts echo context to params.
func (w *ServerInterfaceWrapper) CreateApiKey(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateApiKey(ctx)
	return err
}

// DestroyApiKey converts echo context to params.
func (w *ServerInterfaceWrapper) DestroyApiKey(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "apiKeyId" -------------
	var apiKeyId string

	err = runtime.BindStyledParameterWithOptions("simple", "apiKeyId", ctx.Param("apiKeyId"), &apiKeyId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter apiKeyId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DestroyApiKey(ctx, apiKeyId)
	return err
}

// GetApplications converts echo context to params.
func (w *ServerInterfaceWrapper) GetApplications(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{})

	ctx.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApplicationsParams
	// ------------- Required query parameter "environmentID" -------------

	err = runtime.BindQueryParameter("form", true, true, "environmentID", ctx.QueryParams(), &params.EnvironmentID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter environmentID: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetApplications(ctx, params)
	return err
}

// CreateApplication converts echo context to params.
func (w *ServerInterfaceWrapper) CreateApplication(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateApplication(ctx)
	return err
}

// GetApplicationById converts echo context to params.
func (w *ServerInterfaceWrapper) GetApplicationById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "applicationID" -------------
	var applicationID ApplicationId

	err = runtime.BindStyledParameterWithOptions("simple", "applicationID", ctx.Param("applicationID"), &applicationID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter applicationID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{})

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetApplicationById(ctx, applicationID)
	return err
}

// AddEndpoint converts echo context to params.
func (w *ServerInterfaceWrapper) AddEndpoint(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "applicationID" -------------
	var applicationID string

	err = runtime.BindStyledParameterWithOptions("simple", "applicationID", ctx.Param("applicationID"), &applicationID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter applicationID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AddEndpoint(ctx, applicationID)
	return err
}

// SendMessage converts echo context to params.
func (w *ServerInterfaceWrapper) SendMessage(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "applicationID" -------------
	var applicationID ApplicationId

	err = runtime.BindStyledParameterWithOptions("simple", "applicationID", ctx.Param("applicationID"), &applicationID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter applicationID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SendMessage(ctx, applicationID)
	return err
}

// GetEnvironments converts echo context to params.
func (w *ServerInterfaceWrapper) GetEnvironments(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEnvironments(ctx)
	return err
}

// GetEventTypes converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventTypes(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetEventTypesParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventTypes(ctx, params)
	return err
}

// CreateEventType converts echo context to params.
func (w *ServerInterfaceWrapper) CreateEventType(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateEventType(ctx)
	return err
}

// GetEventTypeById converts echo context to params.
func (w *ServerInterfaceWrapper) GetEventTypeById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "eventTypeID" -------------
	var eventTypeID string

	err = runtime.BindStyledParameterWithOptions("simple", "eventTypeID", ctx.Param("eventTypeID"), &eventTypeID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter eventTypeID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	ctx.Set(ApiKeyAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEventTypeById(ctx, eventTypeID)
	return err
}

// HealthCheck converts echo context to params.
func (w *ServerInterfaceWrapper) HealthCheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.HealthCheck(ctx)
	return err
}

// SignIn converts echo context to params.
func (w *ServerInterfaceWrapper) SignIn(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SignIn(ctx)
	return err
}

// SignUp converts echo context to params.
func (w *ServerInterfaceWrapper) SignUp(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SignUp(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
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

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api-keys", wrapper.GetAllApiKeys)
	router.POST(baseURL+"/api-keys", wrapper.CreateApiKey)
	router.DELETE(baseURL+"/api-keys/:apiKeyId", wrapper.DestroyApiKey)
	router.GET(baseURL+"/applications", wrapper.GetApplications)
	router.POST(baseURL+"/applications", wrapper.CreateApplication)
	router.GET(baseURL+"/applications/:applicationID", wrapper.GetApplicationById)
	router.POST(baseURL+"/applications/:applicationID/endpoints", wrapper.AddEndpoint)
	router.POST(baseURL+"/applications/:applicationID/messages", wrapper.SendMessage)
	router.GET(baseURL+"/environments", wrapper.GetEnvironments)
	router.GET(baseURL+"/event-types", wrapper.GetEventTypes)
	router.POST(baseURL+"/event-types", wrapper.CreateEventType)
	router.GET(baseURL+"/event-types/:eventTypeID", wrapper.GetEventTypeById)
	router.GET(baseURL+"/health", wrapper.HealthCheck)
	router.POST(baseURL+"/signin", wrapper.SignIn)
	router.POST(baseURL+"/signup", wrapper.SignUp)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
