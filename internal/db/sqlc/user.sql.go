// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password
) VALUES (
  $1, $2
) RETURNING username, hashed_password, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.HashedPassword)
	var i User
	err := row.Scan(&i.Username, &i.HashedPassword, &i.CreatedAt)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_password, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(&i.Username, &i.HashedPassword, &i.CreatedAt)
	return i, err
}
