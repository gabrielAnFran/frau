package entity

import (

	"github.com/gabrielAnFran/frauCLI/models"
)

type RecurrentExpensesRepositoryInterface interface {
	AddRecurrentExpenses(models.RecurrentExpense) (models.RecurrentExpense, error)
}
