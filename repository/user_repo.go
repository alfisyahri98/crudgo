package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"projectgo/model"
)

type UserRepo interface {
	CreateUser(user *model.User) error
	GetAllUsers() ([]*model.User, error)
	FindByUsername(username string) (*model.User, error)
}

type userRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepoImpl{db: db}
}

func (u *userRepoImpl) CreateUser(user *model.User) error {
	var existingUser model.User
	err := u.db.Where("username = ?", user.Username).First(&existingUser).Error
	if err == nil {
		return errors.New("username already exists")
	}

	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	
	return nil
}

func (u *userRepoImpl) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *userRepoImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("username %s not found", username)
	}
	return &user, nil
}
