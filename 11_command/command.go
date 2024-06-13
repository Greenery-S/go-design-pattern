package command

import "fmt"

//// * 命令模式, 有各种命令, 有一个执行者, 有一个调用者, 调用者调用命令, 命令执行者执行命令
// 命令执行者能够执行所有类型的命令, 调用者调用设定的命令, 等待执行结果

// MotherBoard 主板, 执行者, 能够执行所有命令
type MotherBoard struct{}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

//// * 命令接口, 所有命令都要实现这个接口

// Command 命令接口, 所有命令都要实现这个接口
type Command interface {
	Execute()
}

// StartCommand 开机命令
type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

// RebootCommand 重启命令
type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

//// * 调用者, 调用命令, 等待执行结果: Box上的按钮可以绑定不同的命令, 按下按钮后, 命令执行

// Box 调用者, 调用命令, 等待执行结果
type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
