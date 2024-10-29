package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an income or expense",
	Long:  `Edit every aspect of a transaction. Choose between income or expense.`,

	Run: editHandler,
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().IntP("id", "i", 0, "ID of the transaction")
	editCmd.Flags().StringP("type", "t", "income", "Type of the transaction (income or expense)")
	editCmd.Flags().StringP("description", "d", "", "Description of the income")
	editCmd.Flags().IntP("amount", "a", 0, "Amount of income")
}

func editHandler(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetInt("id")
	transactionType, _ := cmd.Flags().GetString("type")

	var description string
	var amount int
	descriptionFlag, _ := cmd.Flags().GetString("description")
	if descriptionFlag != "" {
		description = descriptionFlag
	} else {
		transaction, err := getTransactionById(id, transactionType)
		if err != nil {
			fmt.Println("Error getting transaction:", err)
			return
		}
		description = transaction.Description
	}

	amountFlag, _ := cmd.Flags().GetInt("amount")
	if amountFlag != 0 {
		amount = amountFlag
	} else {
		transaction, err := getTransactionById(id, transactionType)
		if err != nil {
			fmt.Println("Error getting transaction:", err)
			return
		}
		amount = transaction.Amount
	}

	switch transactionType {
	case "income":
		editTransaction(id, description, amount, "income")
		fmt.Println("Income edited successfully")
	case "expense":
		editTransaction(id, description, amount, "expense")
		fmt.Println("Expense edited successfully")
	default:
		fmt.Println("Invalid transaction type. Please use 'income' or 'expense'.")
	}
}

func getTransactionById(id int, transactionType string) (Transaction, error) {
	db, err := connectDatabase()
	if err != nil {
		return Transaction{}, err
	}
	defer db.Close()

	var transaction Transaction
	var query string
	switch transactionType {
	case "income":
		query = "SELECT id, description, amount FROM income WHERE id = ?"
	case "expense":
		query = "SELECT id, description, amount FROM expenses WHERE id = ?"
	default:
		return Transaction{}, fmt.Errorf("invalid transaction type: %s", transactionType)
	}

	err = db.QueryRow(query, id).Scan(&transaction.ID, &transaction.Description, &transaction.Amount)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func editTransaction(id int, description string, amount int, transactionType string) {
	db, err := connectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	switch transactionType {
	case "income":
		_, err = db.Exec("UPDATE income SET description = ?, amount = ? WHERE id = ?", description, amount, id)
		if err != nil {
			panic(err)
		}
	case "expense":
		_, err = db.Exec("UPDATE expenses SET description = ?, amount = ? WHERE id = ?", description, amount, id)
		if err != nil {
			panic(err)
		}
	}
}
