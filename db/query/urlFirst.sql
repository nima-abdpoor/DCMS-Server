-- name: CreateUrlFirst :one
INSERT INTO urlFirst (id,
                      url_hash)
values ($1, $2) RETURNING *;

-- name: GetUrlFirst :one
SELECT *
FROM urlFirst
WHERE id = $1 LIMIT 1;

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
WHERE id = $1 RETURNING *;