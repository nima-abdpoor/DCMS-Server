// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: requestUrl.sql

package db

import (
	"context"
)

const createRequestUrl = `-- name: CreateRequestUrl :one
INSERT INTO requestUrl (unique_id,
                        request_url)
values ($1, $2) RETURNING id, unique_id, request_url
`

type CreateRequestUrlParams struct {
	UniqueID   int64  `json:"unique_id"`
	RequestUrl string `json:"request_url"`
}

func (q *Queries) CreateRequestUrl(ctx context.Context, arg CreateRequestUrlParams) (Requesturl, error) {
	row := q.db.QueryRowContext(ctx, createRequestUrl, arg.UniqueID, arg.RequestUrl)
	var i Requesturl
	err := row.Scan(&i.ID, &i.UniqueID, &i.RequestUrl)
	return i, err
}

const deleteRequestUrl = `-- name: DeleteRequestUrl :exec
DELETE
FROM requestUrl
WHERE id = $1
`

func (q *Queries) DeleteRequestUrl(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteRequestUrl, id)
	return err
}

const getRequestUrl = `-- name: GetRequestUrl :one
SELECT id, unique_id, request_url
FROM requestUrl
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRequestUrl(ctx context.Context, id int64) (Requesturl, error) {
	row := q.db.QueryRowContext(ctx, getRequestUrl, id)
	var i Requesturl
	err := row.Scan(&i.ID, &i.UniqueID, &i.RequestUrl)
	return i, err
}

const getRequestUrlByUniqueId = `-- name: GetRequestUrlByUniqueId :many
SELECT id, unique_id, request_url
FROM requestUrl
WHERE unique_id = $1
`

func (q *Queries) GetRequestUrlByUniqueId(ctx context.Context, uniqueID int64) ([]Requesturl, error) {
	rows, err := q.db.QueryContext(ctx, getRequestUrlByUniqueId, uniqueID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Requesturl
	for rows.Next() {
		var i Requesturl
		if err := rows.Scan(&i.ID, &i.UniqueID, &i.RequestUrl); err != nil {
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

const listRequestUrls = `-- name: ListRequestUrls :many
SELECT id, unique_id, request_url
FROM requestUrl
ORDER BY id LIMIT $1
OFFSET $2
`

type ListRequestUrlsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRequestUrls(ctx context.Context, arg ListRequestUrlsParams) ([]Requesturl, error) {
	rows, err := q.db.QueryContext(ctx, listRequestUrls, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Requesturl
	for rows.Next() {
		var i Requesturl
		if err := rows.Scan(&i.ID, &i.UniqueID, &i.RequestUrl); err != nil {
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

const updateRequestUrl = `-- name: UpdateRequestUrl :one
UPDATE requestUrl
set request_url = $2
WHERE unique_id = $1 RETURNING id, unique_id, request_url
`

type UpdateRequestUrlParams struct {
	UniqueID   int64  `json:"unique_id"`
	RequestUrl string `json:"request_url"`
}

func (q *Queries) UpdateRequestUrl(ctx context.Context, arg UpdateRequestUrlParams) (Requesturl, error) {
	row := q.db.QueryRowContext(ctx, updateRequestUrl, arg.UniqueID, arg.RequestUrl)
	var i Requesturl
	err := row.Scan(&i.ID, &i.UniqueID, &i.RequestUrl)
	return i, err
}
