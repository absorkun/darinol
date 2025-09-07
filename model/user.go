package model

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"primarykey;autoIncrement"`
	Email     string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"type:varchar(100);not null"`
	Role      string    `json:"role" gorm:"type:varchar(10);default:'user';not null"`
	Todos     []Todo    `json:"todos" gorm:"foreignKey:UserId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
