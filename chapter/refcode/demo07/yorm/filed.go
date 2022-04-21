package yorm

import "reflect"

type FieldInfo struct {
	Name string
	IsPrimaryKey bool
	refValue reflect.Value
}

