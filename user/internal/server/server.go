package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ferjmc/api_ddd/user/config"
	"github.com/ferjmc/api_ddd/user/internal/interceptors"
	"github.com/ferjmc/api_ddd/user/internal/middlewares"
	"github.com/ferjmc/api_ddd/user/internal/session_client"
	userHandlers "github.com/ferjmc/api_ddd/user/internal/user/delivery/http"
	"github.com/ferjmc/api_ddd/user/internal/user/repository"
	"github.com/ferjmc/api_ddd/user/internal/user/usecase"
	"github.com/ferjmc/api_ddd/user/pkg/logger"
	sessionService "github.com/ferjmc/api_ddd/user/proto/session"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

const (
	certFile          = "ssl/server.crt"
	keyFile           = "ssl/server.pem"
	maxHeaderBytes    = 1 << 20
	userCachePrefix   = "users:"
	userCacheDuration = time.Minute * 15
)

// Server
type Server struct {
	echo   *echo.Echo
	logger logger.Logger
	cfg    *config.Config
	//redisConn *redis.Client
	pgxPool *pgxpool.Pool
	//tracer    opentracing.Tracer
}

// NewServer
func NewServer(logger logger.Logger, cfg *config.Config, pgxPool *pgxpool.Pool) *Server {
	return &Server{logger: logger, cfg: cfg, pgxPool: pgxPool, echo: echo.New()}
}

func (s *Server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	validate := validator.New()
	v1 := s.echo.Group("/api/v1")
	usersGroup := v1.Group("/users")
	im := interceptors.NewInterceptorManager(s.logger, s.cfg)

	sessGRPCConn, err := session_client.NewSessionServiceConn(ctx, s.cfg, im)
	if err != nil {
		return fmt.Errorf("sessionclient.NewSessionServiceConn: %w", err)
	}
	defer sessGRPCConn.Close()

	sessServiceClient := sessionService.NewAuthorizationServiceClient(sessGRPCConn)

	userPGRepository := repository.NewUserPGRepository(s.pgxPool)
	userUseCase := usecase.NewUserUseCase(userPGRepository, sessServiceClient, s.logger)

	middlewareManager := middlewares.NewMiddlewareManager(s.logger, s.cfg, userUseCase)

	uh := userHandlers.NewUserHandlers(usersGroup, userUseCase, s.logger, validate, middlewareManager, s.cfg)
	uh.MapUserRoutes()

	s.echo.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.HttpServer.Port)
		s.echo.Server.ReadTimeout = time.Second * s.cfg.HttpServer.ReadTimeout
		s.echo.Server.WriteTimeout = time.Second * s.cfg.HttpServer.WriteTimeout
		s.echo.Server.MaxHeaderBytes = maxHeaderBytes
		if err := s.echo.Start(s.cfg.HttpServer.Port); err != nil {
			s.logger.Fatalf("Error starting TLS Server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		s.logger.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		s.logger.Errorf("ctx.Done: %v", done)
	}

	s.logger.Info("Server Exited Properly")

	if err := s.echo.Server.Shutdown(ctx); err != nil {
		return fmt.Errorf("echo.Server.Shutdown: %w", err)
	}

	s.logger.Info("Server Exited Properly")

	return nil
}
