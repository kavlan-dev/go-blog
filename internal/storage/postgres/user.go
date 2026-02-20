package postgres

import "go-blog/internal/model"

func (s *storage) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *storage) CreateUser(newUser *model.User) error {
	return s.db.Create(&newUser).Error
}
