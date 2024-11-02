package cmd

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
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
	idSet := cmd.Flags().Changed("id")
	transactionType, _ := cmd.Flags().GetString("type")
	transactionTypeSet := cmd.Flags().Changed("type")

	if idSet && transactionTypeSet {
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
		case "expense":
			editTransaction(id, description, amount, "expense")
		default:
			fmt.Println("Invalid transaction type. Please use 'income' or 'expense'.")
		}
	} else if !idSet || !transactionTypeSet {
		// If either ID or transaction type is missing, display the form
		displayEditForm()
		return
	} else {
		// If both ID and transaction type are provided, display the form
		// to prompt the user to edit the transaction
		fmt.Println("Please provide both ID and transaction type when using flags.")
		return
	}

	fmt.Println("Transaction edited successfully")
}

func displayEditForm() {
	var (
		id              string
		description     string
		amount          string
		transactionType string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("ID").
				Validate(validateInteger).
				Value(&id),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Transaction Type").
				Options(
					huh.NewOption("Income", "income"),
					huh.NewOption("Expense", "expense"),
				).
				Value(&transactionType),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Description").
				Value(&description),
		),
		huh.NewGroup(

			huh.NewInput().
				Title("Amount").
				Validate(validateInteger).
				Value(&amount),
		),
	)

	err := form.Run()
	if err != nil {
		panic(err)
	}

	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		panic(err)
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	switch transactionType {
	case "income":
		editTransaction(idInt, description, amountInt, "income")
		fmt.Println("Income inserted successfully")
	case "expense":
		editTransaction(idInt, description, amountInt, "expense")
		fmt.Println("Expense inserted successfully")
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
