package components_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	events "github.com/subscribeddotdev/subscribed-backend/events/go"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/eventdriven"
	"github.com/subscribeddotdev/subscribed-backend/tests/client"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestApplication_Lifecycle(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	app := ff.NewApplication().WithEnvironmentID(env.ID).Save()
	apiKey := ff.NewApiKey().WithOrgID(org.ID).WithEnvironmentID(env.ID).Save()
	eventType := ff.NewEventType().WithOrgID(org.ID).Save()
	endpoint := ff.NewEndpoint().WithEventTypeIDs([]string{eventType.ID}).WithAppID(app.ID).Save()
	apiClient := getClientWithApiKey(t, apiKey.SecretKey)

	pubsub, err := eventdriven.NewPubSub(os.Getenv("AMQP_URL"))
	require.NoError(t, err)
	defer func() {
		require.NoError(t, pubsub.Close())
	}()

	t.Run("send_message", func(t *testing.T) {
		// TODO: assert that MessageSent events have been published with the endpoint_id fixture above
		var foundMessageSentEvent events.MessageSent
		err = pubsub.Subscribe(command.MessageSentEvent, func(event *eventdriven.Event) error {
			fmt.Println("--->", string(event.Payload))
			require.NoError(t, json.Unmarshal(event.Payload, &foundMessageSentEvent))

			fmt.Println("--->")
			return nil
		})
		require.NoError(t, err)

		go func() {
			require.NoError(t, pubsub.Run(ctx))
		}()

		payload, err := gofakeit.JSON(&gofakeit.JSONOptions{
			Type: "object",
			Fields: []gofakeit.Field{
				{Name: "first_name", Function: "firstname"},
				{Name: "last_name", Function: "lastname"},
			},
		})
		require.NoError(t, err)

		resp, err := apiClient.SendMessage(ctx, app.ID, client.SendMessageRequest{
			EventTypeId: eventType.ID,
			Payload:     string(payload),
		})
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, resp.StatusCode)
		msg, err := models.Messages(
			models.MessageWhere.ApplicationID.EQ(app.ID),
			models.MessageWhere.EventTypeID.EQ(eventType.ID),
		).One(ctx, db)
		require.NoError(t, err)
		assert.Equal(t, string(payload), msg.Payload)

		assert.Eventually(t, func() bool {
			return foundMessageSentEvent.EndpointId == endpoint.ID
		}, time.Second*5, time.Millisecond*250)
	})
}
