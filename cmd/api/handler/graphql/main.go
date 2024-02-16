package graphql

import (
	"log/slog"

	dbT "github.com/talgat-ruby/planets-fact-site-graphql/cmd/db/types"
)

type Handler struct {
	db  dbT.DB
	log *slog.Logger
}

func New(db dbT.DB, l *slog.Logger) *Handler {
	return &Handler{db, l}
}
