package router

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	apiT "github.com/talgat-ruby/planets-fact-site-graphql/cmd/api/types"
	dbT "github.com/talgat-ruby/planets-fact-site-graphql/cmd/db/types"
)

// SetupRoutes setup router api
func SetupRoutes(ctx context.Context, app *echo.Echo, api apiT.Api, db dbT.DB, v *validator.Validate) {
	graphqlRouter(app, db, api.GetLog())
}
