package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"projectgo/model"
	"projectgo/repository"
	"projectgo/utils"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetAllUsers() ([]*model.User, error)
	Login(username, password string) (string, string, error)
	GetUserByUsername(username string) (*model.User, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

func (u *UserServiceImpl) CreateUser(user *model.User) error {
	ulidID := utils.GenerateUlid()

	user.UserID = ulidID

	hashPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashPassword

	return u.userRepo.CreateUser(user)
}

func (u *UserServiceImpl) GetAllUsers() ([]*model.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *UserServiceImpl) Login(username, password string) (string, string, error) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	accesstoken, err := utils.GenerateJWTAccessToken(user.UserID)
	if err != nil {
		return "", "", err
	}

	refreshtoken, err := utils.GenerateJWTRefreshToken(user.UserID)
	if err != nil {
		return "", "", err
	}

	return accesstoken, refreshtoken, nil
}

func (u *UserServiceImpl) GetUserByUsername(username string) (*model.User, error) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
