package likes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/talgat-ruby/interactive-comments-api/cmd/db/model"
	"github.com/talgat-ruby/interactive-comments-api/internal/response"
)

type PostRequestQuery struct {
	User *string `query:"user" validate:"required"`
}

type PostRequestBody struct {
	CommentID *int `xml:"commentId" json:"commentId,omitempty" form:"commentId" validate:"required"`
	Rate      *int `xml:"rate" json:"rate,omitempty" form:"rate" validate:"required,oneof=1 0 -1"`
}

func (h *Handler) AddOrEdit(c echo.Context) error {
	ctx := c.Request().Context()
	h.log.InfoContext(ctx, "start AddOrEdit", "path", c.Path())

	reqQuery := new(PostRequestQuery)
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, reqQuery); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail AddOrEdit:: body binding error",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	if err := h.postRequestQueryValidationErrors(ctx, reqQuery); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Delete:: validation errors",
			"path", c.Path(),
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	reqBody, err := h.postRequestBody(ctx, c)
	if err != nil {
		h.log.ErrorContext(
			ctx,
			"fail AddOrEdit:: body binding error",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	if err := h.postRequestValidationErrors(ctx, reqBody); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail AddOrEdit:: validation errors",
			"path", c.Path(),
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	dbInput := postDBInput(reqBody, reqQuery.User)
	if err := h.db.UpsertLike(ctx, dbInput); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail AddOrEdit:: db add fail",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	h.log.InfoContext(ctx, "success AddOrEdit", "path", c.Path())
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) postRequestQueryValidationErrors(_ context.Context, reqParam *PostRequestQuery) error {
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

func (h *Handler) postRequestBody(_ context.Context, c echo.Context) (*PostRequestBody, error) {
	reqBody := new(PostRequestBody)
	if err := (&echo.DefaultBinder{}).BindBody(c, reqBody); err != nil {
		return nil, err
	}

	return reqBody, nil
}

func (h *Handler) postRequestValidationErrors(_ context.Context, reqBody *PostRequestBody) error {
	if err := h.validate.Struct(reqBody); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.StructField() {
			case "CommentID":
				return fmt.Errorf("commentId is invalid")
			case "Rate":
				return fmt.Errorf("rate is invalid")
			}
		}

		return err
	}

	return nil
}

func postDBInput(reqBody *PostRequestBody, username *string) *model.UpsertLikeInput {
	inp := new(model.UpsertLikeInput)

	if reqBody == nil {
		return inp
	}

	inp.Author = username
	inp.CommentID = reqBody.CommentID
	inp.Rate = reqBody.Rate

	return inp
}
