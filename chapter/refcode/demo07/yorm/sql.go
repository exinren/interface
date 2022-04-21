package yorm

import "strings"

//生成sql的插入语句。
func genInsertSQL(e interface{}) (string, []interface{}) {
	tbInfo, _ := ParseStruct(e)
	strInsert := "insert into " + tbInfo.Name
	strFields := ""
	strValues := ""
	var params []interface{}
	for _, v := range tbInfo.Fields {
		strFields += v.Name + ","
		strValues += "?,"
		if v.refValue.CanInterface() {
			params = append(params, v.refValue.Interface())
		}
	}
	strFields = strings.Trim(strFields, ",")
	strValues = strings.Trim(strValues, ",")
	strInsert += "(" + strFields + ") values(" + strValues + ")"
	return strInsert, params
}

// 更新sql语句
func genUpdateSQL(e interface{}) (string, []interface{}) {
	tbInfo, _ := ParseStruct(e)
	strUpdate := "update " + tbInfo.Name + " set "
	strFields := ""
	strWhere := ""
	var params []interface{}
	// 存储主键的值，最后推入数组，当搜索引擎
	var primaryValue interface{}
	for _, v := range tbInfo.Fields {
		if v.IsPrimaryKey {
			strWhere = " where " + v.Name + "=?"
			if v.refValue.CanInterface() {
				primaryValue = v.refValue.Interface()
			}
		} else {
			strFields += v.Name + "=?,"
			if v.refValue.CanInterface() {
				params = append(params, v.refValue.Interface())
			}
		}
	}
	params = append(params, primaryValue)
	strFields = strings.Trim(strFields, ",")
	strUpdate = strUpdate + strFields + strWhere
	return strUpdate, params
}

func genDeleteSQL (e interface{}) (string, []interface{})   {
	tbInfo, _ := ParseStruct(e)
	strUpdate := "update " + tbInfo.Name + " set "
	strFields := ""
	strWhere := ""
	var params []interface{}
	var primaryValue interface{}
	for _, v := range tbInfo.Fields {
		if v.IsPrimaryKey {
			strWhere = " where " + v.Name + "=?"
			if v.refValue.CanInterface() {
				primaryValue = v.refValue.Interface()
			}
		} else {
			if v.Name == "is_del" {
				strFields += v.Name + "=?"
				if v.refValue.CanInterface() {
					params = append(params, v.refValue.Interface())
				}
			}
		}
	}
	params = append(params, primaryValue)
	strFields = strings.Trim(strFields, ",")
	strUpdate = strUpdate + strFields + strWhere
	return strUpdate, params
}
