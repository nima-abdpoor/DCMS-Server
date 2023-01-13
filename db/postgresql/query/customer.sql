-- name: CreateCustomer :one
INSERT INTO customer (username,
                      password,
                      info,
                      email,
                      package_name,
                      sdk_uuid,
                      secret_key)
values ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetCustomer :one
SELECT *
FROM customer
WHERE id = $1 LIMIT 1;

-- name: GetCustomerByUsername :one
SELECT *
FROM customer
WHERE username = $1 LIMIT 1;

-- name: DeleteCustomer :exec
DELETE
FROM customer
WHERE id = $1;

-- name: UpdateCustomer :one
UPDATE customer
set username     = $2,
    password     = $3,
    info         = $4,
    email        = $5,
    package_name = $6,
    sdk_uuid     = $7,
    secret_key   = $8
WHERE id = $1 RETURNING *;