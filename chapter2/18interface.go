package chapter2

import "fmt"

// 接口是隐性的背后的王！
// go中的所有类型都实现了 interface{} 的接口（空接口），这意味着，所有的类型（如 string,int,int64 甚至是自定义的 struct 类型）都就此拥有了 interface{} 接口，这种做法和 java 中的 Object 类型类似
// 那么在一个数据通过 func funcName(interface{}) 的方式传进来的时候，也就意味着这个参数被自动的转为 interface{} 的类型。也即，这个函数支持所有类型
// 正因为这个规则，go语言中提供了断言的功能（后面再说）

// 整体看待：数据类型、结构体、接口、方法，三个元素
/*
结构体可以理解为，将其它数据类型(或其它结构体类型)自定义组合到一起，形成的一个新的数据类型；【开发者设计的入口】
一个新的类型自身包括有数据，方法是自动匹配进来的；
它们是否实现接口，同样是由方法自动匹配的，实现了接口，结构体就方便复用；
在这个过程中，方法即反向绑定了结构体，还实现了接口；
至此为止，go 已经拥有描述世界的能力；

在 go 语言中，同样是有对象的，对象的模板（类）在 go 语言中叫数据结构+接口[可选]+方法[可选]，

小结：整个机制，不得不说简直就是极简到家了。。。。不需要写多余的标识标，哪怕一个字母。


注意：
func sum(x,y int) int { // 函数
    return x + y
}

func (var[可选] Type) imp() string { // 方法【函数上添加一个结构体接收者，便是方法，而接收者即可以是】
    return "这是方法"
}

使用关键点：
1.所有接口的 方法名、参数、返回类型 都匹配了接口才会被实现（接口实现不是强制的，是根据类实现的方法来动态判定的）
2.如果方法有对结构体数据的写操作，方法的结构体参数必须是指针，否则写无效（不报错）
3.接口跟其它语言一样，接口为结构体类型提供的隐式转换
4.组合结构体（组合接口），实现了继承特性
5.接口组合，实现了什么呢？





Go 语言正是通过提供结构体组合和接口组合的机制，
让一个结构体/接口包含另一个结构体/接口类型的匿名成员，
这样就可以通过简单的点运算符 x.f 来访问匿名成员链中嵌套的 x.d.e.f 成员
同样的规则，内嵌匿名成员链的方法也会提升为外部类型的方法？？？
*/

var (
	cindy Phone
	danny Phone
)

// 这里声明了两个 『拥有两个方法的接口』的变量。接口变量的零值是 nil ，可以接受任何满足接口的类型的值。(批量声明接口类型变量)

type Phone interface {
	title()
	call()
	charge()
	electricQuantity() string
}

type NokiaPhone struct {
	title string
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("我是" + nokiaPhone.title + "，我能打电话!")
}
func (nokiaPhone NokiaPhone) charge() {
	fmt.Println("我是" + nokiaPhone.title + "，我能充电!")
}
func (NokiaPhone) electricQuantity() string {
	return "当前电量：80%"
}

type IPhone struct {
	title string
}

func (iPhone IPhone) call() {
	fmt.Println("我是" + iPhone.title + "，我能打电话!")
}

func Example18() {
	fmt.Println("Example:")

	// 使用 new() 函数创建一个结构体实例
	nokia := new(NokiaPhone)
	nokia.title = "瑞典诺基亚"
	nokia.call()
	nokia.charge()
	nokia.electricQuantity()

	ipone := new(IPhone)
	ipone.title = "爱疯"
	ipone.call()

	// 初始化一个结构体实例
	nokia1 := NokiaPhone{title: "诺诺诺"}
	nokia1.call()
	nokia1.charge()
	nokia1.electricQuantity()

	// 声明一个结构体实例
	var nokia2 NokiaPhone
	nokia2.title = "诺什么机"
	nokia2.call()

	// 接口的意义：继承
	vivo1 := Android1{brand: "Vivo2"}
	ix := IPhone1{"X12"}

	printCall(vivo1) // 同一个函数，即可以接受 Android1 结构体类型，也可以接受 IPhone1 结构体类型
	printCall(ix)    // 会隐式地将 vivo 和 ix 对象转换成 Phone1 类型

	// 结构、方法、接口 与指针
	// 重点：如果想要通过方法（包括接口方法，即所有方法）修改属性（属性就是结构体的数据），需要在传入指针的结构体才有效！
	// 所有，在使用上，所有结构体方法，都传入结构体的指针即可。

	var fruit1 fruit        // 声明接口
	fruit1 = &apple{"糖心苹果"} // 所有声明接口类型变量，只能接受实现了此接口的结构体的指针类型！！！
	println(fruit1.getName())

	ap := apple{"红富士"}
	fmt.Println(ap.getName())
	ap.setName("树顶红")
	fmt.Println(ap.getName())

	// 组合接口
	var mouse1 mouse // 非接口类型，使用了最终的实体类型
	mouse1.write("jorry1")
	println("写入的名称为：", mouse1.read())

	var rw1 rw
	rw1 = &mouse{"初始化名称"} // 注意注意：接口声明变量，必须接受同结构体的指针类型！！！
	rw1.write("jorry2")
	println("写入的名称是：", rw1.read())

	// 空接口类型的值包裹着各种非接口值的例子【重】
	var i interface{}
	i = []int{1, 2, 3}
	fmt.Println(i) // [1 2 3]
	i = map[string]int{"Go": 2012}
	fmt.Println(i) // map[Go:2012]
	i = true
	fmt.Println(i) // true
	i = 1
	fmt.Println(i) // 1
	i = "abc"
	fmt.Println(i) // abc

	// 将接口值i中包裹的值清除掉。
	i = nil
	fmt.Println(i) // <nil>
}

type Phone1 interface { // 可以跳转到 实现接口的结构体
	call() string // 可以跳转到 实现的方法
}

type Android1 struct {
	brand string
}

type IPhone1 struct {
	version string
}

func (android Android1) call() string {
	return "I am Android " + android.brand
}

func (iPhone IPhone1) call() string {
	return "I am iPhone " + iPhone.version
}

func printCall(p Phone1) { // 基于接口类型的参数呈现出的多态特性
	// 因为 Phone1 接口本身就规定了 call() 方法必须实现
	fmt.Println(p.call() + ", I can call you!")
}

type fruit interface {
	getName() string
	setName(name string)
}
type apple struct {
	name string
}

func (a apple) getName() string {
	return a.name
}
func (a *apple) setName(name string) { // 如果方法要修改结构体的数据，则匹配的结构体必须是其指针类型
	a.name = name
}

type reader interface {
	read() string
}

type writer interface {
	write(string) bool // 参数可以不需要名称标识符
}

type rw interface { // 组会接口
	reader
	writer
}

type mouse struct {
	name string // 小写，外部不能访问，封装
}

func (m mouse) read() string {
	return m.name
}

func (m *mouse) write(name string) bool { // 写操作，必须传入结构体指针
	m.name = name
	return true
}

// 值接受者 还是 指针接受者？
// 结构体是值类型，给结构体定义方法时，到底俩月和值接受者，还是指针接受者？具体取决于什么？
// https://morven.life/posts/golang-interface-and-composition/
