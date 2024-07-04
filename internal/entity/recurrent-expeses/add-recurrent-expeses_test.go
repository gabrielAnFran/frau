package entity

import (
	"errors"
	"testing"

	"github.com/gabrielAnFran/frauCLI/models"
	"github.com/stretchr/testify/assert"
)

func TestAddrecurrentexpense(t *testing.T) {
	expenses := models.RecurrentExpense{
		ExpenseName: "",
	}

	assert.Equal(t, Addrecurrentexpense(expenses), errors.New("name should not be empty"))
}
