package mediator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMediator(t *testing.T) {
	m := GetMediatorInstance()

	// Toggle, 触发播放全流程: CD驱动器读取数据, CPU处理数据, 显卡显示数据, 声卡播放数据
	m.CD.ReadData("We are the Champion", "Kong fu Panda 3")

	assert.Equal(t, "We are the Champion,Kong fu Panda 3", m.CD.Data)
	assert.Equal(t, "We are the Champion", m.CPU.Sound)
	assert.Equal(t, "Kong fu Panda 3", m.CPU.Video)
	assert.Equal(t, "Kong fu Panda 3", m.Video.Data)
	assert.Equal(t, "We are the Champion", m.Sound.Data)
}
