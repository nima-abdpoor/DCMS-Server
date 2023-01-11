-- name: CreateUrlFirst :one
INSERT INTO urlFirst (unique_id,
                      url_hash)
values ($1, $2) RETURNING *;

-- name: GetUrlFirst :one
SELECT *
FROM urlFirst
WHERE id = $1 LIMIT 1;

-- name: GetUrlFirstByUniqueId :many
SELECT *
FROM urlFirst
WHERE unique_id = $1;

-- name: ListUrlFirsts :many
SELECT *
FROM urlFirst
ORDER BY id LIMIT $1
OFFSET $2;

-- name: DeleteUrlFirst :exec
DELETE
FROM urlFirst
WHERE id = $1;

-- name: UpdateUrlFirst :one
UPDATE urlFirst
set url_hash = $2
WHERE unique_id = $1 RETURNING *;