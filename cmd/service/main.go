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
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/clerkhttp"
	"golang.org/x/sync/errgroup"

	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/events"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/transaction"
	"github.com/subscribeddotdev/subscribed-backend/internal/app"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/logs"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/postgres"
	"github.com/subscribeddotdev/subscribed-backend/internal/ports/http"
)

type Config struct {
	DatabaseUrl            string `envconfig:"DATABASE_URL"`
	Port                   int    `envconfig:"HTTP_PORT"`
	ProductionMode         bool   `envconfig:"PRODUCTION_MODE"`
	AllowedCorsOrigin      string `envconfig:"HTTP_ALLOWED_CORS"`
	AmqpUrL                string `envconfig:"AMQP_URL"`
	ClerkSecretKey         string `envconfig:"CLERK_SECRET_KEY"`
	ClerkEmulatorServerURL string `envconfig:"CLERK_EMULATOR_SERVER_URL"`
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

	if !config.ProductionMode {
		logger.Info("[LOGIN PROVIDER] setting up clerk to work with the simulator", "emulator_url", config.ClerkEmulatorServerURL)
		clerkhttp.SetupClerkForTestingMode(config.ClerkEmulatorServerURL)
	}

	eventPublisher, err := events.NewPublisher(config.AmqpUrL, watermill.NewStdLogger(!config.ProductionMode, !config.ProductionMode))
	if err != nil {
		return err
	}

	txProvider := transaction.NewPsqlProvider(db, eventPublisher, logger)

	fmt.Println(txProvider)

	application := &app.App{
		Command: app.Command{},
	}

	httpserver, err := http.NewServer(http.Config{
		Ctx:               ctx,
		Logger:            logger,
		Application:       application,
		Port:              config.Port,
		IsDebug:           !config.ProductionMode,
		ClerkSecretKey:    config.ClerkSecretKey,
		AllowedCorsOrigin: strings.Split(config.AllowedCorsOrigin, ","),
	})
	if err != nil {
		return err
	}

	g.Go(func() error {
		return httpserver.Start()
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