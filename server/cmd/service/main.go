package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/subscribeddotdev/subscribed/server/internal/app/query"
	"github.com/subscribeddotdev/subscribed/server/internal/common/messaging"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/internal/domain/iam"
	"golang.org/x/sync/errgroup"

	"github.com/subscribeddotdev/subscribed/server/internal/adapters/events"
	"github.com/subscribeddotdev/subscribed/server/internal/adapters/psql"
	"github.com/subscribeddotdev/subscribed/server/internal/adapters/transaction"
	"github.com/subscribeddotdev/subscribed/server/internal/app"
	"github.com/subscribeddotdev/subscribed/server/internal/app/auth"
	"github.com/subscribeddotdev/subscribed/server/internal/app/command"
	"github.com/subscribeddotdev/subscribed/server/internal/common/logs"
	"github.com/subscribeddotdev/subscribed/server/internal/common/observability"
	"github.com/subscribeddotdev/subscribed/server/internal/common/postgres"
	amqpport "github.com/subscribeddotdev/subscribed/server/internal/ports/amqp"
	"github.com/subscribeddotdev/subscribed/server/internal/ports/http"
)

type Config struct {
	DatabaseUrl        string `envconfig:"SBS_DATABASE_URL" required:"true"`
	Port               int    `envconfig:"SBS_HTTP_PORT" required:"true"`
	ProductionMode     bool   `envconfig:"SBS_PRODUCTION_MODE" required:"true"`
	AllowedCorsOrigin  string `envconfig:"SBS_HTTP_ALLOWED_CORS" required:"true"`
	AmqpURL            string `envconfig:"SBS_AMQP_URL" required:"true"`
	HttpJwtSecret      string `envconfig:"SBS_HTTP_JWT_SECRET" required:"true"`
	WebFrontendEnabled bool   `envconfig:"SBS_WEB_FRONTEND_ENABLED" required:"true"`
}

func main() {
	logger := logs.New()

	if err := run(logger); err != nil {
		logger.Fatal("service crashed due to an error", "error", err)
	}

	os.Exit(0)
}

func run(logger *logs.Logger) error {
	config := &Config{}
	err := envconfig.Process("", config)
	if err != nil {
		return fmt.Errorf("unable to load env variables: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	g, gctx := errgroup.WithContext(ctx)

	db, err := postgres.Connect(config.DatabaseUrl)
	if err != nil {
		return err
	}

	err = postgres.ApplyMigrations(db, "misc/sql/migrations")
	if err != nil {
		return err
	}

	applicationRepo := psql.NewApplicationRepository(db)
	endpointRepo := psql.NewEndpointRepository(db)
	eventTypeRepo := psql.NewEventTypeRepository(db)
	memberRepo := psql.NewMemberRepository(db)
	apiKeyRepo := psql.NewApiKeyRepository(db)
	envRepo := psql.NewEnvironmentRepository(db)

	watermillLogger := watermill.NewSlogLogger(logger.Logger)
	publisher, err := messaging.NewAmqpPublisher(config.AmqpURL, logger)
	if err != nil {
		return err
	}
	defer func() { _ = publisher.Close() }()

	subscriber, err := amqp.NewSubscriber(amqp.NewDurableQueueConfig(config.AmqpURL), watermillLogger)
	if err != nil {
		return err
	}
	defer func() { _ = subscriber.Close() }()
	eventPublisher, err := events.NewPublisher(publisher)
	if err != nil {
		return err
	}

	txProvider := transaction.NewPsqlProvider(db, eventPublisher, logger)

	application := &app.App{
		Authorization: auth.NewService(memberRepo, apiKeyRepo),
		Command: app.Command{
			// IAM
			SignUp: observability.NewCommandDecorator[command.Signup](command.NewSignupHandler(txProvider), logger),
			SignIn: observability.NewCommandWithResultDecorator[command.SignIn, *iam.Member](command.NewSignInHandler(memberRepo), logger),

			// Applications
			CreateApplication: observability.NewCommandWithResultDecorator[command.CreateApplication](command.NewCreateApplicationHandler(applicationRepo), logger),

			// Endpoints
			AddEndpoint:         observability.NewCommandDecorator[command.AddEndpoint](command.NewAddEndpointHandler(endpointRepo), logger),
			CallWebhookEndpoint: observability.NewCommandDecorator[command.CallWebhookEndpoint](command.NewCallWebhookEndpointHandler(txProvider), logger),

			// Messages
			SendMessage: observability.NewCommandDecorator[command.SendMessage](command.NewSendMessageHandler(txProvider, endpointRepo), logger),

			// Event types
			CreateEventType: observability.NewCommandWithResultDecorator[command.CreateEventType, domain.EventTypeID](command.NewCreateEventTypeHandler(eventTypeRepo), logger),

			// API Keys
			CreateApiKey:  observability.NewCommandWithResultDecorator[command.CreateApiKey, *domain.ApiKey](command.NewCreateApiKeyHandler(apiKeyRepo, envRepo), logger),
			DestroyApiKey: observability.NewCommandDecorator[command.DestroyApiKey](command.NewDestroyApiKeyHandler(apiKeyRepo), logger),
		},
		Query: app.Query{
			AllEnvironments: observability.NewQueryDecorator[query.AllEnvironments, []*domain.Environment](query.NewEnvironmentsHandler(envRepo), logger),
			AllApiKeys:      observability.NewQueryDecorator[query.AllApiKeys, []*domain.ApiKey](query.NewAllApiKeysHandler(apiKeyRepo), logger),

			// Applications
			AllApplications: observability.NewQueryDecorator[query.AllApplications, query.Paginated[[]domain.Application]](query.NewAllApplicationsHandler(applicationRepo), logger),
			Application:     observability.NewQueryDecorator[query.Application, *domain.Application](query.NewApplicationHandler(applicationRepo), logger),

			// Event types
			AllEventTypes: observability.NewQueryDecorator[query.AllEventTypes, query.Paginated[[]domain.EventType]](query.NewAllEventTypesHandler(eventTypeRepo), logger),
			EventType:     observability.NewQueryDecorator[query.EventType, *domain.EventType](query.NewEventTypeHandler(eventTypeRepo), logger),
		},
	}

	router, err := message.NewRouter(message.RouterConfig{}, watermillLogger)
	if err != nil {
		return err
	}
	defer func() { _ = router.Close() }()

	poisonQueueMiddleware, err := middleware.PoisonQueue(publisher, "poison_queue")
	if err != nil {
		return err
	}

	retryMiddleware := middleware.Retry{
		MaxRetries:      3,
		Logger:          watermillLogger,
		InitialInterval: time.Millisecond * 3,
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(
		// Handle panics
		middleware.Recoverer,

		// Sets the correlation id to the messages' context
		// messaging.CorrelationIdMiddleware,

		// Send failed events to a specific queue
		poisonQueueMiddleware,

		// messaging.ErrorLoggerMiddleware(logger),

		// Retry failed events
		retryMiddleware.Middleware,
	)

	// Event handlers
	eventHandlers := amqpport.NewHandlers(application)
	for _, handler := range eventHandlers {
		router.AddNoPublisherHandler(handler.HandlerName(), handler.EventName(), subscriber, handler.Handle)
	}

	httpserver, err := http.NewServer(http.Config{
		Ctx:                ctx,
		Logger:             logger,
		Application:        application,
		Port:               config.Port,
		JwtSecret:          config.HttpJwtSecret,
		IsDebug:            !config.ProductionMode,
		WebFrontendEnabled: config.WebFrontendEnabled,
		AllowedCorsOrigin:  strings.Split(config.AllowedCorsOrigin, ","),
	})
	if err != nil {
		return err
	}

	g.Go(func() error {
		return httpserver.Start()
	})

	g.Go(func() error {
		return router.Run(ctx)
	})

	// Gracefully termination of services
	g.Go(func() error {
		<-gctx.Done()

		logger.Info("starting gracefully termination")

		tCtx, tCancel := context.WithTimeout(context.Background(), time.Second*5)
		defer tCancel()

		err = httpserver.Stop(tCtx)
		if err != nil {
			return err
		}

		logger.Info("service terminated gracefully")

		return nil
	})

	return g.Wait()
}
