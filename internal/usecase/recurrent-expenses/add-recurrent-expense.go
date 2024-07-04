package usecase

import (
	entity "github.com/gabrielAnFran/frauCLI/internal/entity/recurrent-expeses"
	"github.com/gabrielAnFran/frauCLI/models"
)

type AddRecurrentExpenseUsecase struct {
	ExpensesRepository entity.RecurrentExpensesRepositoryInterface
}

func NewAddRecurrentExpanseUseCase(addExpensesRepository entity.RecurrentExpensesRepositoryInterface) *AddRecurrentExpenseUsecase {
	return &AddRecurrentExpenseUsecase{
		ExpensesRepository: addExpensesRepository,
	}
}

func (a *AddRecurrentExpenseUsecase) Execute(expenses models.RecurrentExpense) (*models.RecurrentExpense, error) {

	res, err := a.ExpensesRepository.AddRecurrentExpenses(expenses)
	if err != nil {
		return nil, err
	}

	return &res, err
}
