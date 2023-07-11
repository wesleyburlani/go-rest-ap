-- name: CreateUser :exec
INSERT INTO users (id, email, password) VALUES($1,$2,$3);

-- name: ListUsers :many
SELECT * FROM users ORDER BY id;
