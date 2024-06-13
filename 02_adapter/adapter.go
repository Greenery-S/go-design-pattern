package _2_adapter

//// * 适配目标

// Target 是适配的目标接口
type Target interface {
	Request() string
}

//// * 原始接口

// Adaptee 是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

type adapteeImpl struct{}

func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

//// * 适配器

// NewAdapter , adapter 是转换Adaptee为Target接口的适配器
// Adaptee --adapter适配--> Target
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

type adapter struct {
	Adaptee // 组合
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
