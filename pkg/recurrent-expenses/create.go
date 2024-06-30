package recurrentexpenses

import (
	"fmt"
)

func CreateRecurrentExpense(name string, amount float64, frequency string) error {
	fmt.Println("Creating recurrent expense:", name, amount, frequency)
	return nil
}
