package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gabrielAnFran/frauCLI/models"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// test
type RecurrentExpensesRepository struct{}

func NewRecurrentExpensesRepository() *RecurrentExpensesRepository {
	return &RecurrentExpensesRepository{}
}

func (a *RecurrentExpensesRepository) AddRecurrentExpenses(expense models.RecurrentExpense) (models.RecurrentExpense, error) {

	db, err := dbConnect()
	if err != nil {
		return models.RecurrentExpense{}, err
	}
	defer db.Close()

	expenseAmoutString := fmt.Sprintf("%f", expense.Amount)
	expenseClientIDString := "0"

	_, err = db.Exec(`INSERT INTO expenses (expense_name, amount, frequency, start_date, end_date, client_id) VALUES (?, ?, ?, ?, ?, ?)`,
		expense.ExpenseName, expenseAmoutString, expense.Frequency, expense.StartDate, expense.EndDate, expenseClientIDString)
	if err != nil {
		return models.RecurrentExpense{}, err
	}


	return expense, nil

}

func dbConnect() (*sql.DB, error) {
	// Connect to database
	dbURL := os.Getenv("TURSO_DATABASE_URL")
	db_token := os.Getenv("TURSO_AUTH_TOKEN")
	url :=  dbURL + ".turso.io?authToken=" + db_token

	spew.Dump(url)

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		return nil, err
	}

	return db, nil
}
