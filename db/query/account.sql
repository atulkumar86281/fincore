-- name: CreateAccount :one
INSERT INTO account (
    owner,
    balance,
    currency
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * from account
where id = $1 limit 1;

-- name: GetAccountForUpdate :one
SELECT * from account
where id = $1 limit 1
FOR NO KEY UPDATE;

-- name: ListAccount :many
SELECT * from account
ORDER BY id
limit $1
OFFSET $2;


-- name: UpdateAccount :one
UPDATE account
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;