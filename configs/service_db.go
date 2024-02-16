package configs

import (
	"context"
	"flag"

	"github.com/sethvargo/go-envconfig"
)

type DBConfig struct {
	DBFile string `env:"DB_FILE"`
}

func newDBConfig(ctx context.Context) (*DBConfig, error) {
	c := &DBConfig{}

	if err := envconfig.Process(ctx, c); err != nil {
		return nil, err
	}

	flag.StringVar(&c.DBFile, "db-file", c.DBFile, "database db-file [DB_FILE]")

	return c, nil
}
