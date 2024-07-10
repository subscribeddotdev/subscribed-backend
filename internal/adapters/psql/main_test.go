package psql_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/psql"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/postgres"
)

var (
	db               *sql.DB
	ctx              context.Context
	environmentRepo  *psql.EnvironmentRepository
	applicationRepo  *psql.ApplicationRepository
	endpointRepo     *psql.EndpointRepository
	organizationRepo *psql.OrganizationRepository
	apiKeyRepo       *psql.ApiKeyRepository
	msgRepo          *psql.MessageRepository
)

// Set up file
func TestMain(m *testing.M) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()

	var err error
	db, err = postgres.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = postgres.ApplyMigrations(db, "../../../misc/sql/migrations")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	environmentRepo = psql.NewEnvironmentRepository(db)
	applicationRepo = psql.NewApplicationRepository(db)
	endpointRepo = psql.NewEndpointRepository(db)
	organizationRepo = psql.NewOrganizationRepository(db)
	apiKeyRepo = psql.NewApiKeyRepository(db)
	msgRepo = psql.NewMessageRepository(db)

	os.Exit(m.Run())
}
