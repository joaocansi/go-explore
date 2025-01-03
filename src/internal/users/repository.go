package users

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByEmail(email string) (*UserModel, error) {
	var user UserModel
	result := r.db.First(&user).Where(&User{Email: email})

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) Create(user *UserModel) (*UserModel, error) {
	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
