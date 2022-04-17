package main

import (
	"fmt"
	"reflect"
)

/*
反射动态修改值
*/

func main()  {
	msg := "没事干"
	value := reflect.ValueOf(&msg)
	// 如果v的类型不是interface或ptr(指针)。则抛出panic
	fmt.Println(value.Elem())	// 没事干
	//reflect.ValueOf(msg).Elem() // panic
	if value.Elem().CanSet() {
		value.Elem().Set(reflect.ValueOf("阿拉丁"))
	}
	fmt.Println(msg)	// 阿拉丁
	// 值类型的无法修改
	v := reflect.ValueOf(msg)
	if v.CanSet() {//false
		value.SetString("小公鸡")
		fmt.Println(msg)
	}
}
