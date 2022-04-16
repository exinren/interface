package sql

type SqlBuilder struct {}

type _sqlBuilder struct {
	IBuilder
}

func (*SqlBuilder) GenSelectSQL(db *DbData) string {
	if db == nil {
		return ""
	}
	switch db.DbType {
	case MYSQL:
		mysql := MySQL{}
		sqlBuilder := _sqlBuilder{&mysql}
		sql := sqlBuilder.GenSelect(db)
		return sql
	case MSSQL:
		mssql := MsSQL{}
		sqlBuilder := _sqlBuilder{&mssql}
		sql := sqlBuilder.GenSelect(db)
		return sql
	default:
		return ""
	}
}

func (*SqlBuilder) GenPageSQL(db *DbData) string {
	if db == nil {
		return ""
	}
	switch db.DbType {
	case MYSQL:
		mysql := MySQL{}
		sqlBuilder := _sqlBuilder{&mysql}
		sql := sqlBuilder.GenPage(db)
		return sql
	case MSSQL:
		mssql := MsSQL{}
		sqlBuilder := _sqlBuilder{&mssql}
		sql := sqlBuilder.GenPage(db)
		return sql
	default:
		return ""
	}
}