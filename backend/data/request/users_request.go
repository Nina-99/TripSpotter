package request

type CreateUserRequest struct {
	Username string `validate:"required,min=2,max=50" json:"username"`
	Email    string `validate:"required,min=2,max=50" json:"email"`
	Password string `validate:"required,min=6,max=100" json:"password"`
	Role     string `validate:"required,min=1,max=10" json:"role"`
}

type UpdateUserRequest struct {
	Id       int    `validate:"required"`
	Username string `validate:"required,min=2,max=50" json:"username"`
	Email    string `validate:"required,min=2,max=50" json:"email"`
	Password string `validate:"required,min=6,max=100" json:"password"`
	Role     string `validate:"required,min=1,max=10" json:"role"`
}

type LoginUserRequest struct {
	Username string `validate:"required,min=2,max=50" json:"username"`
	Password string `validate:"required,min=6,max=100" json:"password"`
}
