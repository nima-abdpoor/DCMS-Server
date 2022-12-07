-- name: CreateUrlSecond :one
INSERT INTO urlSecond (unique_id,
                       url_hash)
values ($1, $2) RETURNING *;

-- name: GetUrlSecond :one
SELECT *
FROM urlSecond
WHERE id = $1 LIMIT 1;

-- name: GetUrlSecondByUniqueId :many
SELECT *
FROM urlSecond
WHERE unique_id = $1;

-- name: ListUrlSeconds :many
SELECT *
FROM urlSecond
ORDER BY unique_id LIMIT $1
OFFSET $2;

-- name: DeleteUrlSecond :exec
DELETE
FROM urlSecond
WHERE id = $1;

-- name: UpdateUrlSecond :one
UPDATE urlSecond
set url_hash = $2
WHERE unique_id = $1 RETURNING *;