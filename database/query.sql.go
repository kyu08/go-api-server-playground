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
  id, screen_name, name,
  bio, is_private, created_at
) values (
  ?, ?, ?, ?, ?, ?
)
`

type CreateUserParams struct {
	ID         string
	ScreenName string
	Name       string
	Bio        string
	IsPrivate  bool
	CreatedAt  time.Time
}

// -- name: GetAuthor :one
// select * from authors
// where id = ? limit 1;
//
// -- name: ListAuthors :many
// select * from authors
// order by name2;
//
// -- name: CreateAuthor :execresult
// insert into authors (
//
//	name2, bio
//
// ) values (
//
//	?, ?
//
// );
//
// -- name: DeleteAuthor :exec
// delete from authors
// where id = ?;
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.ScreenName,
		arg.Name,
		arg.Bio,
		arg.IsPrivate,
		arg.CreatedAt,
	)
}

const getUserByScreenName = `-- name: GetUserByScreenName :one
select id, screen_name, name, bio, is_private, created_at from user
where screen_name = ? limit 1
`

func (q *Queries) GetUserByScreenName(ctx context.Context, screenName string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByScreenName, screenName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.ScreenName,
		&i.Name,
		&i.Bio,
		&i.IsPrivate,
		&i.CreatedAt,
	)
	return i, err
}
