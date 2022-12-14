package chapter2

import (
	"fmt"
)

// 使用协程的格式：go 函数名( 参数列表 )
// 通道，同样支持 range，支持遍历的数据还包括：数组、切片、集合、通道

func Example20() {
	fmt.Println("Example20:")

	// 1.协程之间的独立性
	// 从结果可以看出，两个同时执行的协程，毫无规律的在执行
	//go say("world") // 新开的协程
	//say("hello")    // 当前协程

	// 2.协程之间的数据共享方式：chan
	// 将一个计算任务，人为拆解为多个子任务，分别交由不同的协程执行，最后将结果重新整合，从而达到并行计算的目的
	i := []int{2, 4, 7, 8, 5, 10, 5, 9, -5, 3, 4}
	i1 := i[:len(i)/2] // 这里是以5分割
	i2 := i[len(i)/2:]
	c1 := make(chan int) // 不带缓冲区，为同步通道
	c2 := make(chan int)
	go sum(i1, c1)
	go sum(i2, c2)
	x, y := <-c1, <-c2 // 接收方，如果先执行到这一步，永远处于阻塞等待状态
	println(x, y, x+y)

	// 3.通道/管道（同步、异步）
	// ch := make(chan int)
	// 注意：make() 实例化函数，只用来操作 slice, map, chan (only)
	// 通道的逻辑：
	// 一、不带缓冲区的通道，是同步状态，发送方必须阻塞等待有接收方将数据读取完后才算结束，接收方必须阻塞等待直接有数据接收到才会向下执行，两者在时间上是相互等待状态，两者在传输动作上是同步进行的。
	// 二、带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，比如：发送方已经执行完了，并将数据推送至通道，发送方逻辑就算结束了！而接收方随时都可以读取这个数据（两者可以一前一后）。
	// 三、缓冲区有大小限制，如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值，此时接收方在有值可以接收之前会一直阻塞。
	// 以上三种情况，不论哪种，接收方如果先执行，永远处于阻塞等待状态！
	ch := make(chan int, 2) // 这里必须设置缓冲，才能实现当前协程中同步传递数据
	ch <- 1                 // 在同步执行中，1先进入管道队列
	ch <- 2                 // 2后进入管道队列
	fmt.Println(<-ch)       // 先进先出，channel 就是队列结构
	fmt.Println(<-ch)
	// 缓冲管道相当于在发送方与接收方之间增加了一个数据传输的队列！进而实现两者之间的传输由同步转化为了异步。
	// 所以，管道的缓冲大小表示的就是队列可存储的元素个数。

	// 4.通道/管道关闭
	// close(ch)
	// x, ok := <-ch // 用法类似于 读取集合中的元素
	// 重点：关闭通道并不会丢失里面的数据，只是让读取通道数据的时候不会在读完之后一直阻塞等待新数据写入【关闭，可以理解为将关闭标记插入到队列的最后，range 或者 select 或者 <- 可以自动检测到】
	ch1 := make(chan int, 2)
	ch1 <- 25
	close(ch1)     // 很显示，这里是关闭了，但是管道中的数据还在缓冲区里
	_, ok := <-ch1 // 因为管道内只有一个数据，这里读取了就没有了，这意味着这个读释放了阻塞！！！这时阻塞状态已解除，如果还要读管道就会返回异常参数
	fmt.Printf("关闭通道后，返回：%#v \n", ok)
	//_, ok1 := <-ch1
	//fmt.Printf("关闭通道后，返回：%#v \n", ok1) // 再去
}

func sum(i []int, c chan int) {
	sum := 0
	for _, v := range i {
		sum += v
	}
	c <- sum
	close(c)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

// 复习一下make()的用法：
// make([]string) // 切片
// make(map[string]int) // 集合
// make(chan int) // 通道【格式特殊，有空格】
