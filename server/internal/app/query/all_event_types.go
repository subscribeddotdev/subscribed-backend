package query

import (
	"context"

	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/internal/domain/iam"
)

type AllEventTypes struct {
	PaginationParams
	OrgID string
}

type eventTypesFinder interface {
	FindAll(ctx context.Context, orgID iam.OrgID, pagination PaginationParams) (Paginated[[]domain.EventType], error)
}

type allEventTypesHandler struct {
	eventTypesFinder eventTypesFinder
}

func NewAllEventTypesHandler(eventTypesFinder eventTypesFinder) allEventTypesHandler {
	return allEventTypesHandler{
		eventTypesFinder: eventTypesFinder,
	}
}

func (h allEventTypesHandler) Execute(ctx context.Context, q AllEventTypes) (Paginated[[]domain.EventType], error) {
	return h.eventTypesFinder.FindAll(ctx, iam.OrgID(q.OrgID), q.PaginationParams)
}
