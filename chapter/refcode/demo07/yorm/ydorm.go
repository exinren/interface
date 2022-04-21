package yorm

import "fmt"

func Save(e interface{}) (isOk bool) {
	defer func() {
		if err := recover(); nil != err {
			isOk = false
			fmt.Println("捕获异常", err)
		}
	}()
	strSQL, value := genInsertSQL(e)
	fmt.Println("sql语句：", strSQL)
	fmt.Println("需要的值：", value)
	isOk = true
	return
}

func Update(e interface{}) (isOk bool)  {
	defer func() {
		if err := recover(); nil != err {
			isOk = false
			fmt.Println("捕获异常", err)
		}
	}()
	strSQL, value := genUpdateSQL(e)
	fmt.Println("sql语句：", strSQL)
	fmt.Println("需要的值：", value)
	isOk = true
	return
}

func Delete(e interface{}) (isOk bool) {
	defer func() {
		if err := recover(); nil != err {
			isOk = false
			fmt.Println("捕获异常", err)
		}
	}()
	strSQL, value := genDeleteSQL(e)
	fmt.Println("sql语句：", strSQL)
	fmt.Println("需要的值：", value)
	isOk = true
	return isOk

}
