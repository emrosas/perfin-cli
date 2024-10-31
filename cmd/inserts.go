package cmd

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert income or expense",
	Long:  `Insert an income or expense to your balance. You can add a value and a description.`,

	Run: insertHandler,
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringP("description", "d", "", "Description of the income")
	insertCmd.Flags().IntP("amount", "a", 0, "Amount of income")
	insertCmd.Flags().StringP("type", "t", "income", "Type of the transaction (income or expense)")
}

func insertHandler(cmd *cobra.Command, args []string) {
	description, _ := cmd.Flags().GetString("description")
	amount, _ := cmd.Flags().GetInt("amount")
	transactionType, _ := cmd.Flags().GetString("type")

	switch transactionType {
	case "income":
		insertTransactionToDB(description, amount, "income")
		fmt.Println("Income inserted successfully")
	case "expense":
		insertTransactionToDB(description, amount, "expense")
		fmt.Println("Expense inserted successfully")
	default:
		fmt.Println("Invalid transaction type. Please use 'income' or 'expense'.")
	}
}

func insertTransactionToDB(d string, a int, transactionType string) {
	db, err := connectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	switch transactionType {
	case "income":
		_, err = db.Exec("INSERT INTO income (description, amount) VALUES (?, ?)", d, a)
		if err != nil {
			panic(err)
		}
	case "expense":
		_, err = db.Exec("INSERT INTO expenses (description, amount) VALUES (?, ?)", d, a)
		if err != nil {
			panic(err)
		}
	}
}
