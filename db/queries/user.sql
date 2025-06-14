-- name: CreateUser :one
INSERT INTO users (username, password, first_name, last_name)
VALUES ($1, $2, $3, $4) 
RETURNING id, username, first_name, last_name, last_login;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;
