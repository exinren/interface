package main

import (
	"fmt"
	"interface/sql"
)

func main()  {
	db := sql.NewDbData("Student", "ID,NAME,AGE")
	where := make([]sql.WhereItem, 0)
	where = append(where, sql.WhereItem{
		Field: "XH",
		Value: "9527",
		SqlOp: sql.EQ,
		OrAnd: sql.AND,
	})
	db.Where = where

	orderBy := make([]sql.OrderByItem, 0)
	orderBy = append(orderBy, sql.OrderByItem{
		Field: "XH",
		Order: sql.ASC,
	})
	db.OrderBy = orderBy
	db.DbType = sql.MYSQL
	sqlBuilder := sql.SqlBuilder{}
	sql1 := sqlBuilder.GenSelectSQL(db)
	fmt.Println(sql1)
	db.PageIndex = 0
	db.PageSize = 20
	sql10 := sqlBuilder.GenPageSQL(db)
	fmt.Println(sql10)

	db.DbType = sql.MSSQL
	sqlMs := sqlBuilder.GenSelectSQL(db)
	fmt.Println(sqlMs)
	sql10 = sqlBuilder.GenPageSQL(db)
	fmt.Println(sql10)
}