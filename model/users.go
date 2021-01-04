package model

import (
    "strconv"
	"fmt"
)

type User struct {
	username string
	password string
	age      int
	sex      string
}

var userDatas map[string]Model

func NewUser() *User {
    return &User{}
}

func GetUser(username string) *User{
	if userDatas[username] == nil {
		return nil
	}
	//fmt.Println(userDatas)
	return userDatas[username].(*User)
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) SetPassword(password string) {
	u.password = password
	fmt.Println("u:35::",u)
}

func (u *User) SetAge(age int) {
	u.age = age
}

func (u *User) SetSex(sex string) {
	u.sex = sex
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) GetSex() string {
	return u.sex
}

// 
func (u *User)ToString()string {
	if u != nil {
     return u.username + "," + u.password + "," + strconv.Itoa(u.age) + "," + u.sex
	}
	return ""
}

func (u *User) Save() bool {
	userDatas[u.username] = u
	fmt.Println("u:68::",u)
	return fwrite("users", userDatas)
}

func (u *User) All() []*User {
	var users []*User = make([]*User,0)
	for _, user := range userDatas {
		users = append(users,user.(*User))
	}
	return users
}





