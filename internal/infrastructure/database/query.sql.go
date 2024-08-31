// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :execresult
insert into user (
  id, screen_name, user_name, bio, is_private, created_at
) values (
  ?, ?, ?, ?, ?, ?
)
`

type CreateUserParams struct {
	ID         string
	ScreenName string
	UserName   string
	Bio        string
	IsPrivate  bool
	CreatedAt  time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.ScreenName,
		arg.UserName,
		arg.Bio,
		arg.IsPrivate,
		arg.CreatedAt,
	)
}

const findUserByScreenName = `-- name: FindUserByScreenName :one
select id, screen_name, user_name, bio, is_private, type, created_at from user use index (user_screen_name)
where screen_name = ? limit 1
`

func (q *Queries) FindUserByScreenName(ctx context.Context, screenName string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByScreenName, screenName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.ScreenName,
		&i.UserName,
		&i.Bio,
		&i.IsPrivate,
		&i.Type,
		&i.CreatedAt,
	)
	return i, err
}
