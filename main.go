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

	err = crud.Create(dbConn, "users", []string{"name", "age"}, []string{"john doe", "18"})
	if err != nil {
		fmt.Println("Error creating record:", err)
	}

	defer dbConn.Close()

}
