package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"

	"github.com/talgat-ruby/interactive-comments-api/cmd/api/router"
	apiT "github.com/talgat-ruby/interactive-comments-api/cmd/api/types"
	dbT "github.com/talgat-ruby/interactive-comments-api/cmd/db/types"
	"github.com/talgat-ruby/interactive-comments-api/configs"
	"github.com/talgat-ruby/interactive-comments-api/internal/validator"
)

type server struct {
	log  *slog.Logger
	conf *configs.ApiConfig
	app  *echo.Echo
}

func New(log *slog.Logger, conf *configs.ApiConfig) apiT.Api {
	s := &server{
		log:  log,
		conf: conf,
	}

	return s
}

func (s *server) GetLog() *slog.Logger {
	return s.log
}

func (s *server) Start(ctx context.Context, cancel context.CancelFunc, db dbT.DB) {
	v := validator.New()
	e := echo.New()

	e.Logger.SetOutput(io.Discard)

	srv := http.Server{
		Addr:        fmt.Sprintf(":%d", s.conf.Port),
		Handler:     e,
		IdleTimeout: s.conf.IdleTimeout,
	}

	router.SetupRoutes(ctx, e, s, db, v)

	// Listen from s different goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.ErrorContext(ctx, "server error", "error", err)
		}

		cancel()
	}()

	s.log.InfoContext(ctx, "start server", "PORT", s.conf.Port)

	shutdown := make(chan os.Signal, 1)                    // Create channel to signify s signal being sent
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM) //  When an interrupt or termination signal is sent, notify the channel

	go func() {
		_ = <-shutdown

		s.log.WarnContext(ctx, "gracefully shutting down...")
		if err := srv.Shutdown(ctx); err != nil {
			s.log.ErrorContext(ctx, "server shutdown error", "error", err)
		}
	}()
}
