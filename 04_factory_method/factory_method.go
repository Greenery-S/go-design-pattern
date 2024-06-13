package _4_factory_method

import "fmt"

//// * 工厂对外暴露的接口

// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
	ToString() string
}

// OperatorFactory 是工厂接口(出口)
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

//// * 工厂的加法实现

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (o PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
		symbol:       "+",
	}
}

type PlusOperator struct {
	*OperatorBase
	symbol string
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

func (o PlusOperator) ToString() string {
	return fmt.Sprintf("%d %s %d", o.a, o.symbol, o.b)
}

//// * 工厂的减法实现

// MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (o MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
		symbol:       "-",
	}
}

type MinusOperator struct {
	*OperatorBase
	symbol string
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}
func (o MinusOperator) ToString() string {
	return fmt.Sprintf("%d %s %d", o.a, o.symbol, o.b)
}
