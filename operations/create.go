package operations

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func Create(db *sql.DB, table string, columns []string, values []string) error {
	if len(columns) != len(values) {
		return fmt.Errorf("the number of columns and values do not match")
	}

	placeholders := strings.Repeat("?,", len(values))
	placeholders = placeholders[:len(placeholders)-1]

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(columns, ", "), placeholders)

	interfaceValues := make([]interface{}, len(values))
	for i, v := range values {
		interfaceValues[i] = v
	}

	_, err := db.Exec(query, interfaceValues...)
	return err
}
