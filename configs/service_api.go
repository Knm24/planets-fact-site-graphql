package configs

import (
	"context"
	"flag"
	"time"

	"github.com/sethvargo/go-envconfig"

	"github.com/talgat-ruby/interactive-comments-api/internal/constant"
)

type ApiConfig struct {
	Env         constant.Environment
	Host        string        `env:"HOST,default=localhost"`
	Port        int           `env:"PORT,default=8081"`
	IdleTimeout time.Duration `env:"IDLE_TIMEOUT"`
}

func newApiConfig(ctx context.Context, env constant.Environment) (*ApiConfig, error) {
	c := &ApiConfig{
		Env: env,
	}

	if err := envconfig.Process(ctx, c); err != nil {
		return nil, err
	}

	flag.StringVar(&c.Host, "host", c.Host, "server host [HOST]")
	flag.IntVar(&c.Port, "port", c.Port, "server port [PORT]")
	flag.DurationVar(
		&c.IdleTimeout,
		"idle-timeout",
		c.IdleTimeout,
		"expiration period for access token, use \"10m\", \"3s\" etc [IDLE_TIMEOUT]",
	)

	return c, nil
}
