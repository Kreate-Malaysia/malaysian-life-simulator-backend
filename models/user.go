package models

import "time"

type User struct {
	modelImpl
	UserName     string    `json:"name"`
	Email        string    `json:"email"`
	RegisterTime time.Time `json:"register_time"`
}

func NewUser(userName, email string) *User {
	return &User{
		UserName: userName,
		Email:    email,
	}
}

func (u *User) GetName() string {
	return u.UserName
}

func (u *User) GetEmail() string {
	return u.Email
}