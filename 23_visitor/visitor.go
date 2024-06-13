package visitor

import "fmt"

//// * 访问者模式

// Customer 被访问的对象,比如基础库对象, 有一个Accept方法,接受一个Visitor对象
type Customer interface {
	Accept(Visitor)
}

// Visitor 访问者接口
// 通过Visitor接口,可以访问不同的对象
// 通过visitor来给这些基础库对象添加新的功能, 一般命名为 功能/服务+Visitor
type Visitor interface {
	Visit(Customer)
}

// CustomerCol 顾客集合,批量化发起顾客
type CustomerCol struct {
	customers []Customer
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

// EnterpriseCustomer 企业顾客
type EnterpriseCustomer struct {
	name string
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name: name,
	}
}

func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

// IndividualCustomer 个人顾客
type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}

func (c *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

// ServiceRequestVisitor 服务请求访问者, 为基础类customer添加新的功能
// 服务功能为企业和个人顾客提供服务
type ServiceRequestVisitor struct{}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n", c.name)
	}
}

// AnalysisVisitor only for enterprise
// 企业顾客分析访问者, 为基础类customer添加新的功能
// 服务功能为企业顾客提供分析
type AnalysisVisitor struct{}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s\n", c.name)
	}
}
