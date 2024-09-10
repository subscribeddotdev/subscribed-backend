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
	EventTypeId string `json:"event_type_id"`
	Payload     string `json:"payload"`
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
