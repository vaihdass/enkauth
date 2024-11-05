package app

import (
	"fmt"
	"github.com/vaihdass/enkauth/internal/app/auth"
	"github.com/vaihdass/enkauth/internal/platform"
	"log/slog"
)

type App struct {
	GRPCApp *auth.App
}

func New(cfg *platform.Config, log *slog.Logger) (*App, error) {
	const op = "internal.app.New"

	// TODO: Init storage

	// TODO: init auth service

	gRPCApp, err := auth.New(cfg.GRPC.Port, log)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &App{
		GRPCApp: gRPCApp,
	}, nil
}
