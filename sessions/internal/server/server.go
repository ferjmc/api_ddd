package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"github.com/ferjmc/sessions/config"
	crfRepository "github.com/ferjmc/sessions/internal/csrf/repository"
	csrfUseCase "github.com/ferjmc/sessions/internal/csrf/usecase"
	"github.com/ferjmc/sessions/internal/interceptors"
	"github.com/ferjmc/sessions/internal/session/delivery"
	"github.com/ferjmc/sessions/internal/session/repository"
	"github.com/ferjmc/sessions/internal/session/usecase"
	"github.com/ferjmc/sessions/pkg/logger"
	sessionService "github.com/ferjmc/sessions/proto"
)

// Server
type Server struct {
	logger    logger.Logger
	cfg       *config.Config
	redisConn *redis.Client
}

// NewServer
func NewSessionsServer(logger logger.Logger, cfg *config.Config, redisConn *redis.Client) *Server {
	return &Server{logger: logger, cfg: cfg, redisConn: redisConn}
}

func (s *Server) Run() error {
	ctx := context.Background()

	im := interceptors.NewInterceptorManager(s.logger, s.cfg)
	sessionRedisRepo := repository.NewSessionRedisRepo(s.redisConn, s.cfg.GRPCServer.SessionPrefix, time.Duration(s.cfg.GRPCServer.SessionExpire)*time.Minute)
	sessionUC := usecase.NewSessionUseCase(sessionRedisRepo)
	csrfRepository := crfRepository.NewCsrfRepository(s.redisConn, s.cfg.GRPCServer.CSRFPrefix, time.Duration(s.cfg.GRPCServer.CsrfExpire)*time.Minute)
	csrfUC := csrfUseCase.NewCsrfUseCase(csrfRepository)

	l, err := net.Listen("tcp", s.cfg.GRPCServer.Port)
	if err != nil {
		return err
	}
	defer l.Close()

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.cfg.GRPCServer.MaxConnectionIdle * time.Minute,
		Timeout:           s.cfg.GRPCServer.Timeout * time.Second,
		MaxConnectionAge:  s.cfg.GRPCServer.MaxConnectionAge * time.Minute,
		Time:              s.cfg.GRPCServer.Timeout * time.Minute,
	}),
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
			im.Logger,
		),
	)

	sessGRPCService := delivery.NewSessionsService(s.logger, sessionUC, csrfUC)
	sessionService.RegisterAuthorizationServiceServer(server, sessGRPCService)

	go func() {
		s.logger.Infof("GRPC Server is listening on port: %v", s.cfg.GRPCServer.Port)
		s.logger.Fatal(server.Serve(l))
	}()

	if s.cfg.GRPCServer.Mode != "Production" {
		reflection.Register(server)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		s.logger.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		s.logger.Errorf("ctx.Done: %v", done)
	}

	server.GracefulStop()
	s.logger.Info("Server Exited Properly")

	return nil
}
