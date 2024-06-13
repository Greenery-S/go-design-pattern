# 外观模式

`API` 为 `facade` 模块的外观接口，大部分代码使用此接口简化对 `facade` 类的访问。

`facade` 模块同时暴露了 `a` 和 `b` 两个 `Module` 的 `NewXXX` 和 `interface`，其它代码如果需要使用细节功能时可以直接调用。

facade是一个上帝对象,它知道所有的子系统的功能,并且可以根据客户端的需求调用子系统的功能,并且对子系统的功能进行组合,以满足客户端的需求.

## 使用场景

facade类可以包含实际上很多的类型实例,使用这些实例和方法实现很复杂的功能,但是对于客户端来说,只需要调用facade类的方法就可以实现功能,而不需要知道facade类内部的实现细节.

比如说,我有一个walletFacade, 它包含了account,security,transaction等实例,客户端只需要调用walletFacade的方法就可以实现转账,查询余额等功能,而不需要知道account,security,transaction等实例的实现细节.

又比如,我有一个client,他需要httpclient,需要jwtGetter,需要logger...我可以把他们都封装到一个facade类中,然后client只需要调用facade类的方法就可以实现http请求+获取jwt+记录日志等功能.