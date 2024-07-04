package entity

import (
	"errors"

	"github.com/gabrielAnFran/frauCLI/models"
)

func Addrecurrentexpense(expenses models.RecurrentExpense) error {

	if len(expenses.ExpenseName) == 0 {
		return errors.New("name should not be empty")
	}

	return nil

}
