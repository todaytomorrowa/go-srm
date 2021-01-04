package main

import (
	"fmt"
	"reflect"
	"sixedu/controller"
)

func main() {
	controller.Run()

}

func ref(i interface{}) {
	rt := reflect.TypeOf(i)
	rv := reflect.ValueOf(i)

	fmt.Printf("rt type:%T\n", rt)
	fmt.Printf("rv type:%T\n", rv)

	for i := 0; i < rt.NumField(); i++ {
		fmt.Println("i 中元素字段:", rt.Field(i).Name)
		fmt.Println("i 中元素字段:", rt.Field(i).Name)

	}

	fmt.Println(rt.NumField())

	// rvi := rv.Interface()
	// fmt.Printf("rvi type:%T\n",rvi)
}
