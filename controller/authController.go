package controller

import (
    "fmt"
    "sixedu/model"
    "sixedu/util" 
    "errors"
    "strconv"

)

type AuthController struct {

}

func (a *AuthController)Login() bool {
    fmt.Println("请输入用户名：")
    username := util.CRead()
    fmt.Println("请输入密码：")
    password := util.CRead()

    user := model.GetUser(username)

    if user == nil {
        fmt.Println("用户不存在")
        return false
    }

    if user.GetPassword() == password {
        fmt.Println("登入成功")
        view = "index_view"
        return true
    } else {
        fmt.Println("密码错误")
        view = "auth_view"
        return false
    }
    return false 
}

func (a *AuthController) Register() error {
    view = "index_view"
    fmt.Println("请输入你需注册的用户信息 username,password,age,sex")
    fmt.Println("输入你的用户名：")
    username := util.CRead()

    user := model.GetUser(username)

    if user != nil && user.GetUsername() == username {
        view = "auth_view"
        fmt.Println("用户名已经存在,请换一个用户名再注册")
        return errors.New("用户名已经存在")
    }
    

    fmt.Println("输入密码：")
    password := util.CRead()
    fmt.Println("确认密码：")
    confirmPassword := util.CRead()

    if password != confirmPassword {
        return errors.New("输入的密码不一致！")
    }

    fmt.Println("password:",password)

    fmt.Println("输入年龄：")
    age,_ := strconv.Atoi(util.CRead())
    fmt.Println("输入性别：")
    sex := util.CRead()

    result := authService.Register(username,password,age,sex)
    if result {
        fmt.Println("注册成功")
        return nil
    } else {
        return errors.New("注册失败")
    }


    return nil
}