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
	Long:  `Get an overview of your income. You can insert a value and a description.`,

	Run: getIncome,
}

func init() {
	rootCmd.AddCommand(overviewCmd)
}

func getOverview(cmd *cobra.Command, args []string) {
	fmt.Println("Querying overview...")
	income, expenses, balance := queryOverview()
	fmt.Println("Income: ", income)
	fmt.Println("Expenses: ", expenses)
	fmt.Println("Balance: ", balance)
}

func getIncome(cmd *cobra.Command, args []string) {
	fmt.Println("Querying income...")
	incomes, err := queryIncome()
	if err != nil {
		panic(err)
	}

	for _, income := range incomes {
		fmt.Printf("%d | %s | %d ", income.ID, income.Description, income.Amount)
	}
}
