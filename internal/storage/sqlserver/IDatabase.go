package sqlserver

import "database/sql"

type IDatabase interface {
	InitDatabase() bool
	GetDB() *sql.DB
}
