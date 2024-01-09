package main

import (
	db "sqlease/database"
	crud "sqlease/operations"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connDetails := db.ConnectionString()

	dbConn, errConn := db.ConnectToDatabase(connDetails.Username, connDetails.Password, connDetails.Hostname, connDetails.Port, connDetails.Database)

	if errConn != nil {
		panic("Error connecting to the database: " + errConn.Error())
	}

	errCreate := crud.Create(dbConn, "users", []string{"name", "age"}, []string{"jayson tatum", "25"})
	if errCreate != nil {
		panic("Error when creating record: " + errCreate.Error())
	}

	_, errRead := crud.Read(dbConn, "name", "users", 1)
	if errRead != nil {
		panic("Error when reading record: " + errRead.Error())
	}

	errUpdate := crud.Update(dbConn, "users", "name", "james", 1)
	if errUpdate != nil {
		panic("Error when updating record: " + errUpdate.Error())
	}

	errDelete := crud.Delete(dbConn, "users", "name", 99)
	if errDelete != nil {
		panic("Error when reading record: " + errDelete.Error())
	}

	defer dbConn.Close()

}
