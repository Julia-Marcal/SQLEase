package main

import (
	"fmt"

	db "sqlease/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.ConnectToDatabase("username", "password", "hostname", 1234, "databasename")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

}
