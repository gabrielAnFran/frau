package listexpenses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListExpenses(t *testing.T) {

	err := ListExpenses()
	assert.Nil(t, err)

}
