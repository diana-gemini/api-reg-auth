package types

import (
	"database/sql"
)

func CreateTables(db *sql.DB) error {
	stmt := `	
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email VARCHAR(255) UNIQUE NOT NULL,
			password CHAR(60),
			token TEXT,
			expires DATETIME
		);`

	_, err := db.Exec(stmt)

	return err
}
