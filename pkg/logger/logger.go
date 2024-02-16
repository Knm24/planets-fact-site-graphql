package logger

import (
	"log/slog"
	"os"
)

func New(isLocal bool) *slog.Logger {
	if isLocal {

		opts := &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}
		return slog.New(slog.NewTextHandler(os.Stdout, opts))
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
