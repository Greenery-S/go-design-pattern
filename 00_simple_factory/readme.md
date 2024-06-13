# 简单工厂模式

`go` 语言没有构造函数一说，所以一般会定义 `NewXXX` 函数来初始化相关类。
`NewXXX` 函数返回接口时就是简单工厂模式，也就是说 `Golang` 的一般推荐做法就是简单工厂。

在这个 `simplefactory` 包中只有API 接口和 `NewAPI` 函数为包外可见，封装了实现细节。

## 实际场景

通常在new client的时候,会遇到需要实例化略有不同的client,例如我需要new一个cn的client和i18n的client.

假如他们的鉴权略有不同时,相比写两个不同的NewXXXCN,NewXXXI18N函数,倒不如使用指明type的简单工厂模式,这时候就可以使用简单工厂模式.
