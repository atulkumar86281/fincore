-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;


-- name: GetUser :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET full_name = $2,
    email = $3
WHERE username = $1
RETURNING *;


-- name: UpdateUserPassword :one
UPDATE users
SET hashed_password = $2,
    password_last_updated = now()
WHERE username = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE username = $1
RETURNING username;