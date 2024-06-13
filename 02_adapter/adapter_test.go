package _2_adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expect = "adaptee method"

func TestAdapter_Request(t *testing.T) {
	adaptee := NewAdaptee()       // 被适配的原接口
	target := NewAdapter(adaptee) // 适配器,将原接口转换为目标接口
	res := target.Request()
	assert.Equal(t, expect, res)
}
