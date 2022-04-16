package sql

import (
	"strconv"
	"strings"
)

type MySQL struct {}

func (*MySQL) GenSelect (db *DbData) string {
	return Select(db)
}

func (this *MySQL) GenPage(db *DbData) string {
	if db == nil {
		return ""
	}
	s := make([]string, 0)
	str := this.GenSelect(db)
	s = append(s, str)
	s = append(s, " limit ")
	i := (db.PageIndex - 1) * db.PageIndex
	s = append(s, strconv.Itoa(i))
	s = append(s, ",")
	s = append(s, strconv.Itoa(db.PageSize))
	return strings.Join(s, "")
}