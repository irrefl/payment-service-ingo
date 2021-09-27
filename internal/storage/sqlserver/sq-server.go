package sqlserver

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"payment-service/internal/storage/credentials"
)

var sqlServerType *sql.DB

type SqlServer struct {
	db IDatabase
}

func (SqlServer) GetDB() *sql.DB {
	return sqlServerType
}

func (SqlServer) InitDatabase() bool {

	cb := credentials.NewCredentialBuilder()

	var cr = cb.ServerInfo().
		WithServerName("localhost").
		WithPortNumber(1433).
		WithDbName("goTest").
		UserAuthInfo().
		WithUserName("sa").
		WithPassword("$prod1234").
		GetConnectionString().
		Build()

	var conString = cr.GetConnectionString()
	var err error

	// Create connection pool
	sqlServerType, err = sql.Open("sqlserver", conString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = sqlServerType.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	return true
}
