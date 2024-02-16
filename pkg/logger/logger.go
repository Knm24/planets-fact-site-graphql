package logger

import (
	"log"
	"log/slog"
	"os"
)

func New(isLocal bool) *slog.Logger {
	if isLocal {
		opts := PrettyHandlerOptions{
			SlogOpts: slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		}
		h := &PrettyHandler{
			h: slog.NewJSONHandler(os.Stdout, &opts.SlogOpts),
			l: log.New(os.Stdout, "", 0),
		}
		return slog.New(h)
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
