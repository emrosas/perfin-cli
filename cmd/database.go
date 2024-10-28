package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func initDatabase() {
	db, err := connectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS income (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			description TEXT,
			amount INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
	)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS expenses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			description TEXT,
			amount INTEGER
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
	)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		`CREATE VIEW total_income AS
			SELECT SUM(amount) as income FROM income;`,
	)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		`CREATE VIEW total_expenses AS
			SELECT SUM(amount) as expenses FROM expenses;`,
	)
	if err != nil {
		panic(err)
	}

}

func connectDatabase() (*sql.DB, error) {
	dbPath := filepath.Join(getUserHomeDir(), "perfin-database", "perfin.db")
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		// Create a new database file
		err = os.MkdirAll(filepath.Dir(dbPath), 0755)
		_, err = os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		initDatabase()
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}
	return db, nil
}

func getUserHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return homeDir
}

type Transaction struct {
	ID          int
	Description string
	Amount      int
}
