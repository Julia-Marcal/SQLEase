package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase(username string, password string, hostname string, port int, databasename string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, databasename)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
