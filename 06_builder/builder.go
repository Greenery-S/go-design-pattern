package builder

//// * 生成器模式: 需要主管和生成器两个角色，主管负责调用生成器的方法，生成器负责实现具体的构建逻辑

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Builder 是生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	builder Builder
}

// Construct Product
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

//// * 生成产品1, 产品1有3个部分，每个部分都是字符串

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}

//// * 生成产品2, 产品2有3个部分，每个部分都是整数

type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}

func (b *Builder2) GetResult() int {
	return b.result
}
