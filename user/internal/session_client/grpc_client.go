package session_client

import (
	"context"
	"time"

	"github.com/ferjmc/api_ddd/user/config"
	"github.com/ferjmc/api_ddd/user/internal/interceptors"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	backoffLinear = 100 * time.Millisecond
)

func NewSessionServiceConn(ctx context.Context, cfg *config.Config, manager *interceptors.InterceptorManager) (*grpc.ClientConn, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(backoffLinear)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted),
	}

	sessGRPCConn, err := grpc.DialContext(
		ctx,
		cfg.GRPCServer.SessionGrpcServicePort,
		grpc.WithUnaryInterceptor(manager.GetInterceptor()),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		return nil, err
	}

	return sessGRPCConn, nil
}
