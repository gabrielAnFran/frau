package listexpenses

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

type Expense struct {
	ExpenseName string
	Amount      float64
	Frequency   string
	StartDate   string
	EndDate     string
}

func NewCmdListExpenses() *cobra.Command {
	return &cobra.Command{
		Use:     "list-expenses",
		Short:   "List all expenses",
		Long:    "List all expenses",
		Example: "frau list-expenses",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := ListExpenses()
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func ListExpenses() error {

	// TODO: Get the expenses from the database
	expenses := []Expense{
		{ExpenseName: "Netflix",
			Amount:    9.99,
			Frequency: "Monthly",
			StartDate: "2021-01-01",
			EndDate:   "2021-12-31",
		},
		{ExpenseName: "Supermarket",
			Amount:    250.0,
			Frequency: "Monthly",
			StartDate: "2021-01-01",
			EndDate:   "2021-12-31",
		},
		{ExpenseName: "Gym",
			Amount:    40.0,
			Frequency: "Monthly",
			StartDate: "2021-01-01",
			EndDate:   "2021-12-31",
		},
	}

	// Create a new tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "ExpenseName\tAmount\tFrequency\tStartDate\tEndDate\t")

	for _, expense := range expenses {
		fmt.Fprintf(w, "%s\t%.2f\t%s\t%s\t%s\t\n", expense.ExpenseName, expense.Amount, expense.Frequency, expense.StartDate, expense.EndDate)

	}

	w.Flush()

	return nil
}
