-- name: CreateUser :one
INSERT INTO users 
(
  username, 
  password, 
  first_name, 
  last_name
)
VALUES ($1, $2, $3, $4) 
RETURNING id, username, first_name, last_name, last_login;

-- name: UsernameExists :one
SELECT EXISTS (
    SELECT *
    FROM users 
    WHERE username = @username
);

-- name: GetUserByUsername :one
SELECT id, username, password, first_name, last_name
FROM users 
WHERE username = @username;
 
-- name: UpdateLastLogin :exec
UPDATE users
SET last_login = now()
WHERE id = @id;

-- name: UpdateLastLogout :exec
UPDATE users
SET last_logout = now()
WHERE id = @id;

-- name: UserExists :one
SELECT EXISTS (
    SELECT *
    FROM users 
    WHERE id = @id
);
