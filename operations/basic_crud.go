package operations

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Create(db *sql.DB, table string, column string, value string) error {
	_, err := db.Exec("INSERT INTO "+table+" ("+column+") VALUES (?)", value)
	return err
}

func Read(db *sql.DB, id int, column string, table string) (string, string, error) {
	var username, email string
	err := db.QueryRow("SELECT "+column+" FROM "+table+" WHERE id = ?", id).Scan(&username, &email)
	if err != nil {
		return "", "", err
	}
	return username, email, nil
}

func Update(db *sql.DB, table string, column string, id int, newValue string) error {
	_, err := db.Exec("UPDATE "+table+" SET "+column+" = ? WHERE id = ?", newValue, id)
	return err
}

func Delete(db *sql.DB, table string, column string, id int) error {
	_, err := db.Exec("DELETE FROM "+table+" WHERE "+column+" = ?", id)
	return err
}
