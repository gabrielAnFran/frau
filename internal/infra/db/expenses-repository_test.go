package db

import (
	"testing"

	"github.com/gabrielAnFran/frauCLI/models"
	"github.com/stretchr/testify/assert"
)

func TestAddRecurrentExpenses(t *testing.T) {
	repo := NewRecurrentExpensesRepository()

	expense := models.RecurrentExpense{
		ExpenseName: "Test",
		Amount:      100,
		Frequency:  "Monthly",
		StartDate:   "2021-01-01",
		EndDate:     "2021-12-31",

	}

	result, err := repo.AddRecurrentExpenses(expense)

	assert.Equal(t, result, expense)
	assert.Nil(t, err)

}
