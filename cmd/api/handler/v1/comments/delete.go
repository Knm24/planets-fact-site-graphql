package comments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/talgat-ruby/interactive-comments-api/cmd/db/model"
	"github.com/talgat-ruby/interactive-comments-api/internal/response"
)

type DeleteRequestParam struct {
	ID *int `param:"id" validate:"required,gt=0"`
}

type DeleteRequestQuery struct {
	User *string `query:"user" validate:"required"`
}

func (h *Handler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	h.log.InfoContext(ctx, "start Delete", "path", c.Path())

	reqParam := new(DeleteRequestParam)
	if err := (&echo.DefaultBinder{}).BindPathParams(c, reqParam); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Delete:: body binding error",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	if err := h.deleteRequestParamValidationErrors(ctx, reqParam); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Delete:: validation errors",
			"path", c.Path(),
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	reqQuery := new(DeleteRequestQuery)
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, reqQuery); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Delete:: body binding error",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	if err := h.deleteRequestQueryValidationErrors(ctx, reqQuery); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Delete:: validation errors",
			"path", c.Path(),
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	dbInput := deleteDBInput(reqParam.ID, reqQuery.User)
	if err := h.db.DeleteComment(ctx, dbInput); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Delete:: db add fail",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	h.log.InfoContext(ctx, "success Delete", "path", c.Path())
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) deleteRequestParamValidationErrors(_ context.Context, reqParam *DeleteRequestParam) error {
	if err := h.validate.Struct(reqParam); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.StructField() {
			case "ID":
				return fmt.Errorf("id is invalid")
			}
		}

		return err
	}

	return nil
}

func (h *Handler) deleteRequestQueryValidationErrors(_ context.Context, reqParam *DeleteRequestQuery) error {
	if err := h.validate.Struct(reqParam); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.StructField() {
			case "User":
				return fmt.Errorf("user is invalid")
			}
		}

		return err
	}

	return nil
}

func deleteDBInput(id *int, username *string) *model.DeleteCommentInput {
	inp := new(model.DeleteCommentInput)

	inp.ID = id
	inp.Username = username

	return inp
}
