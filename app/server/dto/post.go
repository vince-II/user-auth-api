package dto

type Post struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type CreatePostParams struct {
	Content string `json:"content" validate:"required"`
}
