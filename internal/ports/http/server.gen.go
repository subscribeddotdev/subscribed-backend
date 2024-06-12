// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
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

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ClerkWebhookEmailAddress defines model for ClerkWebhookEmailAddress.
type ClerkWebhookEmailAddress struct {
	// EmailAddress User's email address
	EmailAddress string `json:"email_address"`

	// Id Unique identifier for the email address
	Id string `json:"id"`

	// LinkedTo (Array is empty for this event)
	LinkedTo *[]map[string]interface{} `json:"linked_to"`

	// Object Object type (always "email_address" for this event)
	Object       *string `json:"object,omitempty"`
	Verification *struct {
		// Status Verification status (e.g., "verified", "unverified")
		Status *string `json:"status,omitempty"`

		// Strategy Verification strategy (e.g., "ticket", "link")
		Strategy *string `json:"strategy,omitempty"`
	} `json:"verification,omitempty"`
}

// ClerkWebhookUserCreatedData defines model for ClerkWebhookUserCreatedData.
type ClerkWebhookUserCreatedData struct {
	// Birthday User's birthday (empty string if not set)
	Birthday *string `json:"birthday,omitempty"`

	// CreatedAt Timestamp (epoch milliseconds) representing user creation time
	CreatedAt      int                        `json:"created_at"`
	EmailAddresses []ClerkWebhookEmailAddress `json:"email_addresses"`

	// ExternalAccounts (Array is empty for this event)
	ExternalAccounts *[]map[string]interface{} `json:"external_accounts,omitempty"`

	// ExternalId User's external identifier
	ExternalId *string `json:"external_id"`

	// FirstName User's first name
	FirstName *string `json:"first_name"`

	// Gender User's gender (empty string if not set)
	Gender *string `json:"gender,omitempty"`

	// Id Unique identifier for the user
	Id string `json:"id"`

	// ImageUrl User's image URL (may be redacted)
	ImageUrl *string `json:"image_url,omitempty"`

	// LastName User's last name
	LastName *string `json:"last_name"`

	// LastSignInAt Timestamp (epoch milliseconds) representing last sign-in time
	LastSignInAt *int `json:"last_sign_in_at"`

	// Object Object type (always "user" for this event)
	Object *string `json:"object,omitempty"`

	// PasswordEnabled Whether the user has password authentication enabled
	PasswordEnabled bool `json:"password_enabled"`

	// PhoneNumbers (Array is empty for this event)
	PhoneNumbers *[]map[string]interface{} `json:"phone_numbers,omitempty"`

	// PrimaryEmailAddressId Unique identifier for the primary email address
	PrimaryEmailAddressId *string `json:"primary_email_address_id"`

	// PrimaryPhoneNumberId Unique identifier for the primary phone number (null if not set)
	PrimaryPhoneNumberId *string `json:"primary_phone_number_id"`

	// PrimaryWeb3WalletId Unique identifier for the primary web3 wallet (null if not set)
	PrimaryWeb3WalletId *string `json:"primary_web3_wallet_id"`

	// PrivateMetadata User's private metadata (empty object for this event)
	PrivateMetadata *map[string]interface{} `json:"private_metadata,omitempty"`

	// ProfileImageUrl User's profile image URL (may be redacted)
	ProfileImageUrl *string `json:"profile_image_url,omitempty"`

	// PublicMetadata User's public metadata (empty object for this event)
	PublicMetadata *map[string]interface{} `json:"public_metadata,omitempty"`

	// TwoFactorEnabled Whether two-factor authentication is enabled
	TwoFactorEnabled bool `json:"two_factor_enabled"`

	// UnsafeMetadata User's unsafe metadata (empty object for this event)
	UnsafeMetadata *map[string]interface{} `json:"unsafe_metadata,omitempty"`

	// UpdatedAt Timestamp (epoch milliseconds) representing user update time
	UpdatedAt *int `json:"updated_at,omitempty"`

	// Username Username (null if not set)
	Username *string `json:"username"`

	// Web3Wallets (Array is empty for this event)
	Web3Wallets *[]map[string]interface{} `json:"web3_wallets,omitempty"`
}

// CreateAccountRequest defines model for CreateAccountRequest.
type CreateAccountRequest struct {
	Data ClerkWebhookUserCreatedData `json:"data"`

	// Object Event type (always "user.created" for this event)
	Object string `json:"object"`

	// Type Event type (always "user.created" for this event)
	Type string `json:"type"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Error Error custom error code such as 'email_in_use'
	Error string `json:"error"`

	// Message A description about the error
	Message string `json:"message"`
}

// DefaultError defines model for DefaultError.
type DefaultError = ErrorResponse

// CreateAccountParams defines parameters for CreateAccount.
type CreateAccountParams struct {
	SvixId        string `json:"svix-id"`
	SvixTimestamp string `json:"svix-timestamp"`
	SvixSignature string `json:"svix-signature"`
}

// CreateAccountJSONRequestBody defines body for CreateAccount for application/json ContentType.
type CreateAccountJSONRequestBody = CreateAccountRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Creates a new account
	// (POST /webhooks/account)
	CreateAccount(ctx echo.Context, params CreateAccountParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateAccount converts echo context to params.
func (w *ServerInterfaceWrapper) CreateAccount(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateAccountParams

	headers := ctx.Request().Header
	// ------------- Required header parameter "svix-id" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("svix-id")]; found {
		var SvixId string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for svix-id, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "svix-id", runtime.ParamLocationHeader, valueList[0], &SvixId)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter svix-id: %s", err))
		}

		params.SvixId = SvixId
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter svix-id is required, but not found"))
	}
	// ------------- Required header parameter "svix-timestamp" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("svix-timestamp")]; found {
		var SvixTimestamp string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for svix-timestamp, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "svix-timestamp", runtime.ParamLocationHeader, valueList[0], &SvixTimestamp)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter svix-timestamp: %s", err))
		}

		params.SvixTimestamp = SvixTimestamp
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter svix-timestamp is required, but not found"))
	}
	// ------------- Required header parameter "svix-signature" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("svix-signature")]; found {
		var SvixSignature string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for svix-signature, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "svix-signature", runtime.ParamLocationHeader, valueList[0], &SvixSignature)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter svix-signature: %s", err))
		}

		params.SvixSignature = SvixSignature
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter svix-signature is required, but not found"))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateAccount(ctx, params)
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

	router.POST(baseURL+"/webhooks/account", wrapper.CreateAccount)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RY3W7bOBN9FYLfBzQFHCvd3hS+WjfxAl20WyBJtxdNIIyoscRGIlVyZNco/O4LkpJl",
	"W3Li/PTOFMdnhsMzh0P+4kKXlVaoyPLJL27QVlpZ9IMLnENd0MwYbdxYaEWoyP2EqiqkAJJaRd+tVu6b",
	"FTmW4H793+CcT/j/og48CrM28miXjRu+Xq9HPEUrjKwcGJ/wKctQoZGCoTNlprMdNT58dOcFmruvmORa",
	"381KkMU0TQ1aP1cZXaEhGdaBbjaGbnrX4ReL5pVl3oq1ViNOqwr5hFsyUmV8PeIyHfivkj9qZDJFRXIu",
	"0bC5NoxyfBiukOoO05h0H/VkagysmHRBVbRqMN1wgYpe8xGXhKVfSgOrk+8oyMGquiggKZBPyNS48QsO",
	"0c03lj2fn/135szZCRRLWFl2s5u6Gz4QSW9dCzRy3nCjvxeWgOqBTfh3618sGLETHGfjEbtpIDG94W5U",
	"q248GIElA4TZ6kEvwazzQ1LcIQUvbnsG8dej/aSvR9zgj1oaTPnk2x7dPG+2d/t2PdrhrqPfuUEgTC+A",
	"oJ+yRBrKU1gdZG5rwE4CXUKkTM6Z0sQsDu+TCD5jGCDDtSzREpQVO8FKi5yVsiikRaFVal8zg5VB6yiv",
	"MlZbNMyDuaSSLLHzJhVhhsa528lKWNiGxPfpxcEyX/e5jT8JjYIiBiF03Ujay5XWQXeDytCoSmOzJRH8",
	"YJF2mzOXxlKsoMSDyN6EeZMjADNUKZqDYGH6cQR6nB46mgyilJBhXJviYGzegn25/MhOSlixBJnBFARh",
	"OhhWAQ9lzlkcnTgPZ2WmYqmeXSveswM7lZtSORDBVuk8TrJdoo9T6gqsXWqTxqhcAAPb+TVHyrHbQJaD",
	"Ze3fGNSUu4U1etqibDwlWhcIyrvKtcJY1WWC5jdXZWVkCWYV7whO/DiyNhi9Q/xBtrTOt9f7RN8eggUI",
	"duI87xXk0cEsMXkbL6EokJ4Yi0NgAeHJoSyAMC6RIG1OucHabAxZa9hKUtjvw6zu+FAZPZcFxg8LS2P5",
	"WIGp6qSQ4oileLtnrISWOp6DIG2OqNClPg22+2XpXNxTmbWyMD9iY4LdM1ZTV+nL9RsB7HC34YwOnwFu",
	"5mk83iql3ypjey2lbyF7cj3IkJ3Grt92+e7TG0xDj3SJP2q01G87WzYc25zt97H3HFwzl46hc2vcxH7c",
	"+RU+/Bbwvfz7XGzW01jfDpB893bbv4m2V+m9kP1FV9SWdNnceoVOkdla5AwsexV2Uaq4tvhqKBUlWgvZ",
	"QDambGvMINE1heupj+ShdbdWLfzt0N3HoqiNpNWVI0VY53sEg2ZaU+4vMH70lzalq33+99dr3tzkvRb5",
	"2S6UnKgK7wJt73yhxUCxOTs7iaJMUl4nY6HLyNaJs0gwTTWluNj6cJqAuEOVRpez6cWn2bh0leIPhycC",
	"+d5VzXX7MAKB6H6n+MQ18KVUeixyUBko+WfmJhw47714XG3A3WWugXdXRoENi4KU8U8frp8ZdfTxw/ns",
	"nyu/fldBaEr7eX6FZiEFPj0XI06SCk+kockFGhuWejY+G7/x0lChgkryCX87Phu/9fJGud/maBkExUbN",
	"Nc7XkQ4a5arJn2sfUj7ZVTKPYaBE8j3mt19cep4gpOHSFZJoF/LnqRfUjulB7bsHrF5V3IdF7eH1cpDu",
	"igBUG3wU5G0wRkvvdbp6sRe7weNivSsVLjT/Yev58I+zN/2qPb+cTa9nF6EI/PPiIfcbrGjnHdIrTl26",
	"tnRDAMuAKVwy2BCBIHMU4F6CbtdBpsyiJUZXQpMoKrSAIteWJu/O3p3x9e36vwAAAP//ZpnZyxYVAAA=",
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
