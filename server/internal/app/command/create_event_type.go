package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/internal/domain/iam"
)

type CreateEventType struct {
	OrgID         iam.OrgID
	Name          string
	Description   string
	Schema        string
	SchemaExample string
}

type CreateEventTypeHandler struct {
	repo domain.EventTypeRepository
}

func NewCreateEventTypeHandler(repo domain.EventTypeRepository) CreateEventTypeHandler {
	return CreateEventTypeHandler{
		repo: repo,
	}
}

func (c CreateEventTypeHandler) Execute(ctx context.Context, cmd CreateEventType) (domain.EventTypeID, error) {
	eventType, err := domain.NewEventType(cmd.OrgID.String(), cmd.Name, cmd.Description, cmd.Schema, cmd.SchemaExample)
	if err != nil {
		return "", err
	}

	return eventType.ID(), c.repo.Insert(ctx, eventType)
}
