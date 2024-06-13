package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxy_Do(t *testing.T) {
	var sub Subject

	// 代理模式
	sub = &Proxy{}
	res := sub.Do()
	assert.Equal(t, "pre:real:after", res)

	// 不使用代理
	sub = &RealSubject{}
	res = sub.Do()
	assert.Equal(t, "real", res)
}
