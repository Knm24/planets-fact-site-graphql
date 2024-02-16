package db

import (
	"log/slog"

	"github.com/talgat-ruby/planets-fact-site-graphql/cmd/db/model"
	"github.com/talgat-ruby/planets-fact-site-graphql/cmd/db/types"
	"github.com/talgat-ruby/planets-fact-site-graphql/configs"
)

func New(log *slog.Logger, conf *configs.DBConfig) (types.DB, error) {
	return model.New(log, conf)
}
