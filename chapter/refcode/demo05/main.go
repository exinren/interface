package main

import (
	"fmt"
	"reflect"
)

/**
动态调用方法
 */

type PC struct {
	Name string
}

func (p *PC) GetName() string {
	return p.Name
}

func (p PC) Sum(a, b int) int {
	return a + b
}

func main()  {
	pc := PC{
		Name: "外星人",
	}
	vt := reflect.ValueOf(&pc)
	vm := vt.MethodByName("GetName")
	res := vm.Call(nil)
	fmt.Println("GetName返回的：",res[0].String())
	vm = vt.MethodByName("Sum")
	res = vm.Call([]reflect.Value{reflect.ValueOf(3), reflect.ValueOf(5)})
	fmt.Println("Sum(3, 5) = ", res[0].Int())
	// 非指针类型只能调用值接受方法
	vt = reflect.ValueOf(pc)
	vm = vt.MethodByName("Sum")
	res = vm.Call([]reflect.Value{reflect.ValueOf(3), reflect.ValueOf(5)})
	fmt.Println("Sum(3, 5) = ", res[0].Int())
	// vm = vt.MethodByName("GetName")
	// res = vm.Call(nil)  报错
}
