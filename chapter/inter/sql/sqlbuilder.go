package sql

type SqlBuilder struct {}

type _sqlBuilder struct {
	IBuilder
}

func (*SqlBuilder) GenSelectSQL(db *DbData) string {
	if db == nil {
		return ""
	}
	var sqlBuilder IBuilder
	switch db.DbType {
	case MYSQL:
		mysql := MySQL{}
		sqlBuilder = _sqlBuilder{&mysql}

	case MSSQL:
		mssql := MsSQL{}
		sqlBuilder = _sqlBuilder{&mssql}
	default:
		return ""
	}
	sql := sqlBuilder.GenSelect(db)
	return sql
}

func (*SqlBuilder) GenPageSQL(db *DbData) string {
	if db == nil {
		return ""
	}
	var sqlBuilder IBuilder
	switch db.DbType {
	case MYSQL:
		mysql := MySQL{}
		sqlBuilder = _sqlBuilder{&mysql}
	case MSSQL:
		mssql := MsSQL{}
		sqlBuilder = _sqlBuilder{&mssql}
	default:
		return ""
	}
	sql := sqlBuilder.GenPage(db)
	return sql
}