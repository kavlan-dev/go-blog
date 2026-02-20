package service

import (
	"go-blog/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type userStorage interface {
	FindUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
}

type userService struct {
	db userStorage
}

func NewUserService(db userStorage) userService {
	return userService{db: db}
}

func (s userService) AuthenticateUser(authUser *model.User) (*model.User, error) {
	user, err := s.db.FindUserByUsername(authUser.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authUser.Password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s userService) RegisterUser(newUser *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPassword)

	if err := s.db.CreateUser(newUser); err != nil {
		return err
	}

	return nil
}
