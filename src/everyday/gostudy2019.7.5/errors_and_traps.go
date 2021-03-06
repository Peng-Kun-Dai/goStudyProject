package main

//永远不要使用形如 var p*a 声明变量，这会混淆指针声明和乘法运算（参考4.9小节）
//永远不要在for循环自身中改变计数器变量（参考5.4小节）
//永远不要在for-range循环中使用一个值去改变自身的值（参考5.4.4小节）
//永远不要将goto和前置标签一起使用（参考5.6小节）
//永远不要忘记在函数名（参考第6章）后加括号()，尤其调用一个对象的方法或者使用匿名函数启动一个协程时
//永远不要使用new()一个map，一直使用make（参考第8章）
//当为一个类型定义一个String()方法时，不要使用fmt.Print或者类似的代码（参考10.7小节）
//永远不要忘记当终止缓存写入时，使用Flush函数（参考12.2.3小节）
//永远不要忽略错误提示，忽略错误会导致程序奔溃（参考13.1小节）
//不要使用全局变量或者共享内存，这会使并发执行的代码变得不安全（参考14.1小节）
//println函数仅仅是用于调试的目的

//最佳实践：对比以下使用方式：
//使用正确的方式初始化一个元素是切片的映射，例如map[type]slice（参考8.1.3小节）
//一直使用逗号，ok或者checked形式作为类型断言（参考11.3小节）
//使用一个工厂函数创建并初始化自己定义类型（参考10.2小节-18.4小节）
//仅当一个结构体的方法想改变结构体时，使用结构体指针作为方法的接受者，否则使用一个结构体值类型10.6.3小节
