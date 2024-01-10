package operations

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func Delete(db *sql.DB, table string, column string, id int) error {
	_, err := db.Exec("DELETE FROM "+table+" WHERE "+column+" = ?", id)
	return err
}

func DeleteGroup(db *sql.DB, table string, ids []int) error {
	if len(ids) == 0 {
		return errors.New("no ids provided")
	}

	placeholders := make([]string, len(ids))
	for i := range placeholders {
		placeholders[i] = "?"
	}

	placeholderStr := strings.Join(placeholders, ",")

	query := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", table, placeholderStr)

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	_, err := db.Exec(query, args...)
	return err
}
