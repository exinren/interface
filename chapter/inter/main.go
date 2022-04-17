package main

import (
	"fmt"
	sql2 "interface/chapter/inter/sql"
)

func main()  {
	db := sql2.NewDbData("Student", "ID,NAME,AGE")
	where := make([]sql2.WhereItem, 0)
	where = append(where, sql2.WhereItem{
		Field: "XH",
		Value: "9527",
		SqlOp: sql2.EQ,
		OrAnd: sql2.AND,
	})
	db.Where = where

	orderBy := make([]sql2.OrderByItem, 0)
	orderBy = append(orderBy, sql2.OrderByItem{
		Field: "XH",
		Order: sql2.ASC,
	})
	db.OrderBy = orderBy
	db.DbType = sql2.MYSQL
	sqlBuilder := sql2.SqlBuilder{}
	sql1 := sqlBuilder.GenSelectSQL(db)
	fmt.Println(sql1)
	db.PageIndex = 0
	db.PageSize = 20
	sql10 := sqlBuilder.GenPageSQL(db)
	fmt.Println(sql10)

	db.DbType = sql2.MSSQL
	sqlMs := sqlBuilder.GenSelectSQL(db)
	fmt.Println(sqlMs)
	sql10 = sqlBuilder.GenPageSQL(db)
	fmt.Println(sql10)

}