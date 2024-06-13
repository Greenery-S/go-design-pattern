package bridge

import "fmt"

//// * 桥接模式

// 桥接模式是一种结构型设计模式， 可将一个大类或一系列紧密相关的类拆分为抽象和实现两个独立的层次结构， 从而能在开发时分别使用。
// 该模式有助于组合不同的抽象和实现， 并独立地对它们进行修改。
// 通过将类的实现细节从其抽象表示中分离， 可以使其更容易修改和理解。

// AbstractMessage 是消息发送类的抽象接口
type AbstractMessage interface {
	SendMessage(text, to string)
}

// MessageImplementer 是消息发送类的实现接口
type MessageImplementer interface {
	Send(text, to string)
}

// MessageSMS 是消息发送实现类的实现
type MessageSMS struct{}

// ViaSMS 是消息发送实现类的工厂方法
func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send \"%s to %s\" via SMS\n", text, to)
}

// MessageEmail 是消息发送实现类的实现
type MessageEmail struct{}

// ViaEmail 是消息发送实现类的工厂方法
func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send \"%s to %s\" via Email\n", text, to)
}

// CommonMessage 是消息发送抽象类的实现
// 该类将消息发送实现类的实例作为成员变量， 并在其方法中调用实现类的方法
type CommonMessage struct {
	method MessageImplementer
}

// NewCommonMessage 是 CommonMessage 类的构造函数
// 该函数接受一个消息发送实现类的实例作为参数,实现桥接
func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

// UrgencyMessage 是消息发送抽象类的实现
// 该类将消息发送实现类的实例作为成员变量， 并在其方法中调用实现类的方法
type UrgencyMessage struct {
	method MessageImplementer
}

// NewUrgencyMessage 是 UrgencyMessage 类的构造函数
// 该函数接受一个消息发送实现类的实例作为参数,实现桥接
func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}
