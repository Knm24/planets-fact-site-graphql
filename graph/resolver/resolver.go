package resolver

import (
	"log/slog"

	dbT "github.com/talgat-ruby/planets-fact-site-graphql/cmd/db/types"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Log *slog.Logger
	DB  dbT.DB
}
