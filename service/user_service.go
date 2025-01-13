package service

import (
	"projectgo/model"
	"projectgo/repository"
	"projectgo/utils"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetAllUsers() ([]*model.User, error)
	Login(username, password string) error
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
func (u *UserServiceImpl) Login(username, password string) error {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return err
	}

	if err := utils.CheckPasswordHash(password, user.Password); err != nil {
		return err
	}

	return nil
}
