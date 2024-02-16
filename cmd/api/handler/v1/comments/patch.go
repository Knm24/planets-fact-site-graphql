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

type PatchRequestParam struct {
	ID *int `param:"id" validate:"required,gt=0"`
}

type PatchRequestQuery struct {
	User *string `query:"user" validate:"required"`
}

type PatchRequestBody struct {
	Content string `xml:"content" json:"content" form:"content" validate:"required"`
}

func (h *Handler) Edit(c echo.Context) error {
	ctx := c.Request().Context()
	h.log.InfoContext(ctx, "start Edit", "path", c.Path())

	reqParam := new(PatchRequestParam)
	if err := (&echo.DefaultBinder{}).BindPathParams(c, reqParam); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Delete:: body binding error",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	if err := h.patchRequestParamValidationErrors(ctx, reqParam); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Delete:: validation errors",
			"path", c.Path(),
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	reqQuery := new(PatchRequestQuery)
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, reqQuery); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Edit:: body binding error",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	if err := h.patchRequestQueryValidationErrors(ctx, reqQuery); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Delete:: validation errors",
			"path", c.Path(),
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	reqBody, err := h.patchRequestBody(ctx, c)
	if err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Edit:: body binding error",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	if err := h.patchRequestValidationErrors(ctx, reqBody); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Edit:: validation errors",
			"path", c.Path(),
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	dbInput := patchDBInput(reqBody, reqParam.ID, reqQuery.User)
	if err := h.db.UpdateComment(ctx, dbInput); err != nil {
		h.log.ErrorContext(
			ctx,
			"fail Edit:: db add fail",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	h.log.InfoContext(ctx, "success Edit", "path", c.Path())
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) patchRequestParamValidationErrors(_ context.Context, reqParam *PatchRequestParam) error {
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

func (h *Handler) patchRequestQueryValidationErrors(_ context.Context, reqParam *PatchRequestQuery) error {
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

func (h *Handler) patchRequestBody(_ context.Context, c echo.Context) (*PatchRequestBody, error) {
	reqBody := new(PatchRequestBody)
	if err := (&echo.DefaultBinder{}).BindBody(c, reqBody); err != nil {
		return nil, err
	}

	return reqBody, nil
}

func (h *Handler) patchRequestValidationErrors(_ context.Context, reqBody *PatchRequestBody) error {
	if err := h.validate.Struct(reqBody); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.StructField() {
			case "Content":
				return fmt.Errorf("content is required")
			}
		}

		return err
	}

	return nil
}

func patchDBInput(reqBody *PatchRequestBody, id *int, username *string) *model.UpdateCommentInput {
	inp := new(model.UpdateCommentInput)

	if reqBody == nil {
		return inp
	}

	inp.ID = id
	inp.Author = username
	inp.Content = reqBody.Content

	return inp
}
