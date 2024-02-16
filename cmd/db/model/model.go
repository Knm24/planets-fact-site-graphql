package model

import (
	"database/sql"
	"log/slog"

	"github.com/talgat-ruby/interactive-comments-api/cmd/db/database"
	"github.com/talgat-ruby/interactive-comments-api/configs"
)

type Model struct {
	log  *slog.Logger
	conf *configs.DBConfig
	db   *sql.DB
}

func New(log *slog.Logger, conf *configs.DBConfig) (*Model, error) {
	db, err := database.NewDB("./database.db?_foreign_keys=on")
	if err != nil {
		return nil, err
	}

	m := &Model{
		log:  log,
		conf: conf,
		db:   db,
	}

	return m, nil
}
