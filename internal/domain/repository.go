package domain

import (
	"context"
	"errors"
)

var (
	ErrApiKeyExists        = errors.New("the api key already exists")
	ErrEnvironmentNotFound = errors.New("environment not found")
)

type EnvironmentRepository interface {
	ByID(ctx context.Context, id ID) (*Environment, error)
	Insert(ctx context.Context, env *Environment) error
}

type ApplicationRepository interface {
	Insert(ctx context.Context, application *Application) error
}

type EventTypeRepository interface {
	Insert(ctx context.Context, eventType *EventType) error
}

type EndpointRepository interface {
	Insert(ctx context.Context, endpoint *Endpoint) error
}

type MessageRepository interface {
	Insert(ctx context.Context, message *Message) error
}

type ApiKeyRepository interface {
	Insert(ctx context.Context, apiKey *ApiKey) error
}
