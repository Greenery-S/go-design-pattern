package mediator

import (
	"fmt"
	"strings"
)

//// * 中介模式, 用于解耦多个对象之间的交互关系. 假设有CD驱动器, CPU, 显卡, 声卡等设备, 他们之间的交互关系如下:
// 1. CD驱动器读取数据后, 通知中介者数据已经读取完毕
// 2. CPU接收到数据后, 通知中介者数据已经处理完毕
// 3. 显卡接收到数据后, 通知中介者数据已经显示完毕
// 4. 声卡接收到数据后, 通知中介者数据已经播放完毕

var mediator *Mediator

type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

// CDDriver , CD驱动器, 读取数据
type CDDriver struct {
	Data string
}

func (c *CDDriver) ReadData(music, image string) {
	c.Data = strings.Join([]string{music, image}, ",")

	fmt.Printf("CDDriver: reading data %s\n", c.Data)
	GetMediatorInstance().changed(c)
}

// CPU , 处理CD驱动器读取的数据
type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]

	fmt.Printf("CPU: split data with Sound: \"%s\", Video: \"%s\"\n", c.Sound, c.Video)
	GetMediatorInstance().changed(c)
}

// VideoCard , 显卡, 显示CPU处理的图像数据
type VideoCard struct {
	Data string
}

func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("VideoCard: display \"%s\"\n", v.Data)
	GetMediatorInstance().changed(v)
}

// SoundCard , 声卡, 播放CPU处理的音频数据
type SoundCard struct {
	Data string
}

func (s *SoundCard) Play(data string) {
	s.Data = data
	fmt.Printf("SoundCard: play \"%s\"\n", s.Data)
	GetMediatorInstance().changed(s)
}

// GetMediatorInstance , 获取中介者实例, 全局唯一, 可以使用单例模式
func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{
			CD:    &CDDriver{},
			CPU:   &CPU{},
			Video: &VideoCard{},
			Sound: &SoundCard{},
		}
	}
	return mediator
}

// changed , 中介者接收到设备通知后, 调用(通知)其他设备, 做下一步操作
func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver: // CD驱动器读取数据后, 中介者收到通知, 通知CPU处理数据
		m.CPU.Process(inst.Data)
	case *CPU: // CPU处理数据后, 中介者收到通知, 通知显卡显示数据
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)
	case *VideoCard, *SoundCard: // 显卡或声卡显示数据后, 中介者收到通知, 无需通知其他设备
		fmt.Printf("Mediator: %T finished\n", i)
	}
}
