package cmd

import (
	"database/sql"
	_ "modernc.org/sqlite"
	"os"
)

func initDatabase() {
	db := connectDatabase()
	defer db.Close()

	_, err := db.Exec(
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
			SELECT SUM(amount) FROM income;`,
	)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		`CREATE VIEW total_expenses AS
			SELECT SUM(amount) FROM expenses;`,
	)
	if err != nil {
		panic(err)
	}

	_, err O= db.Exec(
		`-- CREATE VIEW total_balance AS
--   SELECT (SELECT SUM(amount) FROM income) - (SELECT SUM(amount) FROM expenses) AS balance;
`
	)

}

func connectDatabase() *sql.DB {
	_, err := os.Stat("./db.db")
	if os.IsNotExist(err) {
		// Create a new database file
		_, err = os.Create("./db.db")
		if err != nil {
			return nil
		}
		initDatabase()
	}

	db, err := sql.Open("sqlite", "./db.db")
	if err != nil {
		panic(err)
	}
	return db
}

func insertIncomeToDB(d string, a int) {
	db := connectDatabase()
	defer db.Close()

	_, err := db.Exec("INSERT INTO income (description, amount) VALUES (?, ?)", d, a)
	if err != nil {
		panic(err)
	}
}

func insertExpenseToDB(d string, a int) {
	db := connectDatabase()
	defer db.Close()

	_, err := db.Exec("INSERT INTO income (description, amount) VALUES (?, ?)", d, a)
	if err != nil {
		panic(err)
	}
}
