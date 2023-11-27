package main

import (
	"fmt"

	db "sqlease/database"
	crud "sqlease/operations"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConn, err := db.ConnectToDatabase("username", "My_hard_password_1234!", "localhost", 3306, "mydb")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	crud.Create(dbConn, "my_table", "my_column", "my_value")

	defer dbConn.Close()

}
