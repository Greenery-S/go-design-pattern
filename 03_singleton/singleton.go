package _3_singleton

import "sync"

//// * 单例模式

type Singleton interface {
	foo()
}

//// * package中唯一的实例

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 用于获取单例模式对象,懒惰初始化,并发安全
func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

type singleton struct{}

func (s singleton) foo() {}
