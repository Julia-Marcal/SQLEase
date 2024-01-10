package operations

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Read(db *sql.DB, column string, table string, id int) (string, error) {
	err := db.QueryRow("SELECT "+column+" FROM "+table+" WHERE id = ?", id).Scan(&column)
	if err != nil {
		return "", err
	}
	return column, nil
}
