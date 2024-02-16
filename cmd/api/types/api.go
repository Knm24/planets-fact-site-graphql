package types

import (
	"context"
	"log/slog"

	dbT "github.com/talgat-ruby/planets-fact-site-graphql/cmd/db/types"
)

type Api interface {
	Start(ctx context.Context, cancel context.CancelFunc, d dbT.DB)
	GetLog() *slog.Logger
}

type Middleware interface{}
