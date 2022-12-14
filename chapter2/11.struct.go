package chapter2

import (
	"fmt"
)

// 注意：结构体只是值类型，它与数组的行为几乎是一样的。

// 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合
// 结构体是开发人员自定义类型集合，是一种新类型的创建方式，跟数组一样，数组是同类型数据的集合，而结构体则完全开放了类型限制（像极了其它语言的类的定义）
// 结构体中的属性和关联方法，叫成员属性和成员方法（即，所属结构体的变量才叫属性，所属结构体的函数才叫方法）
// 组合结构体（组合接口），实现了继承特性

// 匿名结构体（啥用？）
var (
	alice Circle
	bob   Circle
)

// 初始化结构体变量

type Circle struct {
	radius float64
}

// 该 method 属于 Circle 类型对象中的方法
// 在结构体中“定义方法”，取决于函数是否接受了指定的结构体：
// 一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针，所有给定类型的方法属于该类型的方法集

func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func Example11() {
	fmt.Println("Example11:")
	// 结构体的”实例化“直接声明即可

	var varint int
	_ = varint

	// 同上，Circle 也是一个类型，结构体也是一样，只声明也行
	var c1 Circle // 声明一个 Circle 类型即可！！！【在 Go 语言中，所有的语法组织，理解都趋于扁平化，Circle 类型同 int 类型一样的用法，其它语言可能要 new Circle。。。。】
	c1.radius = 20
	fmt.Println("c1的面积：", c1.getArea())

	c2 := Circle{25} // 同 c2 := Circle{radius: 25}
	fmt.Println("c2的面积：", c2.getArea())

	c3 := &Circle{radius: 30} // 同 c3 := &Circle{30}
	fmt.Println("c3的面积：", c3.getArea())

	// 注意，结构体的返回值
	fmt.Printf("%T, %T, %T, %T", varint, c1, &c2, &c3) // int, chapter2.Circle, *chapter2.Circle, **chapter2.Circle

	// 注意：初始化属性可以有以下方式
	// 一、按顺序赋对应类型的值即可
	// 二、可以使用 key => value 格式，这种方式可以不按照顺序
	// 三、可以初始化部分值，需要忽略的属性，可以不填充
	type Books struct {
		title   string
		author  string
		subject string
		book_id int
	}
	b1 := Books{
		"go开发教程",
		"jorry",
		"文章标题",
		25, // 注意，当结构体自身有两个以上属性时，最后一个逗号必须带上
	}
	b2 := Books{
		author:  "哈",
		title:   "go提升", // 可以不按照顺序
		subject: "书名",
		book_id: 222,
	}
	b3 := Books{
		title: "xxxx", // 可以不用填充全部属性
	}
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3) // 重点：没有填充的属性，并不代表没有，而是全部填充了默认值

	// 结构体简单的使用
	type Address struct {
		Province    string
		City        string
		ZipCode     int
		PhoneNumber string
	}
	addr := Address{ // 跟数组的使用完全一样
		"四川", // 每个类型之间，使用逗号分隔
		"成都",
		610000,
		"18888888888",
	}
	println()
	fmt.Printf("%T \n", addr)
	fmt.Println("打印一下类型变量", addr) // {四川 成都 610000 0} // 这种格式，想起了 js 的对象

	// 结构体的引用初始化
	type People struct {
		name  string
		child *People // 结构体不能包含自身，比如 People 中的字段不能是 People 类型，但却可能是 *People
	}
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
				child: &People{
					name: "我儿子",
				},
			},
		},
	}
	fmt.Printf("%T \n", relation)
	println(relation.name)       // 引用或者变量类型，这个返回值是一样的！！！
	println(relation.child.name) // 引用或者变量类型，这个返回值是一样的！！！  这里需要深入研究，为什么会这样？？？

	// 匿名结构体(普通类型)
	// 匿名结构体的类型名是结构体包含字段成员的详细描述，匿名结构体在使用时需要重新定义，造成大量重复的代码，因此开发中较少使用
	msg := struct {
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	printMsgType(msg)

	// 引用类型
	msgptr := &struct {
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	printMsgTypePtr(msgptr)

	// 组合结构体，实现了继承
	c := cat{animal: animal{name: "animal", age: 2}, name: "cat"}
	fmt.Println(c)
	fmt.Println("name", c.name)
	fmt.Println("name", c.animal.name)
	fmt.Println("age", c.age)
	fmt.Println("age", c.animal.age)

	// 从结构体的定义出发，本质它是一个定义出来的类型，使用这个类型时，与普通类型是完全一样的，传值特性也是一样的

	// 在 go 语言中，将父结构嵌入到子结构，就是结构体组合操作，这个叫继承
	// 想从子结构转型为父级，直接： son.parent 即可

	// 结构体的值引用
	m1 := cat{}    // 值类型
	m2 := new(cat) // 指针结构体
	Change(m1, m2)
	fmt.Println(m1, m2)
}

func Change(m1 cat, m2 *cat) {
	m1.name = "小明" // 不变
	m2.name = "小红" // 影响
}

// 打印消息类型, 传入匿名结构体（普通类型）
func printMsgType(msg struct {
	id   int
	data string
}) {
	fmt.Printf("匿名结构体（跟匿名函数很像）：%T \n", msg) // 结构体的整个类型就是这样的：struct { id int; data string }
}

// 引用类型
func printMsgTypePtr(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("匿名结构体，引用传递（跟匿名函数很像）：%T \n", msg) // 结构体的整个类型就是这样的：*struct { id int; data string }
}

type animal struct {
	name string
	age  int
}

type cat struct {
	animal
	name string
}

/*
关于结构体和类，面向对象的对比理解
Go 没有面向对象，而我们知道常见的 Java、C++ 等语言中，实现类的方法做法都是编译器隐式的给函数加一个 this 指针
而在 Go 里，这个 this 指针需要明确的申明出来，其实和其它 OO 语言并没有很大的区别
C++ 是这样写的：
class Circle {
  private:
    float radius;
  public:
    float getArea() {
       return 3.14 * radius * radius;
    }
}
// 其中 getArea 经过编译器处理大致变为
float getArea(Circle *const c) { // 可以看到，不管是 C++ 和 Go 本质是一样的，形式可能不太一样，而且更多的在于他们的编译器到底为他们做了多少（甚至可以说，一个语言的使用体验，取决于如何定义编译器）
  // ...
}

Go 代码是这样写的：
func (c Circle) getArea() float64 {
    // c.radius 即为 Circle 类型对象中的属性（在 Go 中，结构体的属性是需要显示指定的）
    return 3.14 * c.radius * c.radius // 这里相当于 this.radius，但又由于 Go 的方法是定义在结构体之外的，所以使用 this 显得不易理解。。。。。
}
*/

// 结构体是类型组合体 + 功能体？
// 如此看的话，它跟类没有什么两样了....而结构体的定义显得更加抽象

//reflect.TypeOf()
//unsafe.Sizeof()
