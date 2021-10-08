package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/ferjmc/user/config"
	"github.com/ferjmc/user/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	maxConn           = 50
	healthCheckPeriod = 3 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
	lazyConnect       = false
)

func NewPgxConn(cfg *config.Config, logger logger.Logger) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlPort,
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlDbname,
		cfg.Postgres.PostgresqlSSLMode,
		cfg.Postgres.PostgresqlPassword,
	)

	poolCfg, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConns = maxConn
	poolCfg.HealthCheckPeriod = healthCheckPeriod
	poolCfg.MaxConnIdleTime = maxConnIdleTime
	poolCfg.MaxConnLifetime = maxConnLifetime
	poolCfg.MinConns = minConns
	poolCfg.LazyConnect = lazyConnect
	poolCfg.ConnConfig.Logger = logger

	connPool, err := pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("pgx.ConnectConfig %w", err)
	}

	return connPool, nil
}
