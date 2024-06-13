package main

import "design-pattern/17_memento/game"

func main() {
	// 初始化hp=100,mp=50
	gm := game.Game{}
	gm.Play(+100, +50)
	gm.Status()
	// gm.XXX // 无法访问私有属性

	// 保存当前状态
	checkpoint := gm.Save()
	//checkpoint.(*game.gameMemento)
	//Cannot use the unexported type 'gameMemento' in the current package, 私有接口无法断言

	// 战斗
	gm.Play(-50, -30)
	gm.Play(-3, -5)
	gm.Status()

	// 恢复状态
	gm.Load(checkpoint)
	gm.Status()
}
