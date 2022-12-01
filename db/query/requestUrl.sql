-- name: CreateRequestUrl :one
INSERT INTO requestUrl (unique_id,
                        request_url)
values ($1, $2) RETURNING *;

-- name: GetRequestUrl :one
SELECT *
FROM requestUrl
WHERE unique_id = $1 LIMIT 1;

-- name: ListRequestUrls :many
SELECT *
FROM requestUrl
ORDER BY id LIMIT $1
OFFSET $2;

-- name: DeleteRequestUrl :exec
DELETE
FROM requestUrl
WHERE id = $1;

-- name: UpdateRequestUrl :one
UPDATE requestUrl
set request_url = $2
WHERE unique_id = $1 RETURNING *;