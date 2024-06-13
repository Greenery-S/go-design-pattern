package main

import (
	. "design-pattern/05_abstract_factory"
	"fmt"
	"os"
)

// 小插曲: pipe可以将fmt.print的输出,从stdout重定向到pipe中,这样就可以检查print的是否符合预期.
// TODO: golang i/o 需要进一步学习...

func main() {
	// 创建一个管道
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println("Error creating pipe:", err)
		return
	}

	// 保存当前的os.Stdout，并将os.Stdout重定向到管道
	oldStdout := os.Stdout
	os.Stdout = w

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

	// 在一个goroutine中写入数据到标准输出
	go func() {
		defer func() {
			// 恢复os.Stdout
			os.Stdout = oldStdout
			w.Close()
		}()
		//fmt.Println("Hello, stdout!")
		mainDAO.SaveOrderMain()
		detailDAO.SaveOrderDetail()
	}()

	// 在主goroutine中从管道读取数据
	buf := make([]byte, 128)
	n, err := r.Read(buf)
	if err != nil {
		fmt.Println("Error reading from pipe:", err)
		return
	}

	fmt.Println("Read from pipe:", string(buf[:n]))
}
