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

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;