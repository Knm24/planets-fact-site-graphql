package model

import (
	"context"
	"fmt"
)

type DBComment struct {
	ID        int
	Content   string
	Author    string
	AvatarUrl string
	Likes     int
	Duration  string
	IsMine    bool
	MyRate    int
	ParentID  *int
	Addressee *string
}

type Reply struct {
	ID        int
	Content   string
	Author    string
	AvatarUrl string
	Likes     int
	Duration  string
	IsMine    bool
	MyRate    int
	Addressee string
}

type Comment struct {
	ID        int
	Content   string
	Author    string
	AvatarUrl string
	Likes     int
	Duration  string
	IsMine    bool
	MyRate    int
	Replies   []*Reply
}

func (m *Model) ReadComments(ctx context.Context, username string) ([]*Comment, error) {
	m.log.InfoContext(ctx, "start ReadComments")

	parentIds, err := m.getParentCommentsId(ctx, username)
	if err != nil {
		m.log.ErrorContext(ctx, "fail ReadComments", "error", err)
		return nil, err
	}

	comments, err := m.getComments(ctx, username, parentIds)
	if err != nil {
		m.log.ErrorContext(ctx, "fail ReadComments", "error", err)
		return nil, err
	}

	m.log.InfoContext(ctx, "success ReadComments")
	return comments, nil
}

func (m *Model) getParentCommentsId(ctx context.Context, username string) ([]*int, error) {
	m.log.InfoContext(ctx, "start getParentCommentsId")

	sqlStatement := `
		SELECT c.OID
		FROM main.comment c
		WHERE c.parent_id IS NULL
	`

	rows, err := m.db.QueryContext(ctx, sqlStatement)
	if err != nil {
		m.log.ErrorContext(ctx, "fail getParentCommentsId", "error", err)
		return nil, err
	}

	defer rows.Close()

	ids := make([]*int, 0)

	for rows.Next() {
		id := new(int)

		if err = rows.Scan(&id); err != nil {
			m.log.ErrorContext(ctx, "fail getParentCommentsId", "error", err)
			return nil, err
		}

		ids = append(ids, id)
	}

	m.log.InfoContext(ctx, "success getParentCommentsId")
	return ids, nil
}

func (m *Model) getComments(ctx context.Context, username string, pIds []*int) ([]*Comment, error) {
	m.log.InfoContext(ctx, "start getComments")

	sqlStatement := `
		SELECT
			c.id as id,
			c.content as content,
			c.author as author,
			c.addressee as addressee,
			CASE
				WHEN (strftime('%Y', 'now') - strftime('%Y', c.created_at)) > 0
					THEN 'More than ' || (strftime('%Y', 'now') - strftime('%Y', c.created_at)) || ' year(s) ago'
				WHEN (strftime('%m', 'now') - strftime('%m', c.created_at)) > 0
					THEN 'More than ' || (strftime('%m', 'now') - strftime('%m', c.created_at)) || ' month(s) ago'
				WHEN (strftime('%d', 'now') - strftime('%d', c.created_at)) > 0
					THEN 'More than ' || (strftime('%d', 'now') - strftime('%d', c.created_at)) || ' day(s) ago'
				WHEN (strftime('%H', 'now') - strftime('%H', c.created_at)) > 0
					THEN 'More than ' || (strftime('%H', 'now') - strftime('%H', c.created_at)) || ' hour(s) ago'
				WHEN (strftime('%M', 'now') - strftime('%M', c.created_at)) > 0
					THEN 'More than ' || (strftime('%M', 'now') - strftime('%M', c.created_at)) || ' minute(s) ago'
				WHEN (strftime('%S', 'now') - strftime('%S', c.created_at)) > 0
					THEN 'More than ' || (strftime('%S', 'now') - strftime('%S', c.created_at)) || ' second(s) ago'
				ELSE 'now'
			END AS duration,
			u.avatar_url as avatar_url,
			u.username == ? as is_mine,
			pc.OID as parent_id,
			CASE
				WHEN l.count is NULL THEN 0
				ELSE l.count
			END AS likes,
			CASE
				WHEN l2.rate is NULL THEN 0
				ELSE l2.rate
			END AS my_rate
		FROM main.comment c
		LEFT JOIN main.user_ u ON c.author = u.username
		LEFT JOIN main.comment pc ON c.parent_id = pc.OID
		LEFT JOIN
			(
				SELECT
					comment_id,
					SUM(rate) as count
				FROM
					like_
				GROUP BY
					comment_id
			) as l
			ON
				c.OID = l.comment_id
				LEFT JOIN
			like_ l2
			ON
				c.OID == l2.comment_id AND l2.author == ?
		ORDER BY c.created_at DESC;
	`
	rows, err := m.db.QueryContext(ctx, sqlStatement, username, username)
	if err != nil {
		m.log.ErrorContext(ctx, "fail getComments", "error", err)
		return nil, err
	}
	defer rows.Close()

	// NOTE: get all comments from db
	dbComments := make([]*DBComment, 0)
	for rows.Next() {
		c := new(DBComment)

		if err = rows.Scan(
			&c.ID,
			&c.Content,
			&c.Author,
			&c.Addressee,
			&c.Duration,
			&c.AvatarUrl,
			&c.IsMine,
			&c.ParentID,
			&c.Likes,
			&c.MyRate,
		); err != nil {
			m.log.ErrorContext(ctx, "fail getComments", "error", err)
			return nil, err
		}

		dbComments = append(dbComments, c)
	}

	// NOTE: map parent ids to value
	mIds := make(map[int]*Comment, len(pIds))
	for _, id := range pIds {
		if id != nil {
			mIds[*id] = new(Comment)
		}
	}

	for _, c := range dbComments {
		if _, ok := mIds[c.ID]; ok {
			mIds[c.ID] = &Comment{
				ID:        c.ID,
				Content:   c.Content,
				Author:    c.Author,
				AvatarUrl: c.AvatarUrl,
				Likes:     c.Likes,
				Duration:  c.Duration,
				IsMine:    c.IsMine,
				MyRate:    c.MyRate,
				Replies:   make([]*Reply, 0),
			}
		}
	}

	// NOTE: append replies to comments
	for _, c := range dbComments {
		if c != nil && c.ParentID != nil {
			if _, ok := mIds[*c.ParentID]; ok {
				r := &Reply{
					ID:        c.ID,
					Content:   c.Content,
					Author:    c.Author,
					AvatarUrl: c.AvatarUrl,
					Likes:     c.Likes,
					Duration:  c.Duration,
					IsMine:    c.IsMine,
					MyRate:    c.MyRate,
				}
				if c.Addressee != nil {
					r.Addressee = *c.Addressee
				}
				mIds[*c.ParentID].Replies = append(mIds[*c.ParentID].Replies, r)
			}
		}
	}

	comments := make([]*Comment, len(pIds), len(pIds))
	for i, id := range pIds {
		if id != nil {
			comments[i] = mIds[*id]
		}
	}

	m.log.InfoContext(ctx, "success getComments")
	return comments, nil
}

type CreateCommentInput struct {
	Author    *string
	Content   string
	ParentID  *int
	Addressee *string
}

func (m *Model) CreateComment(ctx context.Context, input *CreateCommentInput) error {
	m.log.InfoContext(ctx, "start CreateComment")

	sqlStatement := `
		INSERT INTO comment (author, content, parent_id, addressee)
		SELECT ?, ?, ?, ?
		WHERE ? IS NULL OR (
		    ? IS NOT NULL AND EXISTS (
				SELECT * FROM comment c WHERE c.id = ? AND c.parent_id IS NULL
			)
		);
	`

	res, err := m.db.ExecContext(
		ctx,
		sqlStatement,
		input.Author,
		input.Content,
		input.ParentID,
		input.Addressee,
		input.ParentID,
		input.ParentID,
		input.ParentID,
	)
	if err != nil {
		m.log.ErrorContext(ctx, "fail CreateComment", "error", err)
		return err
	}

	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if n == 0 {
		return fmt.Errorf("no record was inserted, please verify parent id")
	}

	m.log.InfoContext(ctx, "success CreateComment")
	return nil
}

type UpdateCommentInput struct {
	ID      *int
	Author  *string
	Content string
}

func (m *Model) UpdateComment(ctx context.Context, input *UpdateCommentInput) error {
	m.log.InfoContext(ctx, "start UpdateComment")

	sqlStatement := `
		UPDATE comment
		SET content = ?
		WHERE id = ? AND author = ?;
	`

	res, err := m.db.ExecContext(
		ctx,
		sqlStatement,
		input.Content,
		input.ID,
		input.Author,
	)
	if err != nil {
		m.log.ErrorContext(ctx, "fail UpdateComment", "error", err)
		return err
	}

	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if n == 0 {
		return fmt.Errorf("no record was update, please verify request")
	}

	m.log.InfoContext(ctx, "success UpdateComment")
	return nil
}

type DeleteCommentInput struct {
	ID       *int
	Username *string
}

func (m *Model) DeleteComment(ctx context.Context, input *DeleteCommentInput) error {
	m.log.InfoContext(ctx, "start DeleteComment")

	sqlStatement := `
		DELETE FROM comment WHERE id == ? AND author == ?;
	`

	res, err := m.db.ExecContext(
		ctx,
		sqlStatement,
		*input.ID,
		*input.Username,
	)
	if err != nil {
		m.log.ErrorContext(ctx, "fail DeleteComment", "error", err)
		return err
	}

	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if n == 0 {
		return fmt.Errorf("no record was deleted, please check you request")
	}

	m.log.InfoContext(ctx, "success DeleteComment")
	return nil
}
