package domain

import (
	"context"
	"errors"
)

var (
	ErrApiKeyExists        = errors.New("the api key already exists")
	ErrApiKeyNotFound      = errors.New("api key not found")
	ErrAppNotFound         = errors.New("application not found")
	ErrEndpointNotFound    = errors.New("endpoint not found")
	ErrEventTypeNotFound   = errors.New("event type not found")
	ErrMessageNotFound     = errors.New("message not found")
	ErrEnvironmentNotFound = errors.New("environment not found")
)

type EnvironmentRepository interface {
	ByID(ctx context.Context, id EnvironmentID) (*Environment, error)
	Insert(ctx context.Context, env *Environment) error
	FindAll(ctx context.Context, orgID string) ([]*Environment, error)
}

type ApplicationRepository interface {
	Insert(ctx context.Context, application *Application) error
}

type EventTypeRepository interface {
	Insert(ctx context.Context, eventType *EventType) error
	ByID(ctx context.Context, orgID string, id EventTypeID) (*EventType, error)
}

type EndpointRepository interface {
	Insert(ctx context.Context, endpoint *Endpoint) error
	ByID(ctx context.Context, id EndpointID) (*Endpoint, error)
	ByEventTypeIdOrAppID(ctx context.Context, eventTypeID *EventTypeID, appID ApplicationID) ([]*Endpoint, error)
}

type MessageRepository interface {
	ByID(ctx context.Context, id MessageID) (*Message, error)
	Insert(ctx context.Context, message *Message) error
	SaveMessageSendAttempt(ctx context.Context, attempt *MessageSendAttempt) error
}

type ApiKeyRepository interface {
	Insert(ctx context.Context, apiKey *ApiKey) error
	Destroy(ctx context.Context, orgID string, id ApiKeyID) error
	FindBySecretKey(ctx context.Context, sk SecretKey) (*ApiKey, error)
	FindAll(ctx context.Context, orgID string, envID EnvironmentID) ([]*ApiKey, error)
}
