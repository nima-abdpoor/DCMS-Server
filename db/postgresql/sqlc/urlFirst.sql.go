// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: urlFirst.sql

package db

import (
	"context"
)

const createUrlFirst = `-- name: CreateUrlFirst :one
INSERT INTO urlFirst (unique_id,
                      url_hash)
values ($1, $2) RETURNING id, unique_id, url_hash
`

type CreateUrlFirstParams struct {
	UniqueID int64  `json:"unique_id"`
	UrlHash  string `json:"url_hash"`
}

func (q *Queries) CreateUrlFirst(ctx context.Context, arg CreateUrlFirstParams) (Urlfirst, error) {
	row := q.db.QueryRowContext(ctx, createUrlFirst, arg.UniqueID, arg.UrlHash)
	var i Urlfirst
	err := row.Scan(&i.ID, &i.UniqueID, &i.UrlHash)
	return i, err
}

const deleteUrlFirst = `-- name: DeleteUrlFirst :exec
DELETE
FROM urlFirst
WHERE id = $1
`

func (q *Queries) DeleteUrlFirst(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUrlFirst, id)
	return err
}

const getUrlFirst = `-- name: GetUrlFirst :one
SELECT id, unique_id, url_hash
FROM urlFirst
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUrlFirst(ctx context.Context, id int64) (Urlfirst, error) {
	row := q.db.QueryRowContext(ctx, getUrlFirst, id)
	var i Urlfirst
	err := row.Scan(&i.ID, &i.UniqueID, &i.UrlHash)
	return i, err
}

const getUrlFirstByUniqueId = `-- name: GetUrlFirstByUniqueId :many
SELECT id, unique_id, url_hash
FROM urlFirst
WHERE unique_id = $1
`

func (q *Queries) GetUrlFirstByUniqueId(ctx context.Context, uniqueID int64) ([]Urlfirst, error) {
	rows, err := q.db.QueryContext(ctx, getUrlFirstByUniqueId, uniqueID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Urlfirst
	for rows.Next() {
		var i Urlfirst
		if err := rows.Scan(&i.ID, &i.UniqueID, &i.UrlHash); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUrlFirsts = `-- name: ListUrlFirsts :many
SELECT id, unique_id, url_hash
FROM urlFirst
ORDER BY id LIMIT $1
OFFSET $2
`

type ListUrlFirstsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUrlFirsts(ctx context.Context, arg ListUrlFirstsParams) ([]Urlfirst, error) {
	rows, err := q.db.QueryContext(ctx, listUrlFirsts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Urlfirst
	for rows.Next() {
		var i Urlfirst
		if err := rows.Scan(&i.ID, &i.UniqueID, &i.UrlHash); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUrlFirst = `-- name: UpdateUrlFirst :one
UPDATE urlFirst
set url_hash = $2
WHERE unique_id = $1 RETURNING id, unique_id, url_hash
`

type UpdateUrlFirstParams struct {
	UniqueID int64  `json:"unique_id"`
	UrlHash  string `json:"url_hash"`
}

func (q *Queries) UpdateUrlFirst(ctx context.Context, arg UpdateUrlFirstParams) (Urlfirst, error) {
	row := q.db.QueryRowContext(ctx, updateUrlFirst, arg.UniqueID, arg.UrlHash)
	var i Urlfirst
	err := row.Scan(&i.ID, &i.UniqueID, &i.UrlHash)
	return i, err
}
