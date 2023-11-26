package main

import (
	"fmt"

	db "sqlease/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.ConnectToDatabase("username", "My_hard_password_1234!", "localhost", 3306, "mydb")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

}
