package infrastructure

import "database/sql"

func OpenMysqlDB(conectionString string) (*sql.DB, error) {
	return sql.Open("mysql", conectionString)
}
