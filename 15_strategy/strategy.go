package strategy

import "fmt"

//// * 策略模式

// Payment 是存算分离的.
// - PaymentContext 是存储支付信息的.
// - PaymentStrategy 是支付策略的接口.
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

// NewPayment 用于创建一个 Payment 实例.
func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}

// Pay 上下文+策略结合,实现支付.
func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

//// * Cash

// Cash 是支付策略之一.
type Cash struct{}

// Pay 实现了 PaymentStrategy.Pay 方法.
func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

//// * Bank

// Bank 是支付策略之一.
type Bank struct{}

// Pay 实现了 PaymentStrategy.Pay 方法.
func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s\n", ctx.Money, ctx.Name, ctx.CardID)
}
