package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;auto_increment:true"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserProps struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
