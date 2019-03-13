package service

import (
	"model"
	"sync"
)

var User = &userService{
	mutex: &sync.Mutex{},
}

type userService struct {
	mutex *sync.Mutex
}

func (u *userService) GetUserByID(uid uint) *model.User {
	var user model.User

	if err := db.Where("id = ?", uid).First(&user).Error; err != nil {
		return nil
	}
	return &user
}
