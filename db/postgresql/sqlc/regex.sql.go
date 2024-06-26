// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: regex.sql

package db

import (
	"context"
)

const createRegex = `-- name: CreateRegex :one
INSERT INTO regex (url_id,
                   regex,
                   start_index,
                   finish_index)
values ($1, $2, $3, $4) RETURNING id, url_id, regex, start_index, finish_index
`

type CreateRegexParams struct {
	UrlID       int64  `json:"url_id"`
	Regex       string `json:"regex"`
	StartIndex  int32  `json:"start_index"`
	FinishIndex int32  `json:"finish_index"`
}

func (q *Queries) CreateRegex(ctx context.Context, arg CreateRegexParams) (Regex, error) {
	row := q.db.QueryRowContext(ctx, createRegex,
		arg.UrlID,
		arg.Regex,
		arg.StartIndex,
		arg.FinishIndex,
	)
	var i Regex
	err := row.Scan(
		&i.ID,
		&i.UrlID,
		&i.Regex,
		&i.StartIndex,
		&i.FinishIndex,
	)
	return i, err
}

const deleteRegex = `-- name: DeleteRegex :exec
DELETE
FROM regex
WHERE id = $1
`

func (q *Queries) DeleteRegex(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteRegex, id)
	return err
}

const getRegex = `-- name: GetRegex :one
SELECT id, url_id, regex, start_index, finish_index
FROM regex
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRegex(ctx context.Context, id int64) (Regex, error) {
	row := q.db.QueryRowContext(ctx, getRegex, id)
	var i Regex
	err := row.Scan(
		&i.ID,
		&i.UrlID,
		&i.Regex,
		&i.StartIndex,
		&i.FinishIndex,
	)
	return i, err
}

const getRegexByUrlId = `-- name: GetRegexByUrlId :many
SELECT id, url_id, regex, start_index, finish_index
FROM regex
WHERE url_id = $1
`

func (q *Queries) GetRegexByUrlId(ctx context.Context, urlID int64) ([]Regex, error) {
	rows, err := q.db.QueryContext(ctx, getRegexByUrlId, urlID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Regex
	for rows.Next() {
		var i Regex
		if err := rows.Scan(
			&i.ID,
			&i.UrlID,
			&i.Regex,
			&i.StartIndex,
			&i.FinishIndex,
		); err != nil {
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

const listRegexes = `-- name: ListRegexes :many
SELECT id, url_id, regex, start_index, finish_index
FROM regex
ORDER BY id LIMIT $1
OFFSET $2
`

type ListRegexesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRegexes(ctx context.Context, arg ListRegexesParams) ([]Regex, error) {
	rows, err := q.db.QueryContext(ctx, listRegexes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Regex
	for rows.Next() {
		var i Regex
		if err := rows.Scan(
			&i.ID,
			&i.UrlID,
			&i.Regex,
			&i.StartIndex,
			&i.FinishIndex,
		); err != nil {
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

const updateRegex = `-- name: UpdateRegex :one
UPDATE regex
set regex        = $2,
    start_index  = $3,
    finish_index = $4
WHERE url_id = $1 RETURNING id, url_id, regex, start_index, finish_index
`

type UpdateRegexParams struct {
	UrlID       int64  `json:"url_id"`
	Regex       string `json:"regex"`
	StartIndex  int32  `json:"start_index"`
	FinishIndex int32  `json:"finish_index"`
}

func (q *Queries) UpdateRegex(ctx context.Context, arg UpdateRegexParams) (Regex, error) {
	row := q.db.QueryRowContext(ctx, updateRegex,
		arg.UrlID,
		arg.Regex,
		arg.StartIndex,
		arg.FinishIndex,
	)
	var i Regex
	err := row.Scan(
		&i.ID,
		&i.UrlID,
		&i.Regex,
		&i.StartIndex,
		&i.FinishIndex,
	)
	return i, err
}
