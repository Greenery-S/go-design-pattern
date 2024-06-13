package composite

import "fmt"

//// * 组合模式

// Component 组件接口, 定义了组件的基本行为, 涵盖了树节点和叶子结点.
type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

const (
	LeafNode = iota
	CompositeNode
)

func NewComponent(kind int, name string, val any) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf(val)
	case CompositeNode:
		c = NewComposite()
	}

	c.SetName(name)
	return c
}

// component 是Component接口的实现, 是树节点和叶子结点的共同基类.
// 树节点和叶子结点通过内嵌来继承这个基类, 可以重写基类的方法.
type component struct {
	Component // 内嵌接口来实现继承,可以不实现部分方法.
	parent    Component
	name      string
}

// 实现必要的基类方法
// 由于内嵌了Component接口, 所以可以不实现AddChild和Print方法.
//func (c *component) AddChild(Component) {}
//func (c *component) Print(string) {}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

//// * 叶子结点

type Leaf struct {
	component // 内嵌基类, 继承基类的方法.
	val       any
}

func NewLeaf(val any) *Leaf {
	return &Leaf{val: val}
}

func (c *Leaf) Print(pre string) {
	fmt.Printf("%s-%s:%v\n", pre, c.Name(), c.val)
}

//// * 树节点

type Composite struct {
	component             // 内嵌基类, 继承基类的方法.
	childs    []Component // 子节点
}

func NewComposite() *Composite {
	return &Composite{
		childs: make([]Component, 0),
	}
}

func (c *Composite) AddChild(child Component) {
	child.SetParent(c)                 // 子->父
	c.childs = append(c.childs, child) // 父->子
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, c.Name())
	for _, child := range c.childs {
		child.Print(pre + "  ")
	}
}
