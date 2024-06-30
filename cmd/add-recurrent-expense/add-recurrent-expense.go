package addrecurrentexpense

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

type Expense struct {
	ExpenseName string
	Amount      float64
	Frequency   string
	StartDate   string
	EndDate     string
}

func NewCmdAddRecurrentExpenses() *cobra.Command {
	expense := &Expense{}
	cmd := cobra.Command{
		Use:     "add-recurrent-expense",
		Short:   "Add a recurrent expense",
		Long:    "Add a recurrent expense",
		Example: "frau add-recurrent-expense --name='Netflix' --amount=9.99 --frequency='Monthly' --start-date='2021-01-01' --end-date='2021-12-31'",

		RunE: func(cmd *cobra.Command, args []string) error {

			err := PersistRecurrentExpense(expense)
			if err != nil {
				return err
			}

			spew.Dump(expense)

			return nil
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&expense.ExpenseName, "name", "", "Expense name")
	flags.Float64Var(&expense.Amount, "amount", 0, "Expense amount")
	flags.StringVar(&expense.Frequency, "frequency", "", "Expense frequency")
	flags.StringVar(&expense.StartDate, "start-date", "", "Expense start date")
	flags.StringVar(&expense.EndDate, "end-date", "", "Expense end date")
	return &cmd

}

func PersistRecurrentExpense(expense *Expense) error {

	spew.Dump(expense)

	// Persistance logic here

	return nil
}
