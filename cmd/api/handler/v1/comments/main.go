package comments

import (
	"log/slog"

	"github.com/go-playground/validator/v10"

	dbT "github.com/talgat-ruby/interactive-comments-api/cmd/db/types"
)

type Handler struct {
	db       dbT.DB
	validate *validator.Validate
	log      *slog.Logger
}

func New(db dbT.DB, v *validator.Validate, l *slog.Logger) *Handler {
	return &Handler{db, v, l}
}
