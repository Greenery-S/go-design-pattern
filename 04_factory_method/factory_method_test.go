package _4_factory_method

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func computeReport(factory OperatorFactory, a, b int) string {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return fmt.Sprintf("%s = %d", op.ToString(), op.Result())
}

func TestOperator(t *testing.T) {
	var factory OperatorFactory
	factory = PlusOperatorFactory{}
	assert.Equal(t, "1 + 2 = 3", computeReport(factory, 1, 2))
	factory = MinusOperatorFactory{}
	assert.Equal(t, "4 - 2 = 2", computeReport(factory, 4, 2))
}
