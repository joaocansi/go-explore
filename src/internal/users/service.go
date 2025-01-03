package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) Create(data CreateUserInput) (*User, error) {
	_, err := u.repository.FindByEmail(data.Email)
	if err == nil {
		return nil, fmt.Errorf("user already exists")
	}

	hashedPassword, err := HashPassword(data.Password)
	if err != nil {
		return nil, fmt.Errorf("not able to hash password")
	}

	createdUser, err := u.repository.Create(&UserModel{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPassword,
	})

	if err != nil {
		return nil, fmt.Errorf("not able to create user")
	}

	return createdUser.ToDomain(), nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
