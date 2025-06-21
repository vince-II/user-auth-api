-- name: CreatePost :one
INSERT INTO post
(
  user_id,
  content,
  created_at
)
VALUES(
  $1, $2, $3
)
RETURNING *;

-- name: UpdatePost :one
UPDATE post
SET content = @content
WHERE id = @id
RETURNING *;

-- name: DeletePost :exec
DELETE 
FROM post
WHERE id = @id;

-- name: GetPost :one
SELECT *
FROM post
where id = @id;
