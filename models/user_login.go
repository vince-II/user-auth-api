package models

type UserLogin struct {
	UserID      int    `json:"user_id"`
	Token       string `json:"token"`
	LastLoginAt string `json:"last_login_at"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
}
