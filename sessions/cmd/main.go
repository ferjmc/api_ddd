package main

import (
	"log"
	"os"

	"github.com/ferjmc/sessions/config"
	"github.com/ferjmc/sessions/internal/server"
	"github.com/ferjmc/sessions/pkg/logger"
	"github.com/ferjmc/sessions/pkg/redis"
)

func main() {
	configPath := config.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Info("Starting sessions server")
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s",
		cfg.GRPCServer.AppVersion,
		cfg.Logger.Level,
		cfg.GRPCServer.Mode,
	)

	appLogger.Infof("Success parsed config: %#v", cfg.GRPCServer.AppVersion)

	redisClient := redis.NewRedisClient(cfg)
	appLogger.Info("Redis connected")

	appLogger.Infof("%-v", redisClient.PoolStats())

	s := server.NewSessionsServer(appLogger, cfg, redisClient)

	appLogger.Fatal(s.Run())
}
