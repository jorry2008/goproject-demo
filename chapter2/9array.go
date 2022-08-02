package chapter2

import "fmt"

// Go 语言中有数组，只是没有 php 那么灵活，相比 Java 其实是差不多的
// 数组是具有"相同类型"的一组已编号且长度固定的数据项序列（连续内存空间），这种类型可以是任意的原始类型，例如：整型、字符串或者自定义类型，对数组整体而言类型为 [size]type,对数组元素而言每个元素都是原基本类型
// 重点：数组的定义是 [SIZE]variable_type，说明使用数组表达一个具体的类型的方式就是明确的数组大小和唯一的元素类型，两者关联在一起，才算是完整的类型！
// 数组是包含所有类型的，只是要求元素的类型保持一致即可，比如：整型、字符串型、浮点型、指针类型、其它自定义类型...

var emptyArr []int            // 声明一个空数组，默认为 nil
var emptyArr1 = []int{}       // 声明一个空数组，并初始化
var emptyArr2 []int = []int{} // 声明一个空数组，同上

//var numbers5 [5]int = [5]int{20} // 因为元素的类型为 int 初始化了一个 20，剩下了4个空位就当使用默认值 0 来填充【左边的5和右边的5必须保持一致，才被认为是同一个数据类型】
var numbers5 = [5]int{20}                 // 同上，数组明确的大小和元素类型，那么数组类型就明确了，所以变量后面不需要重复带类型
var balance1 [10]float32                  // 只声明不初始化，所有元素填充默认值
var balance2 = [5]float32{1: 2.0, 3: 7.0} // 选择性初始化，将索引为 1 和 3 的元素初始化，其它给默认值

//var varFloatArr = [...]float32{25, 36, 52.2}
var varFloatArr = []float32{25, 36, 52.2} // 同上，当数组大小不确认时，就不用填充 SIZE，不填充不代表没有，在编译器能力内会自动添加 SIZE【只是个快捷技巧】

// 注意：初始化数组中 {} 中的元素个数不能大于 [] 中的数字

func Example9() {
	println("Example9:")

	// 只声明不初始化的数组默认为nil，声明并初始化为空的数组其值就是空【初始化为{}，里面啥也没有，表示0维度的空数组】
	println(emptyArr)                  // 空表示：[0/0]0x0
	println("返回内容:", emptyArr == nil)  // true
	println(emptyArr1)                 // [0/0]0xf553a0
	println("返回内容:", emptyArr1 == nil) // false

	// 数组遍历
	balance2[4] = 25
	for i, f := range balance2 {
		println(i, f)
	}

	for k := 0; k < len(varFloatArr); k++ {
		fmt.Printf("varFloatArr[%d] = %f\n", k, varFloatArr[k])
	}

	// 多维数组，每个维度的子元素（或子数组）不需要强制对齐，长度是没有限制的
	// 即 子数组1 和 子数组2 的长度可不相同，
	//u := [2][2]int{{10, 41}, {10, 22}}
	//u := [][]int{{10, 41}, {10, 22}}
	u := [...][]int{{10, 41}, {10, 22}} // 只有第一个维度才可以使用 ...
	for _, uv := range u {
		for uk, uuv := range uv {
			println(uk, uuv)
		}
	}
	var varMultiArr [][]string // nil
	row1 := []string{"a", "b"}
	row2 := []string{"c", "d", "e", "f", "g"}
	varMultiArr = append(varMultiArr, row1) // append 是动态拓容的，因为主数组和追加数组，都应该是动态的？？？
	varMultiArr = append(varMultiArr, row2)
	for vark, varv := range varMultiArr {
		for varkk, varvv := range varv {
			fmt.Printf("[%d][%d] = %s \n", vark, varkk, varvv)
		}
	}

	// 数组作为函数参数
	aveArr := []int{1, 47, 24, 23, 4, 7}
	fmt.Printf("这是一个浮点平均值：%f \n", getAverage(aveArr))

	// 精度问题
	// 注意：浮点数计算输出有一定的偏差，这个是计算机本身的问题
	// 解决办法：将小数点数值转化为整形计算
	a := 1.690                              // 表示1.69（扩大了一千倍）
	b := 1.700                              // 表示1.70（扩大了一千倍）
	c := a * b                              // 结果应该是 2873000 表示 2.873
	fmt.Println(c)                          // 中间值
	fmt.Println(float64(c) / (1000 * 1000)) // 显示
	// 不采用精度管理原来的结果是 2.8729999999999998
}

func getAverage(arr []int) float32 {
	var sum float32

	for _, a := range arr {
		sum += float32(a)
	}

	return sum / float32(len(arr))
}