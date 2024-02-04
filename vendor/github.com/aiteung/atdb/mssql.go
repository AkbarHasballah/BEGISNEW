package atdb

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

func MssqlConnect(mssqlconn DBInfo) (db *sql.DB) {
	db, err := sql.Open("mssql", mssqlconn.DBString)
	if err != nil {
		panic(fmt.Sprintf("MssqlConnect: %v\n", err))
	}
	return db

}
