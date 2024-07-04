package db

import "github.com/gabrielAnFran/frauCLI/models"

type RecurrentExpensesRepository struct{}

func NewRecurrentExpensesRepository() *RecurrentExpensesRepository {
	return &RecurrentExpensesRepository{}
}

func (a *RecurrentExpensesRepository) AddRecurrentExpenses(expense models.RecurrentExpense) (models.RecurrentExpense, error) {

	return expense, nil

}
