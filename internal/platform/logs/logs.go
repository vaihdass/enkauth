package logs

import "log/slog"

const operationLogKey = "op"

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func WithOp(logger *slog.Logger, operation string) *slog.Logger {
	return logger.With(slog.String(operationLogKey, operation))
}
