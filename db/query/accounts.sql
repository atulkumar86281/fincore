-- name: CreateAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * from accounts
where id = $1 limit 1;

-- name: GetAccountForUpdate :one
SELECT * from accounts
where id = $1 limit 1
FOR NO KEY UPDATE;

-- name: ListAccount :many
SELECT * from accounts
ORDER BY id
limit $1
OFFSET $2;


-- name: UpdateAccount :one
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :one
DELETE FROM accounts
WHERE id = $1
RETURNING id;