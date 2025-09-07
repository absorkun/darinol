package todo

import (
	"time"

	"github.com/absorkun/darinol/model"
)

type TodoGetDto struct {
	Id          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Finished    bool       `json:"finished"`
	UserId      uint       `json:"user_id"`
	User        model.User `json:"user"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TodoCreateDto struct {
	Id          uint   `json:"id,omitempty" form:"id,omitempty" validate:"number"`
	Title       string `json:"title" form:"title" validate:"required,min=4,max=200"`
	Description string `json:"description,omitempty" form:"description,omitempty" validate:"max=500"`
	Finished    bool   `json:"finished,omitempty" form:"finished,omitempty" validate:"boolean"`
	UserId      uint   `json:"user_id" form:"user_id" validate:"required,number"`
}

type TodoUpdateDto struct {
	Id          uint   `json:"id,omitempty" form:"id,omitempty" validate:"number"`
	Title       string `json:"title,omitempty" form:"title,omitempty" validate:"required,min=4,max=200"`
	Description string `json:"description,omitempty" form:"description,omitempty" validate:"max=500"`
	Finished    bool   `json:"finished,omitempty" form:"finished,omitempty" validate:"boolean"`
	UserId      uint   `json:"user_id,omitempty" form:"user_id" validate:"required,number"`
}
