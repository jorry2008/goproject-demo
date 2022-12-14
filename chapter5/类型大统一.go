package chapter5

import "fmt"

// 理解不了传值、传引用，理解不了引用、指针，理解不了各种复合体...本质原因是因为对类型理解不到位！
// go 体系中，只有值类型和引用类型两大值，别想了，就这两大类，它们的行为差别非常大，且 string 极为特殊。

// 参考：https://www.cnblogs.com/xbblogs/p/11102970.html

/*
go语言中的值类型：
	int、float、bool、array、sturct 等【牛逼在于，struct 自定义结构体，也是值类型！！！】
	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	声明一个”值类型“变量时，编译器会在栈中分配一个空间，空间里存储的就是该变量的值。（值本身）

go语言中的引用类型：
	slice、map、channel、interface、func、string 等【不敢相信吧，interface，func，string 居然是引用类型】
	声明一个引用类型的变量，编译器会把实例的内存分配在堆上，引用本身存储在了栈上。
	所谓引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

	注意：nil 可以赋值给引用类型（除 string 外）、error 类型和指针类型
	注意：string 和其他语言一样，是引用类型，string 的底层实现 struct String { byte* str; intgo len; }; 但是因为string不允许修改，每次操作 string 只能生成新的对象，所以在看起来使用时像值类型。【重点】

	重点来了：
	不论是值类型，还是引用类型，它们从函数参数上传递的时候，函数所做的一件事情都是将参数的值重新复制一份到栈中分配一个空间用于存放这些参数的值。
	不同的是，值类型本身存储的就是数值（传递效果就是全新的一个变量与之前变量无关），而引用类型存储的是数值所在堆的地址（传递的效果就是对函数外部传递进来的引用关联的变量都将有影响）。
	因此，对于函数参数传递而言，动作只有一种：就是将值复制过来而已。。。。。

	// 所以，函数的传递背锅了，具体是什么传递，跟函数本身没有任何关系！

go语言中的指针类型：
	一个指针变量指向了一个值的内存地址，指针变量存储的内容就是所指向值的内存地址值。
	当一个指针被定义后没有分配到任何变量时，它的值为 nil。nil 指针也称为空指针
	其实引用类型可以看作对指针的封装，但！！！引用类型不是指针类型！它们是两个不同的东西。

go语言中的字符串类型：
	我们来重新认识一下字符串，这东西真的非常特殊，本质它应该是个
	// 待定
*/

func ASaabb() {

	aaa := new(int) // 值类型
	fmt.Println(aaa)

	ccc := new(struct{}) // 值类型
	fmt.Println(ccc)

	eee := new([]int) // 引用类型
	fmt.Println(eee)

	bbb := new(string) // 特殊的引用类型
	fmt.Println(bbb)

}
