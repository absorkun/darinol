package user

import (
	"time"

	"github.com/absorkun/darinol/model"
)

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

type UserGetDto struct {
	Id        uint         `json:"id"`
	Email     string       `json:"email"`
	Password  string       `json:"-"`
	Role      string       `json:"role"`
	Todos     []model.Todo `json:"-" gorm:"foreignKey:UserId"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type UserCreateDto struct {
	Id       uint   `json:"id,omitempty" form:"id,omitempty" validate:"number"`
	Email    string `json:"email" form:"email" validate:"required,email,max=100"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
	Role     string `json:"role,omitempty"  form:"role,omitempty" validate:"role"`
}

type UserUpdateDto struct {
	Id       uint   `json:"id,omitempty"`
	Email    string `json:"email,omitempty" form:"email" validate:"required,email,max=100"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"required,min=6"`
	Role     string `json:"role,omitempty" form:"role,omitempty" validate:"role"`
}
