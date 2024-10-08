package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	libhttp "net/http"
	"os"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
	"github.com/subscribeddotdev/subscribed/server/internal/app"
	"github.com/subscribeddotdev/subscribed/server/internal/common/logs"
)

const (
	akSecurityScheme  = "ApiKeyAuth"
	jwtSecurityScheme = "BearerAuth"
)

type Server struct {
	logger *logs.Logger
	s      *libhttp.Server
}

type Config struct {
	Application        *app.App
	Port               int
	AllowedCorsOrigin  []string
	Logger             *logs.Logger
	IsDebug            bool
	Ctx                context.Context
	JwtSecret          string
	WebFrontendEnabled bool
}

func NewServer(config Config) (*Server, error) {
	router := echo.New()

	spec, err := GetSwagger()
	if err != nil {
		return nil, err
	}

	routerHandlers := &handlers{
		application: config.Application,
		jwtSecret:   config.JwtSecret,
	}

	registerMiddlewares(router, spec, config)
	RegisterHandlers(router, routerHandlers)

	if config.WebFrontendEnabled {
		registerWebFrontend(router)
	}

	return &Server{
		logger: config.Logger,
		s: &libhttp.Server{
			Handler:           router,
			ReadTimeout:       time.Second * 30,
			ReadHeaderTimeout: time.Second * 30,
			WriteTimeout:      time.Second * 30,
			IdleTimeout:       time.Second * 30,
			Addr:              fmt.Sprintf(":%d", config.Port),
			BaseContext: func(listener net.Listener) context.Context {
				return config.Ctx
			},
		},
	}, nil
}

func (s *Server) Start() error {
	s.logger.Info("http server is running", "port", s.s.Addr)
	err := s.s.ListenAndServe()
	if err != nil && !errors.Is(err, libhttp.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("shutting down the http server")
	return s.s.Shutdown(ctx)
}

func registerMiddlewares(router *echo.Echo, spec *openapi3.T, config Config) {
	router.HTTPErrorHandler = errorHandler(config.Logger)
	router.Use(middleware.RequestID())
	router.Use(middleware.Recover())
	router.Use(loggerMiddleware(config.Logger, config.IsDebug))
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.AllowedCorsOrigin,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))

	authApiKeyMiddleware := apiKeyMiddleware{auth: config.Application.Authorization}
	authJwtMiddleware := jwtMiddleware{secret: config.JwtSecret}

	spec.Servers = nil
	router.Use(oapimiddleware.OapiRequestValidatorWithOptions(spec, &oapimiddleware.Options{
		ErrorHandler: nil,
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
				if input.SecuritySchemeName == akSecurityScheme {
					return authApiKeyMiddleware.Middleware(ctx, input)
				}

				if input.SecuritySchemeName == jwtSecurityScheme {
					return authJwtMiddleware.Middleware(ctx, input)
				}

				return fmt.Errorf("unable to recognise '%s' as the security scheme", input.SecuritySchemeName)
			},
		},
		ParamDecoder: nil,
		UserData:     nil,
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/web")
		},
		MultiErrorHandler:     nil,
		SilenceServersWarning: false,
	}))
}
func registerWebFrontend(router *echo.Echo) {
	wd, _ := os.Getwd()

	web := router.Group("/web")
	web.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  fmt.Sprintf("%s/web", wd),
		Index: "index.html",
		HTML5: true,
	}))
}
