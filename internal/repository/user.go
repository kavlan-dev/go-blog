package repository

import (
	"go-blog/internal/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (s userRepository) FindUserByUsername(username string) (model.User, error) {
	var user model.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s userRepository) CreateUser(newUser model.User) error {
	return s.db.Create(&newUser).Error
}
