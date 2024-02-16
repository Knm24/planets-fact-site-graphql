package router

import (
	"log/slog"

	"github.com/labstack/echo/v4"

	"github.com/talgat-ruby/planets-fact-site-graphql/cmd/api/handler/graphql"
	dbT "github.com/talgat-ruby/planets-fact-site-graphql/cmd/db/types"
)

func graphqlRouter(app *echo.Echo, db dbT.DB, l *slog.Logger) {
	h := graphql.New(db, l)

	app.GET("/graphql", h.GraphQL)
}
