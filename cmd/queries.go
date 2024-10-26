package cmd

import (
	"fmt"

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
	fmt.Println("\nHere's an overview on your finances")
	fmt.Printf("Income: $%d\nExpenses: $%d\nBalance: $%d\n", income, expenses, balance)
}

func getIncome(cmd *cobra.Command, args []string) {
	fmt.Println("Querying income...")
	incomes, err := queryIncome()
	if err != nil {
		panic(err)
	}

	for _, income := range incomes {
		fmt.Printf("%d | %s | $%d\n", income.ID, income.Description, income.Amount)
	}

	sum := 0
	for _, income := range incomes {
		sum += income.Amount
	}
	fmt.Printf("Total income: $%d\n", sum)
}

func getExpenses(cmd *cobra.Command, args []string) {
	fmt.Println("Querying expenses...")
	expenses, err := queryExpenses()
	if err != nil {
		panic(err)
	}

	for _, expense := range expenses {
		fmt.Printf("%d | %s | $%d\n", expense.ID, expense.Description, expense.Amount)
	}

	sum := 0
	for _, expense := range expenses {
		sum += expense.Amount
	}
	fmt.Printf("Total expense: $%d\n", sum)
}
