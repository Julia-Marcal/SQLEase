package main

import (
	"fmt"

	db "sqlease/database"
	crud "sqlease/operations"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connDetails := db.ConnectionString()

	dbConn, err := db.ConnectToDatabase(connDetails.Username, connDetails.Password, connDetails.Hostname, connDetails.Port, connDetails.Database)

	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	crud.Create(dbConn, "my_table", "my_column", "my_value")

	defer dbConn.Close()

}
