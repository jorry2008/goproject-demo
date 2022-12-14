package chapter5

// Go 里面所有的声明，都是统一的 关键字 命名 定义内容 这样的顺序，这种语法为『类型语法（type syntax）』
// 值的定义：『无法进一步求值的表达式（expression）』，例如 4 + 3 * 2 / 1 的值是 10
// 常量和变量，则可以理解为值的容器
// go 语言中，万物皆类型
// Go 没有 类（class）和 继承（inheritance），而是通过 结构体 和 组合（composition）实现面向对象。
// 结构体的零值是所有字段都为对应零值的结构体

/*
接口（interface）：接口是一系列行为（方法签名，方法是一种特殊的函数）的集合。跟其它语言不同，Go 的接口不需要显式声明实现（implementation），一个类型只要实现了接口的所有方法，它就隐式地满足接口。
是否满足接口可以在编译期静态检查，所以是类型安全的。
Go 实现了类型安全的鸭子类型（duck typing） 。
这种设计是 Go 的组合式面向对象的重要组成部分。
*/
