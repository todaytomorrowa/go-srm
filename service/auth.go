package service

import (
    "sixedu/model"
   // "fmt"
)

type AuthService struct {}


func (a *AuthService) Register(username,password string,age int, sex string) bool {
    user := model.NewUser()
    user.SetUsername(username)
    user.SetPassword(password)
    user.SetAge(age)
    user.SetSex(sex)
    user.Save()
   // fmt.Println("service.auth.Register.users:",user)

    return true
}