package auth

type LoginDto struct {
	Email    string `json:"email" form:"email" validate:"required,email,max=100"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type RegisterDto struct {
	Id       uint   `json:"id,omitempty" form:"id,omitempty" validate:"number"`
	Email    string `json:"email" form:"email" validate:"required,email,max=100"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
	Role     string `json:"role,omitempty" form:"role,omitempty" validate:"role"`
}
