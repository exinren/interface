package yorm

import (
	"reflect"
)

// 解析结构体
func ParseStruct(s interface{}) (tInfo *TableInfo, err error)  {
	tInfo = &TableInfo{}
	rv := reflect.ValueOf(s)
	rt := reflect.TypeOf(s)
	method := rv.MethodByName("TableName")
	tableName := method.Call(nil)
	tInfo.Name = tableName[0].String()
	if rt.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}
	for i := 0; i < rt.NumField(); i++ {
		rtf := rt.Field(i)
		rvf := rv.Field(i)
		var f FieldInfo
		if rtf.Tag == "" {
			f = FieldInfo{Name: rtf.Name,IsPrimaryKey: false, refValue: rvf}
		} else {
			field := rtf.Tag.Get("table")
			IsPrimaryKey := false
			if  unique := rtf.Tag.Get("unique"); "true" == unique {
				IsPrimaryKey = true
			}
			f = FieldInfo{
				Name: field,
				IsPrimaryKey: IsPrimaryKey,
				refValue: rvf,
			}
		}
		tInfo.Fields = append(tInfo.Fields, f)
	}
	return
}
