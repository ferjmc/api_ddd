package main

import (
	"log"
	"os"

	"github.com/ferjmc/api_ddd/user/config"
	"github.com/ferjmc/api_ddd/user/internal/server"
	"github.com/ferjmc/api_ddd/user/pkg/logger"
	"github.com/ferjmc/api_ddd/user/pkg/postgres"
)

func main() {
	configPath := config.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Info("Starting user server")
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s",
		cfg.GRPCServer.AppVersion,
		cfg.Logger.Level,
		cfg.GRPCServer.Mode,
	)
	appLogger.Infof("Success parsed config: %#v", cfg.GRPCServer.AppVersion)

	pgxConn, err := postgres.NewPgxConn(cfg, appLogger)
	if err != nil {
		appLogger.Fatal("cannot connect to postgres", err)
	}
	defer pgxConn.Close()

	appLogger.Infof("%-v", pgxConn.Stat())

	s := server.NewServer(appLogger, cfg, pgxConn)
	appLogger.Fatal(s.Run())

}
