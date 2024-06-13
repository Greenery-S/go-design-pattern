package abstractfactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	var (
		// 工厂接口
		factory DAOFactory

		// 一"系列产品"的接口
		mainDAO   OrderMainDAO
		detailDAO OrderDetailDAO
	)

	// RDB工厂
	factory = &RDBDAOFactory{}
	mainDAO = factory.CreateOrderMainDAO()
	detailDAO = factory.CreateOrderDetailDAO()
	assert.Equal(t, "rdb main save", mainDAO.SaveOrderMain())
	assert.Equal(t, "rdb detail save", detailDAO.SaveOrderDetail())

	// XML工厂
	factory = &XMLDAOFactory{}
	mainDAO = factory.CreateOrderMainDAO()
	detailDAO = factory.CreateOrderDetailDAO()
	assert.Equal(t, "xml main save", mainDAO.SaveOrderMain())
	assert.Equal(t, "xml detail save", detailDAO.SaveOrderDetail())
}
