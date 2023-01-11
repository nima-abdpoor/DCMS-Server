-- name: CreateConfig :one
INSERT INTO config (id,
                    is_live,
                    save_request,
                    save_response,
                    save_error,
                    save_success,
                    network_type,
                    repeat_interval,
                    repeat_interval_time_unit,
                    requires_battery_not_low,
                    requires_storage_not_low,
                    requires_charging,
                    requires_device_idl)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING *;

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
set is_live                   = $2,
    save_request              = $3,
    save_response             = $4,
    save_error                = $5,
    save_success              = $6,
    network_type              = $7,
    repeat_interval           = $8,
    repeat_interval_time_unit = $9,
    requires_battery_not_low  = $10,
    requires_storage_not_low  = $11,
    requires_charging         = $12,
    requires_device_idl       = $13
WHERE id = $1 RETURNING *;