package model

import "time"

type Todo struct {
	Id          uint      `json:"id" gorm:"primarykey;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(200);uniqueIndex;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Finished    bool      `json:"finished" gorm:"default:false"`
	UserId      uint      `json:"user_id" gorm:"not null"`
	User        User      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
