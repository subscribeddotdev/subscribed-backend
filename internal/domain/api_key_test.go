package domain_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/tests"
)

func TestNewApiKey(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		apiKeyName   string
		envID        domain.ID
		orgID        domain.ID
		expiresAt    *time.Time
		isTestApiKey bool
	}{
		{
			name:         "create_new_api_key",
			expectedErr:  "",
			apiKeyName:   gofakeit.AppName(),
			envID:        domain.NewID(),
			expiresAt:    nil,
			isTestApiKey: false,
			orgID:        domain.NewID(),
		},
		{
			name:         "create_new_test_api_key",
			expectedErr:  "",
			apiKeyName:   gofakeit.AppName(),
			envID:        domain.NewID(),
			expiresAt:    nil,
			isTestApiKey: true,
			orgID:        domain.NewID(),
		},
		{
			name:         "error_empty_name",
			expectedErr:  "name cannot be empty",
			apiKeyName:   "",
			envID:        domain.NewID(),
			expiresAt:    nil,
			isTestApiKey: false,
			orgID:        domain.NewID(),
		},
		{
			name:         "error_empty_or_invalid_env_id",
			expectedErr:  "envID cannot be empty",
			apiKeyName:   gofakeit.AppName(),
			envID:        domain.ID{},
			expiresAt:    nil,
			isTestApiKey: false,
			orgID:        domain.NewID(),
		},
		{
			name:         "error_expires_at_set_in_the_past",
			expectedErr:  "expiresAt cannot be set in the past",
			apiKeyName:   gofakeit.AppName(),
			envID:        domain.NewID(),
			expiresAt:    tests.ToPtr(time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)),
			isTestApiKey: false,
			orgID:        domain.NewID(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			apiKey, err := domain.NewApiKey(tc.apiKeyName, tc.orgID, tc.envID, tc.expiresAt, tc.isTestApiKey)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)

			assert.Equal(t, tc.envID, apiKey.EnvID())
			assert.True(t, apiKey.CreatedAt().Before(time.Now()))
			assert.NotEmpty(t, apiKey.SecretKey().FullKey())
			assert.NotEmpty(t, apiKey.SecretKey().String())
			assert.NotEqual(t, apiKey.SecretKey().FullKey(), apiKey.SecretKey().String())

			if tc.isTestApiKey {
				assert.Contains(t, apiKey.SecretKey().FullKey(), "_test_")
			} else {
				assert.Contains(t, apiKey.SecretKey().FullKey(), "_live_")
			}
		})
	}
}
