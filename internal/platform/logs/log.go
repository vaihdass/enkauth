package logs

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func MustSetupLogger(envType string) *slog.Logger {
	l, err := SetupLogger(envType)
	if err != nil {
		log.Fatal(err)
	}

	return l
}

func SetupLogger(envType string) (*slog.Logger, error) {
	var logger *slog.Logger

	switch envType {
	case envLocal:
		logger = slog.New(
			NewPrettyHandler(os.Stdout, slog.LevelDebug),
		)
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)
	default:
		return nil, fmt.Errorf("internal.platform.SetupLogger: invalid environment type: %s", envType)
	}

	return logger, nil
}
