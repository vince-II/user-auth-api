package dto

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	LastLogin string `json:"last_login,omitempty"` // optional field
	CreatedAt string `json:"created_at,omitempty"` // optional field
	UpdatedAt string `json:"updated_at,omitempty"` // optional field
}

type RegisterUser struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required"` // check if username is unique
	Password  string `json:"password" validate:"required, min=8,max=32"`
}

type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
