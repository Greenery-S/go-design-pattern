package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

//// * 解释者模式
// 解释者模式是一种行为设计模式， 允许将语言解释器与表达式结合使用， 以解释一个简单语言。
// 解释者模式通常用于解释编程语言和脚本语言， 但也可用于任何领域， 如解释数学表达式、配置文件等。
// 解释者模式的关键是定义语言的语法， 并用解释器解释语言中的表达式。

// Node 解释器节点接口
type Node interface {
	Interpret() int
}

// ValNode 数值节点
type ValNode struct {
	val int
}

func (n *ValNode) Interpret() int {
	return n.val
}

// AddNode 加法节点
type AddNode struct {
	left, right Node
}

func (n *AddNode) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

// MinNode 减法节点
type MinNode struct {
	left, right Node
}

func (n *MinNode) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

//// * 解释器模式实现

// Parser 解释器
type Parser struct {
	exp   []string // 表达式
	index int      // 当前解析位置
	prev  Node     // 上一个解析完毕的节点
	// 构建当前节点时，需要用到上一个解析完毕的
}

// Parse 解析表达式
func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ") // 按空格分割表达式

	for { // 循环解析表达式
		fmt.Println(p.stateCheck())
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] { // 根据当前字符类型，创建不同的当前节点, 并赋值给prev
		case "+":
			p.prev = p.newAddNode()
		case "-":
			p.prev = p.newMinNode()
		default:
			p.prev = p.newValNode()
		}
	}
}

// 展示解析状态, 用于调试, 打印 exp | index | prev 的值
func (p *Parser) stateCheck() string {
	currentExp := p.exp[:p.index]
	if p.prev == nil {
		return "Empty, " + strings.Join(currentExp, " ")
	}
	switch p.prev.(type) {
	case *AddNode:
		return fmt.Sprintf("AddNode %d, ", p.prev.Interpret()) + strings.Join(currentExp, " ")
	case *MinNode:
		return fmt.Sprintf("MinNode %d, ", p.prev.Interpret()) + strings.Join(currentExp, " ")
	case *ValNode:
		return fmt.Sprintf("ValNode %d, ", p.prev.Interpret()) + strings.Join(currentExp, " ")
	default:
		return "Unknown"
	}
}

func (p *Parser) newAddNode() Node {
	p.index++
	return &AddNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newMinNode() Node {
	p.index++
	return &MinNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newValNode() Node {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValNode{
		val: v,
	}
}

func (p *Parser) Result() Node {
	return p.prev
}
