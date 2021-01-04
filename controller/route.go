package controller

import (
	"fmt"
	"reflect"
	"sixedu/util"
	"strconv"
	"strings"
    "errors"
)

var (
	authController *AuthController
	next           string
	view           string
)

func init() {
	authController = &AuthController{}
    fmt.Println(util.GetConfig())
}

func Run() {
    next = "index::Welcome"
    for {
        b := util.CReturn(util.Cfunc(dispatch))
        if b {
            break
        }
    }
    fmt.Println("结束")

}

func dispatch() (bool, error) {
	args := strings.Split(next, "::")
	controller, ok := controllers[args[0]]

	if ok != true {
		//fmt.Println("读取不到控制器：", args[0])
		return false, errors.New("无法根据指定标示查找到控制器:" + args[0])
	}

	//反射执行
	// 1.先传递需要执行的控制器对象
	// 2.根据方法名执行方法
	reflect.ValueOf(controller).MethodByName(args[1]).Call([]reflect.Value{})
    //reflect.ValueOf(controller).MethodByName(args[1]).Call([]reflect.Value{})

	//fmt.Println("view:",view)
	//获取下一步执行到操作
	opers, ok := views[view]
    //fmt.Println("views:",views)
    //fmt.Println("view:",view)
    //fmt.Println("opers:",opers)
	if ok != true {
        return false, errors.New("获取不到视图:"+args[1])
	}

	//数据的处理
	methods, desc := toMethodFormate(opers)

    //fmt.Println(methods)

	util.Coper(desc)

	for {
		input := util.CRead()

		if input == "x" {
			fmt.Println("退出")
            return true,nil
		}

		flag, err := strconv.Atoi(input)

       // fmt.Println("method:",methods)

		if err == nil && flag < len(methods) {
			next = methods[flag]
			break
		}
		fmt.Println("输入有误")
	}

    return false, nil
}
