package _1_facade

import "fmt"

//// * facade pattern

// NewAPI return API instance
func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type API interface {
	Test() string // 模拟一个需要使用到多个模块方法的复杂接口
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (api *apiImpl) Test() string {
	// 模拟一个复杂的逻辑, 需要调用多个模块的方法
	// 在这里,这个Test需要调用a模块的TestA和b模块的TestB
	aRet := api.a.TestA()
	bRet := api.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

//// ** a module

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type AModuleAPI interface {
	TestA() string
}
type aModuleImpl struct{}

func (api *aModuleImpl) TestA() string {
	return "A module running"
}

//// ** b module

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

type BModuleAPI interface {
	TestB() string
}
type bModuleImpl struct{}

func (api *bModuleImpl) TestB() string {
	return "B module running"
}
