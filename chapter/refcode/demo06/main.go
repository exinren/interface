package main

import (
	"errors"
	"fmt"
	"reflect"
)

/**
结构体转map类型
*/

type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Profile
}

type Profile struct {
	Hobby string `json:"hobby" structs:"hobby"`
}

func main() {
	stu := Student{
		Name:   "wangxiaoer",
		Age:    10,
		Gender: "男",
		Profile:Profile{
			Hobby: "下棋",
		},
	}
	m1, _ := StructTranSingleMap(&stu, "json")
	fmt.Println(m1)
	//s1 := Student{}
	//s := map[string]interface{} {
	//	"Name": "王小二",
	//	"Age": 13,
	//	"Gender": "女",
	//}
	//MapToStruct(s, &s1)
	//fmt.Println(s1)
	//fmt.Println(s)
}

/*
in：结构体
tagName： tag名称，如果传入空字符串，就会取属性作为map的键。
返回map类型和一个错误
*/
func StructToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		err := errors.New(fmt.Sprintf("StructToMap函数只接受结构体或者结构体指针; 当前参数的类型: %T", v))
		return nil, err
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tagName == "" {
			out[field.Name] = v.Field(i).Interface()
		} else {
			if tagValue := field.Tag.Get(tagName); tagValue != "" {
				out[tagValue] = v.Field(i).Interface()
			}
		}
	}
	return out, nil
}

/*
根据结构体中的字段或者字段tag中的json的进行转换。
m：map类型
u：结构体
*/
func MapToStruct(m map[string]interface{}, u interface{}) error {
	v := reflect.ValueOf(u)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		if v.Kind() != reflect.Struct {
			return errors.New(fmt.Sprintf("传入必须是结构体"))
		}
		findFromMap := func(key, tagName string) interface{} {
			for k, va := range m {
				if k == key || k == tagName {
					return va
				}
			}
			return nil
		}

		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			val := findFromMap(field.Name, field.Tag.Get("json"))
			if nil != val && reflect.ValueOf(val).Kind() == v.Field(i).Kind() {
				v.Field(i).Set(reflect.ValueOf(val))
			}
		}
	} else {
		return errors.New(fmt.Sprintf("传入必须是结构体的指针类型"))
	}
	return nil
}

/*
将结构体（可嵌套）转换单层map
in：结构体
tag： 标签的字符串。
 */
func StructTranSingleMap(in interface{}, tag string) (map[string]interface{}, error) {
	v := reflect.ValueOf(in)
	// 当前函数只接受结构体
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New(fmt.Sprintf("必须传入结构体"))
	}
	out := make(map[string]interface{})
	queue := make([]interface{}, 0, 1)
	queue = append(queue, in)
	for len(queue) > 0 {
		v1 := reflect.ValueOf(queue[0])
		if v1.Kind() == reflect.Ptr {
			v1 = v1.Elem()
		}
		queue = queue[1:]
		t := v1.Type()
		for i := 0; i < t.NumField(); i++ {
			v2 := v1.Field(i)
			// 指针类型的指定
			if v2.Kind() == reflect.Ptr {
				v2 = v2.Elem()
				// 是否为结构体
				if v2.Kind() == reflect.Struct {
					queue = append(queue, v2.Interface())
				} else {
					// 获取标签的值
					if tagValue := t.Field(i).Tag.Get(tag); tagValue != "" {
						out[tagValue] = v2.Interface()
					}
				}
				break
			}
			// 结构体类型，加入切片中，直接退出，在循环
			if v2.Kind() == reflect.Struct { //内嵌结构体
				queue = append(queue, v2.Interface())
				break
			}
			tl := t.Field(i)
			//普通类型字段
			if tagValue := tl.Tag.Get(tag); tagValue != "" {
				// 存入Map
				out[tagValue] = v2.Interface()
			}
		}
	}
	return out, nil
}
