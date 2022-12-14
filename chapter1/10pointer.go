package chapter1

import (
	"fmt"
)

// 在Go语言中的类型有：int、float、bool、string、array、slice、map、struct、func、interface 等
// 它们都有对应的指针类型，比如：*int、*float 等
// 从结构上可以看出，指针是成对存在的，一个指针声明了，必须要先指向，才可以使用！！！

// 对变量进行取地址（&）操作，可以获得这个变量的指针变量
// 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值

// 指针有什么用？
// 首先，指针可以操作所有的任何的类型，包括 go 语言自带的默认类型，它的作用就是让程序更加轻量高性能，
// 对结构体或引用类型的操作就非常合适，因为它们大型且复杂

// 一个指针变量指向了一个值的内存地址，同样的，在使用指针前你需要声明指针
// 等于 *int
var ppt1 *int
var ppt2 *float32
var ppt3 *string

// 从原理上讲，指针就是存储内存地址的16进制值
// 声明后，指针类型的变量初始值为 16 进制 0，且对应的内存空间不存在

// 重点：声明的指针，先指向对应的类型，才可以使用！！！

func Example10() {

	fmt.Println("Example10:")

	// 只声明的指针类型，其值为 nil
	println(ppt1) // 0x0 十六进制 0
	fmt.Printf("空指针的内容是：%x \n", ppt1)
	fmt.Printf("空指针的内容是：%x \n", ppt2)
	fmt.Printf("空指针的内容是：%x \n", ppt3)       // 只声明的空指针，默认值都是 0x0 ！！！
	fmt.Println("空指针是否等于nil:", ppt1 == nil) // 所以这也就是说，内存地址为 0x0 的位置其实并不存在，因此等于 nil

	// 赋值，必须左右类型一致
	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("a 变量的地址是: %x \n", &a)
	fmt.Printf("ip 变量储存的指针地址: %x \n", ip)
	fmt.Printf("*ip 变量的值: %d \n", *ip)

	// 引用和解引用
	x := 9
	y := 10
	//fmt.Printf("%T", &x) // 指针类型 *int

	*&x = 25 // 引用和解引用，x还是本身，所以这里重新对 x 赋值了

	//z := &y // 简写
	var z *int = &y // z为引用类型，才能接受 &y 引用
	*z = 100        // 函数传引用，然后函数体内部 *z 操作，对原实参产生影响，*z 表示的就是原来最初的那个 y

	println(x) // 25
	println(z) // 0xc000109f40
	println(y) // 100

	// 指针数组
	// 正确理解指针数组：就是原数组的每个元素的指针值组成的一个新数组的集合
	ax := [3]int{10, 100, 200}
	var axptr [3]*int // 定义一个整型指针数组（对应的是整型数组）

	for i := 0; i < 3; i++ {
		axptr[i] = &ax[i]
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("a[%d] = %d\n", i, *axptr[i])
	}

	// 指向指针的指针
	var ptr1 *int
	var pptr2 **int
	ppvalue := 25
	ptr1 = &ppvalue
	pptr2 = &ptr1
	println("指向指针的指针地址：", pptr2)

	//var abc **int
	//abc = &&y

	// 函数参数的引用
	yy := 100
	jj := 200
	swapNum(&yy, &jj)
	println("两个值已经交换了：", yy, jj)

	// 声明、指向、使用
	//var i1 *int      // 声明
	//*i1 = 10         // 没有指向的指针直接使用，这里没有报错是，但用法错误
	//fmt.Println(*i1) // panic: runtime error: invalid memory address or nil pointer dereference
	// 正确写法：
	var i2 *int   // 声明
	i2 = new(int) // 指向（new(Type) *Type，返回对应类型的指针）
	*i2 = 10      // 使用 // int是值类型，可以直接使用！
	fmt.Println(*i2)
}

func swapNum(x, y *int) {
	var temp int
	temp = *x
	*x = *y // 千万要注意这里：*x 接受的，还是需要解引用，才能修改原地址空间的值
	*y = temp

	// 更好的一个写法： *x, *y = *y, *x
	// 不需要返回
}
