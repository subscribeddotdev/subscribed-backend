package app

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/app/auth"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/query"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type CommandHandler[C any] interface {
	Execute(ctx context.Context, cmd C) error
}

type QueryHandler[Q, R any] interface {
	Execute(ctx context.Context, q Q) (R, error)
}

type Command struct {
	AddEndpoint         CommandHandler[command.AddEndpoint]
	SendMessage         CommandHandler[command.SendMessage]
	CreateApplication   CommandHandler[command.CreateApplication]
	CreateOrganization  CommandHandler[command.CreateOrganization]
	CreateEventType     CommandHandler[command.CreateEventType]
	CreateApiKey        CommandHandler[command.CreateApiKey]
	CallWebhookEndpoint CommandHandler[command.CallWebhookEndpoint]
}

type Query struct {
	Environments QueryHandler[query.Environments, []*domain.Environment]
}

type App struct {
	Authorization *auth.Service
	Command       Command
	Query         Query
}
