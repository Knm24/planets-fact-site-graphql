package db

import (
	"log/slog"

	"github.com/talgat-ruby/interactive-comments-api/cmd/db/model"
	"github.com/talgat-ruby/interactive-comments-api/cmd/db/types"
	"github.com/talgat-ruby/interactive-comments-api/configs"
)

func New(log *slog.Logger, conf *configs.DBConfig) (types.DB, error) {
	return model.New(log, conf)
}
