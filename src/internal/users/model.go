package users

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (u *UserModel) ToDomain() *User {
	return &User{
		Name:  u.Name,
		Email: u.Email,
	}
}
