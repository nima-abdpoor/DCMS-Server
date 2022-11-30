-- name: CreateConfig :one
INSERT INTO config (id,
                    sync_type,
                    is_live)
values ($1, $2, $3) RETURNING *;

-- name: GetConfig :one
SELECT *
FROM config
WHERE id = $1 LIMIT 1;

-- name: ListConfigs :many
SELECT *
FROM config
ORDER BY id LIMIT $1
OFFSET $2;

-- name: DeleteConfig :exec
DELETE
FROM config
WHERE id = $1;

-- name: UpdateConfig :one
UPDATE config
set is_live = $2, sync_type = $3
WHERE id = $1 RETURNING *;