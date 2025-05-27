-- name: RegisterUser :one
INSERT INTO users (first_name, last_name)
VALUES ($1, $2) 
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;
