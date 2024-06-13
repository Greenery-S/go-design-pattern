package decorator

//// * 装饰器模式

// Component 定义一个对象接口
type Component interface {
	Calc() int
}

// ConcreteComponent 实现对象接口的基础对象
type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

// MulDecorator 乘法装饰器,由于内嵌,也是一个Component
type MulDecorator struct {
	Component
	num int
}

// WrapMulDecorator 包装一个Component,返回一个MulDecorator,是一个构造函数
func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

// Calc 实现Component接口的方法,返回乘法结果,被装饰的部分
func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

// AddDecorator 加法装饰器,由于内嵌,也是一个Component
type AddDecorator struct {
	Component
	num int
}

// WrapAddDecorator 包装一个Component,返回一个AddDecorator,是一个构造函数
func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

// Calc 实现Component接口的方法,返回加法结果,被装饰的部分
func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}
