-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
OFFSET $1
LIMIT $2;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByPhone :one
SELECT * FROM users
WHERE phone = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (
  email
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users SET
  email = $1,
  phone = $2,
  address = $3
WHERE
  id = $4;

-- name: DeleteAllUsers :exec
DELETE FROM users;
