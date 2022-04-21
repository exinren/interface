package main

import (
	"fmt"
	"interface/chapter/refcode/demo07/entity"
	"interface/chapter/refcode/demo07/yorm"
	"time"
)

func main() {
	stu := entity.Student{
		Id:         "001",
		Name:       "wangxiaoer",
		Age:        10,
		Sex:        1,
		IsDel:      1,
		CreateDate: time.Now(),
		ModifyDate: time.Now(),
	}
	isOk := yorm.Save(&stu)
	//fmt.Println("新增成功",tm)
	if isOk {
		fmt.Println("新增成功")
	}
	isOk = yorm.Update(&stu)
	if isOk {
		fmt.Println("更新成功")
	}

	isOk = yorm.Delete(&stu)
	if isOk {
		fmt.Println("删除成功")
	}
}
