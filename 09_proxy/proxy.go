package proxy

// Subject 是一个接口，定义了 RealSubject 和 Proxy 的共同接口，这样就在任何使用 RealSubject 的地方都可以使用 Proxy。
// 这样就实现了透明代理: Proxy 假装自己是 RealSubject，但是实际上 Proxy 内部持有 RealSubject 的引用。
type Subject interface {
	Do() string
}

// RealSubject 是真实对象，是最终执行业务逻辑的对象。比如数据库,服务器等.
type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

// Proxy 是代理对象，持有 RealSubject 的引用，可以在调用 RealSubject 之前和之后执行一些操作。
type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。。
	res += "pre:"

	// 调用真实对象
	res += p.real.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等。。
	res += ":after"

	return res
}
