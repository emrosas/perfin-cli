package cmd

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
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

	_, err = db.Exec(
		`CREATE VIEW total_balance AS
			SELECT (SELECT SUM(amount) FROM income) - (SELECT SUM(amount) FROM expenses) AS balance;`,
	)
	if err != nil {
		panic(err)
	}

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

	_, err := db.Exec("INSERT INTO expenses (description, amount) VALUES (?, ?)", d, a)
	if err != nil {
		panic(err)
	}
}

func queryOverview() (income int, expenses int, balance int) {
	db := connectDatabase()
	defer db.Close()

	// Get income
	var i int
	rows, err := db.Query("SELECT * FROM total_income")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&income)
		if err != nil {
			panic(err)
		}
		income = i
	}

	// Get expenses
	var e int
	rows, err = db.Query("SELECT * FROM total_expenses")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&expenses)
		if err != nil {
			panic(err)
		}
		expenses = e
	}

	// Get balance
	var b int
	rows, err = db.Query("SELECT balance FROM total_balance")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&b)
		if err != nil {
			panic(err)
		}
		balance = b
	}
	return income, expenses, balance
}

func queryIncome() ([]Transaction, error) {
	db := connectDatabase()
	defer db.Close()

	var incomes []Transaction

	// Get income
	rows, err := db.Query("SELECT id, description, amount FROM income")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var income Transaction
		err := rows.Scan(&income.ID, &income.Description, &income.Amount)
		if err != nil {
			return nil, err
		}
		incomes = append(incomes, income)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return incomes, nil
}

type Transaction struct {
	ID          int
	Description string
	Amount      int
}
