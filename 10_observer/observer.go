package observer

import "fmt"

//// * 观察者模式, 也叫发布订阅模式, 定义了一种一对多的依赖关系, 让多个观察者对象同时监听某一个主题对象, 当主题对象发生变化时, 它的所有观察者都会收到通知并更新

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

// Subject 主题内容发布者
type Subject struct {
	observers []Observer
	content   string // 订阅内容
}

// Attach 增加观察者
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

// 通过调用Update(),通知所有观察者
func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

// Observer 观察者, 用于接收主题内容
type Observer interface {
	Update(*Subject)
}

// NewPaperReader 不同类型的观察者
type NewPaperReader struct {
	name          string
	latestContent string
}

func (r *NewPaperReader) Update(s *Subject) {
	r.latestContent = s.content
	fmt.Printf("%s receive \"%s\" from newspaper\n", r.name, s.content)
}

// MobileReader 不同类型的观察者
type MobileReader struct {
	name          string
	latestContent string
}

func (r *MobileReader) Update(s *Subject) {
	r.latestContent = s.content
	fmt.Printf("%s receive \"%s\" from mobile\n", r.name, s.content)
}
