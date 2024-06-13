package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilder(t *testing.T) {
	var (
		director *Director
		builder1 = &Builder1{}
		builder2 = &Builder2{}
	)

	director = NewDirector(builder1)
	director.Construct()
	assert.Equal(t, "123", builder1.GetResult())

	director = NewDirector(builder2)
	director.Construct()
	assert.Equal(t, 6, builder2.GetResult())
}
