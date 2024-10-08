package domain_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/internal/domain/iam"
)

func TestNewMessage(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		eventTypeID   domain.EventTypeID
		orgID         string
		applicationID domain.ApplicationID
		payload       string
	}{
		{
			name:          "create_new_message",
			expectedErr:   "",
			eventTypeID:   domain.NewEventTypeID(),
			orgID:         iam.NewOrgID().String(),
			applicationID: domain.NewApplicationID(),
			payload:       gofakeit.Sentence(10),
		},
		{
			name:          "error_invalid_event_type_id",
			expectedErr:   "eventTypeID cannot be empty",
			eventTypeID:   domain.EventTypeID(""),
			orgID:         iam.NewOrgID().String(),
			applicationID: domain.NewApplicationID(),
			payload:       gofakeit.Sentence(10),
		},
		{
			name:          "error_invalid_org_id",
			expectedErr:   "orgID cannot be empty",
			eventTypeID:   domain.NewEventTypeID(),
			orgID:         "",
			applicationID: domain.NewApplicationID(),
			payload:       gofakeit.Sentence(10),
		},
		{
			name:          "error_invalid_application_id",
			expectedErr:   "applicationID cannot be empty",
			eventTypeID:   domain.NewEventTypeID(),
			orgID:         iam.NewOrgID().String(),
			applicationID: domain.ApplicationID(""),
			payload:       gofakeit.Sentence(10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			msg, err := domain.NewMessage(&tc.eventTypeID, tc.orgID, tc.applicationID, tc.payload)
			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)

			assert.Contains(t, msg.Id(), "msg_")
			assert.Equal(t, tc.orgID, msg.OrgID())
			assert.Equal(t, tc.eventTypeID.String(), msg.EventTypeID().String())
			assert.Equal(t, tc.applicationID.String(), msg.ApplicationID().String())
			assert.Equal(t, tc.payload, msg.Payload())
		})
	}
}
