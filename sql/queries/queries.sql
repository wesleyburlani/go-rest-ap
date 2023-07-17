-- name: CreateUser :one
INSERT INTO users (email, password)
VALUES(@email,@password) RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = @id;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = @email;

-- name: UpdateUser :one
UPDATE users SET
  email = coalesce(sqlc.narg(email), email),
  password = coalesce(sqlc.narg(password), password)
WHERE id = @id
RETURNING *;

-- name: DeleteUserById :one
DELETE FROM users
WHERE id = @id
RETURNING *;

-- name: DeleteUserByEmail :one
DELETE FROM users
WHERE email = @email
RETURNING *;
