-- name: CreateUrlSecond :one
INSERT INTO urlSecond (id,
                       url_hash,
                       regex,
                       start_index,
                       finish_index)
values ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUrlSecond :one
SELECT *
FROM urlSecond
WHERE id = $1 LIMIT 1;

-- name: ListUrlSeconds :many
SELECT *
FROM urlSecond
ORDER BY id LIMIT $1
OFFSET $2;

-- name: DeleteUrlSecond :exec
DELETE
FROM urlSecond
WHERE id = $1;

-- name: UpdateUrlSecond :one
UPDATE urlSecond
set url_hash     = $2,
    regex        = $3,
    start_index  = $4,
    finish_index = $5
WHERE id = $1 RETURNING *;