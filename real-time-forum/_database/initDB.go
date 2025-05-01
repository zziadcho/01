package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./forum.db?_foreign_keys=on")

	if err != nil {
		return fmt.Errorf("error opening the database: %v", err)
	}

	if err = createTables(); err != nil {
		return fmt.Errorf("error creating the tables: %v", err)
	}
	return nil
}

func createTables() error {
	postTable := `
	CREATE TABLE IF NOT EXISTS PostTable (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(70) NOT NULL,
    content TEXT NOT NULL,
    creation_date DATETIME DEFAULT CURRENT_TIMESTAMP
);
	`
	tables := []string{postTable}

	for _, table := range tables {
		if _, err := DB.Exec(table); err != nil {
			return fmt.Errorf("error creating table: %v", err)
		}
	}
	return nil
}
