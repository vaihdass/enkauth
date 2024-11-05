package main

import (
	"context"
	"github.com/vaihdass/enkauth/internal/app"
	"github.com/vaihdass/enkauth/internal/platform"
	"github.com/vaihdass/enkauth/internal/platform/logs"
	"log"
	"log/slog"
	"os/signal"
	"syscall"
)

func main() {
	cfg := platform.MustLoadConfig()

	logger := logs.MustSetupLogger(cfg.Env)

	logger.Info("starting auth gRPC service application",
		slog.String("op", "cmd.auth.main"),
		slog.String("environment", cfg.Env),
		slog.Int("port", cfg.GRPC.Port),
	)

	application, err := app.New(cfg, logger)
	if err != nil {
		log.Fatal(err)
	}

	go application.GRPCApp.MustRun()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	<-ctx.Done()
	logger.Info("gracefully shutting down auth gRPC service application")
	application.GRPCApp.Stop()
	logger.Info("auth gRPC service application stopped")
}
