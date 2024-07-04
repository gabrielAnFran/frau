package addrecurrentexpense

import (
	"fmt"
	"os"
	"text/tabwriter"

	entity "github.com/gabrielAnFran/frauCLI/internal/entity/recurrent-expeses"
	"github.com/gabrielAnFran/frauCLI/internal/infra/db"
	usecase "github.com/gabrielAnFran/frauCLI/internal/usecase/recurrent-expenses"
	"github.com/gabrielAnFran/frauCLI/models"
	"github.com/spf13/cobra"
)

func NewCmdAddRecurrentExpenses() *cobra.Command {
	expense := &models.RecurrentExpense{}
	cmd := cobra.Command{
		Use:     "add-recurrent-expense",
		Short:   "Add a recurrent expense",
		Long:    "Add a recurrent expense",
		Example: "frau add-recurrent-expense --name='Netflix' --amount=9.99 --frequency='Monthly' --start-date='2021-01-01' --end-date='2021-12-31'",

		RunE: func(cmd *cobra.Command, args []string) error {
			repository := entity.RecurrentExpensesRepositoryInterface(db.NewRecurrentExpensesRepository())
			recurrentExpensesHandler := NewAddRecurrentExpenseHandler(repository)
			err := recurrentExpensesHandler.AddRecurrentExpense(*expense)

			return err

		}}

	flags := cmd.Flags()
	flags.StringVar(&expense.ExpenseName, "name", "", "Expense name")
	flags.Float64Var(&expense.Amount, "amount", 0, "Expense amount")
	flags.StringVar(&expense.Frequency, "frequency", "", "Expense frequency")
	flags.StringVar(&expense.StartDate, "start-date", "", "Expense start date")
	flags.StringVar(&expense.EndDate, "end-date", "", "Expense end date")
	return &cmd

}

type RecurrentExpensesHandler struct {
	RecurrenteExpensesRepository entity.RecurrentExpensesRepositoryInterface
}

func NewAddRecurrentExpenseHandler(addExpensesRepository entity.RecurrentExpensesRepositoryInterface) *RecurrentExpensesHandler {
	return &RecurrentExpensesHandler{
		RecurrenteExpensesRepository: addExpensesRepository,
	}
}

func (h *RecurrentExpensesHandler) AddRecurrentExpense(expense models.RecurrentExpense) error {
	expensesUsecase := usecase.NewAddRecurrentExpanseUseCase(h.RecurrenteExpensesRepository)

	err := entity.Addrecurrentexpense(expense)
	if err != nil {
		return err
	}

	res, err := expensesUsecase.Execute(expense)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	fmt.Fprintf(w, "%s\t%.2f\t%s\t%s\t%s\t\n", res.ExpenseName, res.Amount, res.Frequency, res.StartDate, res.EndDate)

	w.Flush()
	return nil
}
