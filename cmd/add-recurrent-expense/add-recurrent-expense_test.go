package addrecurrentexpense

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRecurrentExpense(t *testing.T) {
	expense := Expense{
		ExpenseName: "Test Expense",
		Amount:      100,
		Frequency:   "Monthly",
		StartDate:   "2021-01-01",
		EndDate:     "2021-12-31",
	}

	err := PersistRecurrentExpense(&expense)

	assert.NoError(t, err)
}
