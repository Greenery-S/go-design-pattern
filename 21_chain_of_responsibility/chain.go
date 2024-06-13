package chain

import "fmt"

//// * 责任链模式

// 责任链模式是一种行为设计模式， 允许你将请求沿着处理者链进行发送， 直至其中一个处理者对其进行处理。
// 责任链模式会为请求创建一个接收者对象的链。 这种模式给予请求的发送者和接收者更多的灵活性。

// Manager 定义处理请求的接口
type Manager interface {
	HaveRight(money int) bool                     // 判断是否有权限处理请求
	HandleFeeRequest(name string, money int) bool // 有权限,就可以处理请求
}

// RequestChain 责任链
// 1. Manager: 处理请求的接口
// 2. successor: 后继者
type RequestChain struct {
	Manager
	successor *RequestChain
}

// SetSuccessor 设置后继者, 初始化时设置
func (r *RequestChain) SetSuccessor(m *RequestChain) {
	r.successor = m
}

// HandleFeeRequest 处理请求
// 1. 如果有权限,就处理请求 (递归出口)
// 2. 如果没有权限,就交给后继者处理 (递归)
func (r *RequestChain) HandleFeeRequest(name string, money int) bool {
	if r.Manager.HaveRight(money) {
		return r.Manager.HandleFeeRequest(name, money)
	}
	if r.successor != nil {
		return r.successor.HandleFeeRequest(name, money)
	}
	return false
}

// 可以不显式定义, 内嵌接口即可
//func (r *RequestChain) HaveRight(money int) bool {
//	return true
//}

//// * 实现

// ProjectManager 项目经理, 处理500以下的请求
type ProjectManager struct{}

func NewProjectManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &ProjectManager{},
	}
}

func (*ProjectManager) HaveRight(money int) bool {
	return money < 500
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	if name == "bob" {
		fmt.Printf("Project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Project manager don't permit %s %d fee request\n", name, money)
	return false
}

// DepManager 部门经理, 处理5000以下的请求
type DepManager struct{}

func NewDepManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &DepManager{},
	}
}

func (*DepManager) HaveRight(money int) bool {
	return money < 5000
}

func (*DepManager) HandleFeeRequest(name string, money int) bool {
	if name == "tom" {
		fmt.Printf("Dep manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Dep manager don't permit %s %d fee request\n", name, money)
	return false
}

// GeneralManager 总经理, 处理5000以上的请求
type GeneralManager struct{}

func NewGeneralManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &GeneralManager{},
	}
}

func (*GeneralManager) HaveRight(money int) bool {
	return true
}

func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	if name == "ada" {
		fmt.Printf("General manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("General manager don't permit %s %d fee request\n", name, money)
	return false
}
