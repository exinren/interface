package main

import (
	"fmt"
	"reflect"
)

/*
获取结构体字段标识
*/

type Student struct {
	id string `json:"id" iskey:"1"`
	Name string `json:"name" table:"t_student"`
	Age int `json:"age"`
}

func main()  {
	stu := Student{Name: "小二郎", Age: 10}
	t := reflect.TypeOf(&stu).Elem()
	m := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = t.Field(i).Tag.Get("json")
	}
	fmt.Println(m)
	if f, ok := t.FieldByName("Name"); ok {
		fmt.Printf("Name字段的table标识符： %v \n",f.Tag.Get("table"))
	}
	if f, ok := t.FieldByName("id"); ok {
		fmt.Printf("id字段的iskey标识符： %v \n",f.Tag.Get("iskey"))
	}
	//map[Age:age Name:name id:id]
	//Name字段的table标识符： t_student
	//id字段的iskey标识符： 1
}