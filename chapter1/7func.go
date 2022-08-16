package chapter1

import (
	"fmt"
	"math"
)

// Fun1 就是一个全局匿名函数
// 定义一个全局函数，如果想被全局调用首字母大写，如果只能包内部调用首字母小写
var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

// init()函数可以不需要，如果有 init() 函数则会先执行该函数，init() 函数不可以有参数，也不可以有返回
func init() {
	fmt.Println("初始化函数")

	// 连声明匿名函数，边调用
	func() {
		fmt.Println("声明函数，并立马调用")
	}() // 在函数中 () 就是调用符号！！！
}

// Example7 局部声明或者初始中 := 操作很常用，也很方便
// 符合大写公有、小写私有规则的还包括：常量、变量、类型、函数名、结构字段
func Example7() {
	println("Example7:")

	// 空白标识符，是一个只写变量，你不能得到它的值，正因为这种特性，空白标识符可用于接受多余的值并执行了初始化，但又可以不在当前作用域下使用，从而避免了错误（局部声明变量不使用报错的问题！！）
	// 空白标识符用来人为抛弃数据，又不至于触发语法错误（语法规则是：在函数体中声明的所有变量必须要使用，函数返回的所有值必须要接收，否则都将报错）
	_, numb, strs := numbers() //只获取函数返回值的后两个
	println(numb, strs)

	// 调用内部私有函数，直接调用即可
	other()

	c, d := swap(20, 30) // 调用函数时，20和30是实际参数
	println("数值交换：", c, d)

	s1, s2 := 100, 200
	quoteSwap(&s1, &s2) // 这里必须是将引用传递过去（即两个变量的地址值）
	println("函数引用交换数值后：", s1, s2)

	// 全局匿名函数传递
	println("使用变量，调用全局匿名函数，结果：", Fun1(25, 25))
	// 局部匿名函数传递
	getSquareRoot := func(x float64) float64 { // 平方根
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	// 函数回调
	fmt.Println("回调函数：", funcback(1024, getSquareRoot))

	// 方法
	// 请查看：struct

	// 闭包
	// 请查看：closure

	// 函数类型（golang中，万物皆类型）
	var nf FuncType
	nf = StrFunc // 名称、参数、返回值，相同，即类型相同，同类型可赋值（函数类型，同 map 一样是引用类型，函数名就是函数体的内存地址入口）
	str := "hello world"
	nf(str) // uint8 hello world
}

type FuncType func(string) // 定义了一个类型，基于函数类型创建

func StrFunc(s string) {
	fmt.Printf("%T ", s[1])
	fmt.Println(s)
}

// 正常情况下，函数声明告诉了编译器函数的名称，返回类型，和参数（一个函数类型，包括：函数名、参数列表、返回列表 三部分）
func swap(a, b int) (int, int) { // 定义函数参数时，a 和 b 叫形式参数，形式参数如同函数体中的局部变量（函数中，参数列表和返回类型列表都是可选的）
	return b, a
}

// 传引用操作，和 c 语言操作完全一样
// 调用函数时，实参做了取地址操作，传入过来的是地址值（也就是指针），所以这里的形参同样也需要使用对应数据类型的指针类型来接收
func quoteSwap(x *int, y *int) { // x, y 在这里就是指针类型，存储的是地址值
	var temp int
	temp = *x // 想要取值，就得解引用操作
	*x = *y   // 重点！接受赋值用的 x 也必须是这种形式 *x 表示取值状态，也表示同类型，否则无法赋值【*x被赋值，即表示x的引用被赋值】
	*y = temp // 到这一步，程序将两个内存中的值交换了
}

// 一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}

// 整个 Go 程序只分内外两种访问类型，大写开头的对外部包（包内和包外）可访问（公有），小写开头的对内部包（所有同包下的文件都可访问）访问（私有）
// 所以我们会发现，在程序中，大写开头的属性和函数全是外部外引入进来的，所有小写开头的属性和函数都是当前包的（注意：是以包为分界线）
func other() {
	fmt.Println("当前包的私有函数")
}

func Other() {
	fmt.Println("其它包的公共函数")
}

func funcback(num float64, f func(x float64) float64) float64 { // 一个函数的定义包括：函数名、参数列表、返回参数列表 三者组合在一起才叫一种函数类型！！！
	return f(num)
}

// 小结：
// 1.函数类型，是引用类型，函数名是变量，存储的就是函数体的内存地址入口。
// 2.go函数的参数，就是值传递，至于传的是值类型还是引用类型，是由参数自身决定。（并不是函数的某种机制）
// 3.函数名、参数列表、返回列表，三者相同，就认为是同一个函数类型。
