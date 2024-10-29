package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an income or expense",
	Long:  `Delete an income or expense from your balance. Enter an ID and type of the transaction.`,

	Run: deleteHandler,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntP("id", "i", 0, "ID of the transaction")
	deleteCmd.Flags().StringP("type", "t", "income", "Type of the transaction (income or expense)")
}

func deleteHandler(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetInt("id")
	transactionType, _ := cmd.Flags().GetString("type")
	err := deleteTransaction(id, transactionType)
	if err != nil {
		fmt.Println("Error deleting transaction:", err)
		return
	}

	fmt.Println("Transaction deleted successfully")
}

func deleteTransaction(id int, transactionType string) error {
	db, err := connectDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	var query string
	switch transactionType {
	case "income":
		query = "DELETE FROM income WHERE id = ?"
	case "expense":
		query = "DELETE FROM expenses WHERE id = ?"
	default:
		return fmt.Errorf("invalid transaction type: %s", transactionType)
	}

	_, err = db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
