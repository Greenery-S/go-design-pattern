package _0_simple_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHiAPI_Say(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	assert.Equal(t, "Hi, Tom", s)
}

func TestHelloAPI_Say(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	assert.Equal(t, "Hello, Tom", s)
}
