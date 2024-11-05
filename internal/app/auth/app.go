package auth

import (
	"fmt"
	authgrpc "github.com/vaihdass/enkauth/internal/grpc/auth"
	"github.com/vaihdass/enkauth/internal/platform/logs"
	"google.golang.org/grpc"
	stdlog "log"
	"log/slog"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
	log        *slog.Logger
}

func New(port int, log *slog.Logger) (*App, error) {
	const op = "internal.app.auth.New"

	s := grpc.NewServer()

	h, err := authgrpc.NewHandler()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	authgrpc.Register(s, h)

	return &App{
		gRPCServer: s,
		log:        log,
		port:       port,
	}, nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		stdlog.Fatal(err)
	}
}

func (a *App) Run() error {
	const op = "internal.app.auth.Run"
	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Warn("auth gRPC server failed to start: failed to listen to tcp", logs.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("auth gRPC server is running", slog.String("addr", l.Addr().String()))

	if err = a.gRPCServer.Serve(l); err != nil {
		log.Warn("auth gRPC server failed to serve", logs.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "internal.app.auth.Stop"

	logs.WithOp(a.log, op).Info("stopping auth gRPC server", slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()
	logs.WithOp(a.log, op).Info("stopped auth gRPC server", slog.Int("port", a.port))
}
