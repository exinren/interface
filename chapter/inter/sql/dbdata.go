package sql

import "strings"

type DbData struct {
	Table	string
	Fields	string
	Where	[]WhereItem
	OrderBy	[]OrderByItem
	PageIndex	int
	PageSize	int
	DbType
}

func NewDbData(table, field string) *DbData {
	return &DbData{
		Table: table,
		Fields: field,
	}
}

// 获取条件的字符串
func (db *DbData) GetWhere() string {
	s := make([]string, 0)
	for _, v := range db.Where {
		s = append(s, DbMapLeft[db.DbType])
		s = append(s, v.Field)
		s = append(s, DbMapRight[db.DbType])
		s = append(s, string(v.SqlOp))
		s = append(s,v.Value)
		s = append(s, string(v.OrAnd))
	}
	return strings.Join(s, "")
}

//获取排序的字符串
func (db *DbData) GetOrderBy() string {
	s := make([]string, 0)
	for _, v := range db.OrderBy {
		s = append(s, DbMapLeft[db.DbType])
		s = append(s, v.Field)
		s = append(s, DbMapRight[db.DbType])
		s = append(s," " + string(v.Order))
		s = append(s, ",")
	}
	// 移除最后的逗号
	s = s[0 : len(s) -1]
	return strings.Join(s, "")
}
