package pkg

import (
	"api/src/internal/users"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&users.UserModel{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
