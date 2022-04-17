package sql

import (
	"strconv"
	"strings"
)

type MsSQL struct {}

func Select(db *DbData) string {
	if db == nil {
		return ""
	}
	s := make([]string, 0)
	s = append(s, "SELECT ")
	fields := strings.Split(db.Fields, ",")
	for _, f := range fields {
		s = append(s, DbMapLeft[db.DbType])
		s = append(s, f)
		s = append(s, DbMapRight[db.DbType])
		s = append(s, ",")
	}
	s = s[0: len(s) - 1]
	s = append(s, " FROM ")
	s = append(s, DbMapLeft[db.DbType])
	s = append(s, db.Table)
	s = append(s, DbMapRight[db.DbType])
	s = append(s, " WHERE ")
	if db.Where == nil {
		s = append(s, "1=1")
	} else {
		s = append(s, db.GetWhere())
	}
	if db.OrderBy != nil {
		s = append(s, " ORDER BY ")
		s = append(s,db.GetOrderBy())
	}
	return strings.Join(s, "")
}

func (*MsSQL) GenSelect (db *DbData) string {
	return Select(db)
}

func (this *MsSQL) GenPage(db *DbData) string {
	if db == nil {
		return ""
	}
	s := make([]string, 0)
	str := this.GenSelect(db)
	s = append(s, str)
	s = append(s, " OFFSET ")
	i := (db.PageIndex - 1) * db.PageIndex
	s = append(s, strconv.Itoa(i))
	s = append(s, " ROWS FETCH NEXT ")
	s = append(s, strconv.Itoa(db.PageSize))
	s = append(s, " ROWS ONLY ")
	return strings.Join(s, "")
}