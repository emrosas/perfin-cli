package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var incomeCmd = &cobra.Command{
	Use:   "income",
	Short: "Insert income",
	Long:  `Insert an income to your balance. You can insert a value and a description.`,

	Run: insertIncome,
}

var expenseCmd = &cobra.Command{
	Use:   "expense",
	Short: "Insert expense",
	Long:  `Insert an expense to your balance. You can insert a value and a description.`,

	Run: insertExpense,
}

func init() {
	rootCmd.AddCommand(incomeCmd)
	incomeCmd.Flags().StringP("description", "d", "", "Description of the income")
	incomeCmd.Flags().IntP("amount", "a", 0, "Amount of income")

	rootCmd.AddCommand(expenseCmd)
	expenseCmd.Flags().StringP("description", "d", "", "Description of the expense")
	expenseCmd.Flags().IntP("amount", "a", 0, "Amount of expense")
}

func insertIncome(cmd *cobra.Command, args []string) {
	description, _ := cmd.Flags().GetString("description")
	amount, _ := cmd.Flags().GetInt("amount")
	insertIncomeToDB(description, amount)
	fmt.Println("Income inserted successfully")
}

func insertExpense(cmd *cobra.Command, args []string) {
	description, _ := cmd.Flags().GetString("description")
	amount, _ := cmd.Flags().GetInt("amount")
	insertExpenseToDB(description, amount)
	fmt.Println("Expense inserted successfully")
}
