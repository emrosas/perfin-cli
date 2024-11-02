package cmd

import (
	"fmt"
	"strconv"

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
	descriptionSet := cmd.Flags().Changed("description")
	amount, _ := cmd.Flags().GetInt("amount")
	amountSet := cmd.Flags().Changed("amount")
	transactionType, _ := cmd.Flags().GetString("type")

	// Check if the user provided both description and amount using flags
	if descriptionSet && amountSet {
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
	} else if !descriptionSet || !amountSet {
		// If either description or amount is missing, display the form
		fmt.Println("Please provide both description and amount when using flags.")
		displayInsertForm()
	} else {
		fmt.Println("Please provide both description and amount when using flags.")
	}
}

func displayInsertForm() {
	var (
		description     string
		amount          string
		transactionType string
	)

	form := huh.NewForm(
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

	switch transactionType {
	case "income":
		insertTransactionToDB(description, amountInt, "income")
		fmt.Println("Income inserted successfully")
	case "expense":
		insertTransactionToDB(description, amountInt, "expense")
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
