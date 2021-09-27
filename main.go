package main

import (
	"payment-service/api"
	"payment-service/api/controllers"
	"payment-service/internal/storage/sqlserver"

	_ "github.com/denisenkom/go-mssqldb"
	_ "payment-service/internal/storage/sqlserver"
)

func main() {

	var sqlServer sqlserver.IDatabase = sqlserver.SqlServer{}
	sqlServer.InitDatabase()

	controller := controllers.NewEmployeeController(sqlServer)
	api.HandleRequests(controller)

}
