package operations

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Update(db *sql.DB, table string, column string, newValue string, id int) error {
	_, err := db.Exec("UPDATE "+table+" SET "+column+" = ? WHERE id = ?", newValue, id)
	return err
}
