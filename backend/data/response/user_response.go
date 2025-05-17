package response

type UserResponse struct {
	Id       int `json:"id"`
	Username int `json:"username"`
	Email    int `json:"email"`
	Password int `json:"pasword"`
	Role     int `json:"role"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
