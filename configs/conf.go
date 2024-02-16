package configs

import (
	"context"
	"flag"

	"github.com/talgat-ruby/interactive-comments-api/internal/constant"
)

type Config struct {
	Env constant.Environment
	Api *ApiConfig
	DB  *DBConfig
}

func NewConfig(ctx context.Context) (*Config, error) {
	conf := &Config{
		Env: getEnv(),
	}

	_ = conf.loadDotEnvFiles()

	// Api config
	if c, err := newApiConfig(ctx, conf.Env); err != nil {
		return nil, err
	} else {
		conf.Api = c
	}

	// DB config
	if c, err := newDBConfig(ctx); err != nil {
		return nil, err
	} else {
		conf.DB = c
	}

	flag.Parse()

	return conf, nil
}
