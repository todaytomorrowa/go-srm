package controller

import (
	"fmt"
)

type IndexController struct {
}

func (i *IndexController) Welcome() {
	view = "auth_view"
	fmt.Println("欢迎来到sixedu")
}

func (i *IndexController) Index() {
	view = "auth_view"
	fmt.Println("进入到首页")
}
