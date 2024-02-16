package router

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/talgat-ruby/interactive-comments-api/cmd/api/handler/v1/comments"
	"github.com/talgat-ruby/interactive-comments-api/cmd/api/handler/v1/likes"
	dbT "github.com/talgat-ruby/interactive-comments-api/cmd/db/types"
)

func v1Group(api *echo.Group, db dbT.DB, v *validator.Validate, l *slog.Logger) {
	g := api.Group("/v1")

	v1formsRouter(g, db, v, l)
	v1likesRouter(g, db, v, l)
}

func v1formsRouter(v1 *echo.Group, db dbT.DB, v *validator.Validate, l *slog.Logger) {
	h := comments.New(db, v, l)

	v1.GET("/comments", h.ReadList)
	v1.POST("/comments", h.Add)
	v1.PATCH("/comments/:id", h.Edit)
	v1.DELETE("/comments/:id", h.Delete)
}

func v1likesRouter(v1 *echo.Group, db dbT.DB, v *validator.Validate, l *slog.Logger) {
	h := likes.New(db, v, l)

	v1.POST("/likes", h.AddOrEdit)
}
