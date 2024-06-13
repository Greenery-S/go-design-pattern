package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcreteComponent_Calc(t *testing.T) {
	baseComponent := &ConcreteComponent{}
	assert.Equal(t, 0, baseComponent.Calc())

	mulBaseComponent := WrapMulDecorator(baseComponent, 2)
	assert.Equal(t, 0, mulBaseComponent.Calc())

	addMulBaseComponent := WrapAddDecorator(mulBaseComponent, 3)
	assert.Equal(t, 3, addMulBaseComponent.Calc())

	mulAddMulBaseComponent := WrapMulDecorator(addMulBaseComponent, 4)
	assert.Equal(t, 12, mulAddMulBaseComponent.Calc())
}
