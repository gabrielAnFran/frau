package usecase

import (
	"fmt"
	"testing"

	"github.com/gabrielAnFran/frauCLI/models"
)

func TestAddRecurrentExpense(t *testing.T) {
	expense := models.RecurrentExpense{
		ExpenseName: "Test Expense",
		Amount:      100,
		Frequency:   "Monthly",
		StartDate:   "2021-01-01",
		EndDate:     "2021-12-31",
	}

	fmt.Println(expense)
	// err := PersistRecurrentExpense(&expense)

	// assert.NoError(t, err)
}
