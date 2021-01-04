package controller

import (
   // "sixedu/service"
    "fmt"
)

type UserController struct {

}

func (u *UserController) List() {
    //展示用户信息
    view = "index_view"
    users := userService.GetList()
    fmt.Println("| username | passsword | age | sex |")
    //fmt.Println(users)
    for _,v := range users {
        fmt.Printf("| %s | %s | %d | %s | \n",v.GetUsername(),v.GetPassword(),v.GetAge(),v.GetSex())
    }

}
