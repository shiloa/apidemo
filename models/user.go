package models

import "time"

type User struct {
	ID         string    `json:"id" gorm:"column:id;primary_key" xml:"id"`
	Email      string    `json:"email" gorm:"column:email" xml:"email"`
	Password   string    `json:"password" gorm:"column:password" xml:"password"`
	Name       string    `json:"name" gorm:"column:name" xml:"name"`
	IsVerified bool      `json:"is_verified" gorm:"column:is_verified" xml:"is_verified"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at" xml:"created_at"`
	Subscribed bool      `json:"subscribed" gorm:"column:subscribed" xml:"subscribed"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at" xml:"updated_at"`
	AuthToken  string    `json:"auth_token" gorm:"column:auth_token" xml:"auth_token"`
}

func (u User) TableName() string {
	return "Users"
}
