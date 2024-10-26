package cmd

import (
	"database/sql"
	"fmt"
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

func queryOverview() (income, expenses, balance int, err error) {
	db := connectDatabase()
	defer db.Close()

	// Get income
	var incomeNullable sql.NullInt64
	err = db.QueryRow("SELECT income FROM total_income").Scan(&incomeNullable)
	if err != nil {
		return 0, 0, 0, err
	}
	if incomeNullable.Valid {
		income = int(incomeNullable.Int64)
	} else {
		fmt.Println("No income found")
	}

	// Get expenses
	var expensesNullable sql.NullInt64
	err = db.QueryRow("SELECT expenses FROM total_expenses").Scan(&expensesNullable)
	if err != nil {
		return 0, 0, 0, err
	}
	if expensesNullable.Valid {
		expenses = int(expensesNullable.Int64)
	} else {
		fmt.Println("No expenses found")
	}

	// Get balance
	var balanceNullable sql.NullInt64
	err = db.QueryRow("SELECT balance FROM total_balance").Scan(&balanceNullable)
	if err != nil {
		return 0, 0, 0, err
	}
	if balanceNullable.Valid {
		balance = int(balanceNullable.Int64)
	}

	return income, expenses, balance, nil
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

func queryExpenses() ([]Transaction, error) {
	db := connectDatabase()
	defer db.Close()

	var expenses []Transaction

	// Get expense
	rows, err := db.Query("SELECT id, description, amount FROM expenses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var expense Transaction
		err := rows.Scan(&expense.ID, &expense.Description, &expense.Amount)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

type Transaction struct {
	ID          int
	Description string
	Amount      int
}
