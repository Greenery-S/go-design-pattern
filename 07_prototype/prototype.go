package prototype

//// * 原型模式: 需要一个原型管理器, 用于存储原型对象. 通过原型管理器获取原型对象的克隆对象. 还有一个原型对象需要实现的接口, 用于克隆自己.

// Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

// PrototypeManager 是原型管理器, 用于存储原型对象, 通过原型管理器获取原型对象的克隆对象
type PrototypeManager struct {
	prototypes map[string]Cloneable
}

// NewPrototypeManager 创建一个原型管理器
func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

// Get 通过原型管理器获取原型对象的克隆对象
func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name].Clone()
}

// Set 设置原型对象
func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

//// * type1 和 type2 是原型对象, 实现了 Cloneable 接口, 用于克隆自己.

type Type1 struct {
	text string
}

// Clone 实现 Cloneable 接口, 用于克隆自己
func (t *Type1) Clone() Cloneable {
	tc := Type1{text: t.text}
	return &tc
}

type Type2 struct {
	val int
}

// Clone 实现 Cloneable 接口, 用于克隆自己
func (t Type2) Clone() Cloneable {
	tc := Type2{val: t.val}
	return tc
}
