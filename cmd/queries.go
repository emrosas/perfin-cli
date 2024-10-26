package cmd

import (
	"fmt"

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
	fmt.Printf("Income: %s\nExpenses: %s\nBalance: %s\n", ac.FormatMoney(income), ac.FormatMoney(expenses), ac.FormatMoney(balance))
}

func getIncome(cmd *cobra.Command, args []string) {
	fmt.Println("Querying income...")
	incomes, err := queryIncome()
	if err != nil {
		panic(err)
	}

	ac := accounting.Accounting{Symbol: "$"}

	for _, income := range incomes {
		fmt.Printf("%d | %s | %s\n", income.ID, income.Description, ac.FormatMoney(income.Amount))
	}

	sum := 0
	for _, income := range incomes {
		sum += income.Amount
	}
	fmt.Printf("Total income: %s\n", ac.FormatMoney(sum))
}

func getExpenses(cmd *cobra.Command, args []string) {
	fmt.Println("Querying expenses...")
	expenses, err := queryExpenses()
	if err != nil {
		panic(err)
	}

	ac := accounting.Accounting{Symbol: "$"}

	for _, expense := range expenses {
		fmt.Printf("%d | %s | %s\n", expense.ID, expense.Description, ac.FormatMoney(expense.Amount))
	}

	sum := 0
	for _, expense := range expenses {
		sum += expense.Amount
	}
	fmt.Printf("Total expense: %s\n", ac.FormatMoney(sum))
}
