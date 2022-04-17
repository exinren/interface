package main

/**
判断是否实现接口
 */

import (
	"fmt"
	"reflect"
)

type IDuck interface {
	SingGua()
}

type Goose struct {}

func (this *Goose) SingGua()  {
	fmt.Println("Goose Sing Gua")
}

func main()  {
	intf := new(IDuck)
	intfType := reflect.TypeOf(intf).Elem()
	goose := Goose{}
	srcType := reflect.TypeOf(&goose)
	if srcType.Implements(intfType) {
		fmt.Printf("%v实现了%v接口\n", srcType, intfType)
	} else {
		fmt.Printf("%v没有实现%v接口\n", srcType, intfType)
	}
	fmt.Println("============================")
	srcType = reflect.TypeOf(goose)
	if srcType.Implements(intfType) {
		fmt.Printf("%v实现了%v接口\n", srcType, intfType)
	} else {
		// 没实现
		fmt.Printf("%v没有实现%v接口\n", srcType, intfType)
	}
}
