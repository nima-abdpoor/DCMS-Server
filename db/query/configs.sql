-- name: CreateConfig :one
INSERT INTO configs (id,
                     is_live,
                     sync_type,
                     valid_request_url,
                     url_id_first,
                     url_id_second,
                     regex,
                     start_index,
                     finish_index)
values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetConfig :one
SELECT * FROM configs
WHERE id = $1 LIMIT 1;

-- name: ListConfigs :many
SELECT * FROM configs
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: DeleteConfig :exec
DELETE FROM configs
WHERE id = $1;

-- name: UpdateConfig :one
UPDATE configs
set is_live = $2 AND
set sync_type = $3 AND
set valid_request_url = $4 AND
set url_id_first = $5 AND
set url_id_second = $6 AND
set regex = $7 AND
set start_index = $7 AND
set
WHERE id = $1
    RETURNING *;