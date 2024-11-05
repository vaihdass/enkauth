package logs

import (
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"io"
	stdlog "log"
	"log/slog"
)

// PrettyHandler only for local environment use
type PrettyHandler struct {
	l     *stdlog.Logger
	level slog.Level
	attrs []slog.Attr
}

// NewPrettyHandler only for local environment use
func NewPrettyHandler(out io.Writer, level slog.Level) *PrettyHandler {
	return &PrettyHandler{
		l:     stdlog.New(out, "", 0),
		level: level,
	}
}

func (h *PrettyHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var b []byte
	var err error

	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(string(b)),
	)

	return nil
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{
		l:     h.l,
		attrs: attrs,
	}
}

// WithGroup returns the same handler, doesn't supported by this implementation
func (h *PrettyHandler) WithGroup(string) slog.Handler {
	return h
}
