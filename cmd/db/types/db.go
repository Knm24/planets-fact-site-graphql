package types

import (
	"context"

	"github.com/talgat-ruby/interactive-comments-api/cmd/db/model"
)

type DB interface {
	ReadComments(ctx context.Context, username string) ([]*model.Comment, error)
	CreateComment(ctx context.Context, input *model.CreateCommentInput) error
	UpdateComment(ctx context.Context, input *model.UpdateCommentInput) error
	DeleteComment(ctx context.Context, input *model.DeleteCommentInput) error
	UpsertLike(ctx context.Context, input *model.UpsertLikeInput) error
}
