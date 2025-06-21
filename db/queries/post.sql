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


-- name: GetAllPostFromUser :one 
SELECT *
FROM post
WHERE user_id = @user_id;
