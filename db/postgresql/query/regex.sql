-- name: CreateRegex :one
INSERT INTO regex (url_id,
                   regex,
                   start_index,
                   finish_index)
values ($1, $2, $3, $4) RETURNING *;

-- name: GetRegex :one
SELECT *
FROM regex
WHERE id = $1 LIMIT 1;

-- name: GetRegexByUrlId :many
SELECT *
FROM regex
WHERE url_id = $1;

-- name: ListRegexes :many
SELECT *
FROM regex
ORDER BY id LIMIT $1
OFFSET $2;

-- name: DeleteRegex :exec
DELETE
FROM regex
WHERE id = $1;

-- name: UpdateRegex :one
UPDATE regex
set regex        = $2,
    start_index  = $3,
    finish_index = $4
WHERE url_id = $1 RETURNING *;