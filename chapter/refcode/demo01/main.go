package main

import (
	"fmt"
	"reflect"
)

/*
通过反射获取结构体字段与方法
*/
type Student struct {
	id   string
	Name string
	Age  int
}
// 值类型
func (stu Student) GetName()  {
	fmt.Println(stu.Name)
}
// 指针类型
func (stu *Student) GetAge() {
	fmt.Println(stu.Age)
}

func (stu Student) getId() {
	fmt.Println(stu.id)
}

// 打印结构体字段和方法
func printFieldMethod(o interface{})  {
	value := reflect.ValueOf(o)
	typeof := value.Type()
	fmt.Println(typeof)
	if typeof.Kind() == reflect.Struct {
		// 获取属性
		for i := 0; i < typeof.NumField(); i++ {
			field := typeof.Field(i)
			if value.Field(i).CanInterface() {
				// 私有属性报错
				v := value.Field(i).Interface()
				fmt.Printf("%s: %v = %v\n",field.Name, field.Type, v)
			} else {
				fmt.Printf("%s: %v = %v\n", field.Name, field.Type, "私有属性")
			}
		}
	}
	// 获取方法(不能获取私有的)
	for i := 0; i < typeof.NumMethod(); i++ {
		method := typeof.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
		//	GetName:func(main.Student)
	}
}

func main()  {
	stu := Student{Name: "xiaoer", Age: 18}
	// 只能获取值接受者的方法
	printFieldMethod(stu)
	fmt.Println("---------------------------------------")
	a := "xiaoming"
	printFieldMethod(a)
	fmt.Println("---------------------------------------")
	// 可以获取指针接收者方法
	printFieldMethod(&stu)
}