package iterator

//// * 迭代器模式

// Aggregate ,迭代元素的核心信息,记录像是定义迭代的的范围,步长等信息. Iterator()获取初始化迭代器.
type Aggregate interface {
	Iterator() Iterator
}

// Iterator ,迭代器接口,包含迭代核心信息+迭代值过程中信息
type Iterator interface {
	First()       // 重置迭代器
	IsDone() bool // 是否迭代完成
	Next() any    // 下一个元素
}

//// * 具体实现

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

// Numbers ,具体的迭代核心信息
type Numbers struct {
	start, end int
}

// Iterator ,获取初始化迭代器
func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		next:    n.start,
	}
}

// NumbersIterator ,迭代器实现, 不需要New方法,因为Numbers已经实现了Iterator()方法
type NumbersIterator struct {
	numbers *Numbers
	next    int
}

// First ,重置迭代器
func (i *NumbersIterator) First() {
	i.next = i.numbers.start
}

// IsDone ,是否迭代完成
func (i *NumbersIterator) IsDone() bool {
	return i.next > i.numbers.end
}

// Next ,下一个元素
func (i *NumbersIterator) Next() any {
	if !i.IsDone() {
		next := i.next
		i.next++
		return next
	}
	return nil
}
