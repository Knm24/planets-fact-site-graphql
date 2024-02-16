package router

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/talgat-ruby/interactive-comments-api/cmd/api/middleware"
	apiT "github.com/talgat-ruby/interactive-comments-api/cmd/api/types"
	dbT "github.com/talgat-ruby/interactive-comments-api/cmd/db/types"
)

// SetupRoutes setup router api
func SetupRoutes(ctx context.Context, app *echo.Echo, api apiT.Api, db dbT.DB, v *validator.Validate) {
	m := middleware.New(api, db)
	m.Logger(ctx, app)

	group := app.Group("/api")
	v1Group(group, db, v, api.GetLog())
}
