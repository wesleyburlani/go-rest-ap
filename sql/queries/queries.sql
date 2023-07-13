-- name: CreateUser :one
INSERT INTO users (id, email, password) VALUES(@id,@email,@password) RETURNING *;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id;

-- name: UpdateUser :one
UPDATE users SET email = coalesce(sqlc.narg(email), email), password = coalesce(sqlc.narg(password), password) WHERE id = @id
RETURNING *;
