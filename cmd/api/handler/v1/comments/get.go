package comments

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/talgat-ruby/interactive-comments-api/cmd/db/model"
	"github.com/talgat-ruby/interactive-comments-api/internal/response"
)

type GetListRequestQuery struct {
	User string `query:"user"`
}

type GetListResponseBody struct {
	Data []*comment `json:"data"`
}

type commentReply struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	AvatarUrl string `json:"avatarUrl"`
	Likes     int    `json:"likes"`
	Duration  string `json:"duration"`
	IsMine    bool   `json:"isMine"`
	MyRate    int    `json:"myRate"`
	Addressee string `json:"addressee"`
}

type comment struct {
	ID        int             `json:"id"`
	Content   string          `json:"content"`
	Author    string          `json:"author"`
	AvatarUrl string          `json:"avatarUrl"`
	Likes     int             `json:"likes"`
	Duration  string          `json:"duration"`
	IsMine    bool            `json:"isMine"`
	MyRate    int             `json:"myRate"`
	Replies   []*commentReply `json:"replies"`
}

func (h *Handler) ReadList(c echo.Context) error {
	ctx := c.Request().Context()
	h.log.InfoContext(ctx, "start ReadList", "path", c.Path())

	reqQuery := new(GetListRequestQuery)
	err := c.Bind(reqQuery)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	comments, err := h.db.ReadComments(ctx, reqQuery.User)
	if err != nil {
		h.log.ErrorContext(
			ctx,
			"fail ReadList:: db add fail",
			"path", c.Path(),
			"error", err,
		)
		return c.JSON(http.StatusBadRequest, response.ErrorWithMessage{Error: response.WithMessage{Message: err.Error()}})
	}

	respBody := mapDBCommentsToRespComments(comments)

	h.log.InfoContext(ctx, "success ReadList", "path", c.Path())
	return c.JSON(http.StatusOK, response.Data{
		Data: respBody,
	})
}

func mapDBCommentsToRespComments(cs []*model.Comment) []*comment {
	respCs := make([]*comment, len(cs))

	for i, c := range cs {
		respCs[i] = mapDBCommentToRespComment(c)
	}

	return respCs
}

func mapDBCommentToRespComment(c *model.Comment) *comment {
	return &comment{
		ID:        c.ID,
		Content:   c.Content,
		Author:    c.Author,
		AvatarUrl: c.AvatarUrl,
		Likes:     c.Likes,
		Duration:  c.Duration,
		IsMine:    c.IsMine,
		MyRate:    c.MyRate,
		Replies:   mapDBCommentRepliesToRespCommentReplies(c.Replies),
	}
}

func mapDBCommentRepliesToRespCommentReplies(cs []*model.Reply) []*commentReply {
	respCs := make([]*commentReply, len(cs))

	for i, c := range cs {
		respCs[i] = mapDBCommentReplyToRespCommentReply(c)
	}

	return respCs
}

func mapDBCommentReplyToRespCommentReply(c *model.Reply) *commentReply {
	return &commentReply{
		ID:        c.ID,
		Content:   c.Content,
		Author:    c.Author,
		AvatarUrl: c.AvatarUrl,
		Likes:     c.Likes,
		Duration:  c.Duration,
		IsMine:    c.IsMine,
		MyRate:    c.MyRate,
		Addressee: c.Addressee,
	}
}
