package user

import (
	"time"

	"github.com/absorkun/darinol/model"
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
	Id       uint   `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

type UserUpdateDto struct {
	Id       uint   `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}
