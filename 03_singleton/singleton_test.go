package _3_singleton

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"testing"
)

func TestGetInstance(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	assert.Equal(t, ins1, ins2)
}

func TestParallelGetInstance(t *testing.T) {
	const parCount = 100
	eg := errgroup.Group{}
	eg.SetLimit(10)
	instances := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		eg.Go(func() error {
			instances[i] = GetInstance()
			return nil
		})
	}
	err := eg.Wait()
	assert.NoError(t, err)

	for i := 1; i < parCount; i++ {
		assert.Equal(t, instances[i], instances[i-1])
	}
}
