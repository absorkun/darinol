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
	Id          uint   `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Finished    bool   `json:"finished,omitempty"`
	UserId      uint   `json:"user_id"`
}

type TodoUpdateDto struct {
	Id          uint   `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Finished    bool   `json:"finished,omitempty"`
	UserId      uint   `json:"user_id,omitempty"`
}
