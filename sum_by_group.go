package main

import "fmt"

func sum(a []int, result chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	result <- sum
}

func main() {
	a := []int{2, 3, 5, 6, 10, -5, 1, 0}
	result := make(chan int)
	go sum(a[:len(a)/2], result)
	go sum(a[len(a)/2:], result)
	x, y := <-result, <-result

	fmt.Println(x, y, x+y)
}

/*
goroutine是Go语言并行设计的核心。goroutine是一种比线程更轻量的实现，
十几个goroutine可能在底层就是几个线程。要使用goroutine只需要简单的
在需要执行的函数前添加go关键字即可。当执行goroutine时候，go语言立即
返回，接着执行剩余的代码，goroutine不阻塞主线程。channel就像一个管道，
但是可以双向传输数据，通过它我们可以接收和发送数据，假如result 是一
个channel那么：result <- value 是将数据发送到result, 而key <- result
就是从result中接收一个数据，就如以上的代码所示，值得注意的地方是
channel只能通过Go语言内建的函数make(chan type)创建，其中type指明了
该channel能传递的数据类型。

下面我们来说说以上的代码。在main函数中，我们声明了一个int类型的切片a，
然后通过内置函数make创建了一个能接收和发送int类型的channel。

然后通过关键字go执行了两个goroutine，这两个goroutine的功能是
分别计算切片a前半部分和后半部分的和。在这里main函数碰到go关键字
派发goroutine执行相应的函数后，立即返回执行剩余的代码，不会等待
goroutine的返回。sum函数中，计算切片的和，然后将结果发送到channel。
接下来main函数，从channel中获取结果，在这里，main函数会阻塞至直到能
从channel result中接收到数据，最后我们打印出了结果。
*/
