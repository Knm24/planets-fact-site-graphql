package model

import (
	"context"
	"fmt"
)

type UpsertLikeInput struct {
	Author    *string
	CommentID *int
	Rate      *int
}

func (m *Model) UpsertLike(ctx context.Context, input *UpsertLikeInput) error {
	m.log.InfoContext(ctx, "start UpsertLike")

	sqlStatement := `
		INSERT INTO like_ (author, comment_id, rate)
		SELECT ?, ?, ?
		WHERE NOT EXISTS (
			SELECT * FROM comment c WHERE c.id = ? AND c.author = ?
		)
		ON CONFLICT(author, comment_id)
			DO UPDATE SET rate = ?;
	`

	res, err := m.db.ExecContext(
		ctx,
		sqlStatement,
		input.Author,
		input.CommentID,
		input.Rate,
		input.CommentID,
		input.Author,
		input.Rate,
	)
	if err != nil {
		m.log.ErrorContext(ctx, "fail UpsertLike", "error", err)
		return err
	}

	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if n == 0 {
		return fmt.Errorf("no record was inserted, please check your request")
	}

	m.log.InfoContext(ctx, "success UpsertLike")
	return nil
}
