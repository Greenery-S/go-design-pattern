package game

import "fmt"

//// * 备忘录模式

// Memento 定义一个备忘录接口, 用于保存状态
type Memento interface{}

// Game 定义一个游戏, 用于保存当前的状态, 并且可以保存和恢复状态
// hp: 生命值, mp: 魔法值都是私有的, 只有通过Play方法修改. 没有提供直接修改的方法,所以不能直接恢复.
type Game struct {
	hp, mp int
}

// gameMemento 用于保存当前的状态,和Game在同一个包中,所以可以访问Game的私有属性
// 所有属性都和Game对应
type gameMemento struct {
	hp, mp int
}

func (g *Game) Play(mpDelta, hpDelta int) {
	g.mp += mpDelta
	g.hp += hpDelta
}

// Save 用于保存当前的状态, 并返回一个Memento接口
// 由于调用Save在Game包外发生,所以不能断言为gameMemento,也就不能访问私有属性
func (g *Game) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

// Load 用于恢复状态, 传入一个Memento接口,并恢复状态
// 由于调用Load,将Memento传回Game包,所以可以断言为gameMemento,也就可以访问私有属性
func (g *Game) Load(m Memento) {
	gm := m.(*gameMemento)
	g.mp = gm.mp
	g.hp = gm.hp
}

func (g *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", g.hp, g.mp)
}
