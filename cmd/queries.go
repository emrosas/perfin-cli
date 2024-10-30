package cmd

import (
	"database/sql"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/leekchan/accounting"
	"github.com/spf13/cobra"
)

var overviewCmd = &cobra.Command{
	Use:   "overview",
	Short: "Get overview",
	Long:  `Get an overview of your balance. You can insert a value and a description.`,

	Run: getOverview,
}

var queryIncomeCmd = &cobra.Command{
	Use:   "income",
	Short: "Get income overview",
	Long:  `Get an overview of your income. Listing all income with description and amount, as well as a total.`,

	Run: getIncome,
}

var queryExpensesCmd = &cobra.Command{
	Use:   "expenses",
	Short: "Get expenses overview",
	Long:  `Get an overview of your expenses. Listing all expenses with description and amount, as well as a total.`,

	Run: getExpenses,
}

func init() {
	rootCmd.AddCommand(overviewCmd)
	rootCmd.AddCommand(queryIncomeCmd)
	rootCmd.AddCommand(queryExpensesCmd)
}

func getOverview(cmd *cobra.Command, args []string) {
	fmt.Println("Querying overview...")
	income, expenses, balance, err := queryOverview()
	if err != nil {
		panic(err)
	}
	ac := accounting.Accounting{Symbol: "$"}
	fmt.Println("\nHere's an overview on your finances")

	tbl := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("CATEGORY", "AMOUNT")

	tbl.Row("Income", ac.FormatMoney(income))
	tbl.Row("Expenses", ac.FormatMoney(expenses))
	tbl.Row("Balance", ac.FormatMoney(balance))

	fmt.Println(tbl)
}

func getIncome(cmd *cobra.Command, args []string) {
	fmt.Println("Querying income...")
	incomes, err := queryIncome()
	if err != nil {
		panic(err)
	}

	ac := accounting.Accounting{Symbol: "$"}
	var rows [][]string

	for _, income := range incomes {
		rows = append(rows, []string{fmt.Sprintf("%d", income.ID), income.Description, ac.FormatMoney(income.Amount)})
	}

	tbl := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#7cab00"))).
		Headers("ID", "DESCRIPTION", "AMOUNT").
		Rows(rows...)

	sum := 0
	for _, income := range incomes {
		sum += income.Amount
	}

	tbl.Row("", "", "")
	tbl.Row("", "Total Income", ac.FormatMoney(sum))

	fmt.Println(tbl)

}

func getExpenses(cmd *cobra.Command, args []string) {
	fmt.Println("Querying expenses...")
	expenses, err := queryExpenses()
	if err != nil {
		panic(err)
	}

	ac := accounting.Accounting{Symbol: "$"}

	var rows [][]string

	for _, expense := range expenses {
		rows = append(rows, []string{fmt.Sprintf("%d", expense.ID), expense.Description, ac.FormatMoney(expense.Amount)})
	}

	tbl := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#da3f3f"))).
		Headers("ID", "DESCRIPTION", "AMOUNT").
		Rows(rows...)

	sum := 0
	for _, expense := range expenses {
		sum += expense.Amount
	}

	tbl.Row("", "", "")
	tbl.Row("", "Total Expenses", ac.FormatMoney(sum))

	fmt.Println(tbl)

}

func queryOverview() (income, expenses, balance int, err error) {
	db, err := connectDatabase()
	if err != nil {
		return 0, 0, 0, err
	}
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

	// Calculate balance
	balance = income - expenses

	return income, expenses, balance, nil
}

func queryIncome() ([]Transaction, error) {
	db, err := connectDatabase()
	if err != nil {
		return nil, err
	}

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
	db, err := connectDatabase()
	if err != nil {
		return nil, err
	}

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
