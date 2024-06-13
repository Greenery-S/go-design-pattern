package visitor

import "testing"

func TestCustomerCol_Accept(t *testing.T) {
	// 首先新建两种顾客: 个人顾客和企业顾客
	bob := NewIndividualCustomer("bob")
	bytedance := NewEnterpriseCustomer("bytedance")

	// 通过customer的accept方法,接受visitor, 来拓展功能
	bob.Accept(&ServiceRequestVisitor{})
	bob.Accept(&AnalysisVisitor{}) // 个人顾客不支持分析功能,这里什么都不会做
	bytedance.Accept(&ServiceRequestVisitor{})
	bytedance.Accept(&AnalysisVisitor{}) // 企业顾客支持分析功能
	
	// 使用customerCol来批量调用accept方法
	customerCol := CustomerCol{}

	// 注册所有顾客
	customerCol.Add(bob)
	customerCol.Add(bytedance)

	// 批量调用accept方法,触发服务
	customerCol.Accept(&ServiceRequestVisitor{})
	customerCol.Accept(&AnalysisVisitor{})
}
