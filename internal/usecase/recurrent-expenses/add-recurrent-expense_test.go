package usecase

import (
	"testing"

	"github.com/gabrielAnFran/frauCLI/models"
	"github.com/stretchr/testify/mock"
)

// Mock for RecurrentExpensesRepositoryInterface
type MockRecurrentExpensesRepository struct {
	mock.Mock
}

func (m *MockRecurrentExpensesRepository) AddRecurrentExpenses(expense models.RecurrentExpense) (models.RecurrentExpense, error) {
	args := m.Called(expense)
	return args.Get(0).(models.RecurrentExpense), args.Error(1)
}

func TestAddRecurrentExpenseUsecase_Execute(t *testing.T) {
	// mockRepo := new(MockRecurrentExpensesRepository)
	// usecase := NewAddRecurrentExpanseUseCase(mockRepo)

	// expense := models.RecurrentExpense{
	// 	// fill with appropriate fields
	// }

	// t.Run("success", func(t *testing.T) {
	// 	mockRepo.On("AddRecurrentExpenses", expense).Return(expense, nil)

	// 	result, err := usecase.Execute(expense)

	// 	assert.NoError(t, err)
	// 	assert.Equal(t, &expense, result)
	// 	mockRepo.AssertExpectations(t)
	// })

	// t.Run("failure", func(t *testing.T) {
	// 	mockError := errors.New("some error")
	// 	mockRepo.On("AddRecurrentExpenses", expense).Return(models.RecurrentExpense{}, mockError)

	// 	_, err := usecase.Execute(expense)

	// 	assert.NotNil(t, err)
	// })
}
