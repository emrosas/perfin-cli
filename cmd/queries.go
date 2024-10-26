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
