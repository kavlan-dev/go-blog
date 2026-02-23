package service

import (
	"go-blog/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type userRepository interface {
	FindUserByUsername(username string) (model.User, error)
	CreateUser(user model.User) error
}

type userService struct {
	repo userRepository
}

func NewUserService(repo userRepository) *userService {
	return &userService{repo: repo}
}

func (s userService) AuthenticateUser(authUser model.UserRequest) (model.User, error) {
	user, err := s.repo.FindUserByUsername(authUser.Username)
	if err != nil {
		return model.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authUser.Password)); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s userService) RegisterUser(newUser model.UserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hashedPassword)

	user := newUser.ToUser()

	if err := s.repo.CreateUser(user); err != nil {
		return err
	}

	return nil
}
