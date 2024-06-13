package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumbers_Iterator(t *testing.T) {
	var aggr Aggregate
	var iter Iterator

	aggr = NewNumbers(3, 10)
	iter = aggr.Iterator()

	assert.Equal(t, 3, iter.Next().(int))
	assert.Equal(t, 4, iter.Next().(int))

	i := 3
	for iter.First(); !iter.IsDone(); i++ {
		c := iter.Next()
		assert.Equal(t, i, c.(int))
	}

	assert.Nil(t, iter.Next())
}
