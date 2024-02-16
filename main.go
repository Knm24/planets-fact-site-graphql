package main

import (
	"context"
	"log/slog"

	"github.com/talgat-ruby/planets-fact-site-graphql/cmd/api"
	"github.com/talgat-ruby/planets-fact-site-graphql/cmd/db"
	"github.com/talgat-ruby/planets-fact-site-graphql/configs"
	"github.com/talgat-ruby/planets-fact-site-graphql/internal/constant"
	"github.com/talgat-ruby/planets-fact-site-graphql/pkg/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// conf
	conf, err := configs.NewConfig(ctx)
	if err != nil {
		panic(err)
	}

	// setup logger
	log := logger.New(conf.Env == constant.EnvironmentLocal)

	// configure db service
	d, err := db.New(log.WithGroup(string(constant.DB)).With("service", constant.DB), conf.DB)
	if err != nil {
		log.LogAttrs(
			ctx, slog.LevelError, "initialize service error",
			slog.String("service", "database"),
			slog.Any("error", err),
		)
		panic(err)
	}
	log.LogAttrs(
		ctx, slog.LevelInfo, "initialize service",
		slog.String("service", "database"),
	)

	// configure gateway service
	srv := api.New(log.With("service", constant.Api), conf.Api)
	log.LogAttrs(
		ctx, slog.LevelInfo, "initialize service",
		slog.String("service", "api"),
	)
	// start gateway service
	srv.Start(ctx, cancel, d)

	<-ctx.Done()
	// Your cleanup tasks go here
	log.LogAttrs(
		ctx, slog.LevelInfo, "cleaning up ...",
	)

	log.LogAttrs(
		ctx, slog.LevelInfo, "server was successful shutdown.",
	)
}
