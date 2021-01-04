package service 

import (
    "sixedu/model"
)

type UserService struct{}

var (
    user *model.User
)

func init() {
    user = model.NewUser()
}

func (u *UserService)GetList() []*model.User {
    return user.All()
}