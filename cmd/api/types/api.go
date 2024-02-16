package types

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"

	dbT "github.com/talgat-ruby/interactive-comments-api/cmd/db/types"
)

type Api interface {
	Start(ctx context.Context, cancel context.CancelFunc, d dbT.DB)
	GetLog() *slog.Logger
}

type Middleware interface {
	Logger(ctx context.Context, app *echo.Echo)
}
