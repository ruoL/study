package main

import (
	"fmt"
	"runtime"
	"time"
)

func SayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello--->")
		runtime.Gosched()
	}
}

func SayWrold() {
	for i := 0; i < 10; i++ {
		fmt.Println("World")
		runtime.Gosched()
	}
}

func main_01() {
	go SayHello()
	go SayWrold()
	time.Sleep(1 * time.Second)
	fmt.Println("cpu is : ", runtime.NumCPU())
}

func producer(c chan int) {
	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i
	}

}

func consumer_01(c, f chan int) {
	for {
		if v, ok := <-c; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
	f <- 1
}

func consumer(c, f chan int) {
	for v := range c {
		fmt.Println(v)

	}
	f <- 1
}

func main_02() {
	buff := make(chan int)
	flag := make(chan int)
	go producer(buff)
	go consumer(buff, flag)
	<-flag
}

func main_03() {
	c := make(chan int)
	select {
	case <-c:
		fmt.Println("input data ")
	case <-time.After(5 * time.Second):
		fmt.Println("chao shi tui chu")

	}
}

//------------------------------------------------------------//
//假如我们有一个函数叫做f(s)
//这里我们使用通常的同步调用来调用函数
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	//为了能让这个函数, 运行使用go f(s)
	//这个协程将和调用他的协程并行执行
	go f("goroutine")
	//可以为一个匿名函数开启一个协程

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	//上面的协程在调用之后就开始异步执行了,所以程序不用等待他们执行完成就跳到这里来了,下面的scanln用来从命令行获取一个输入然后才让main函数执行,如果没有下面的语句
	//程序到了这里就会直接退出,而上面的协程还没有来得及执行main函数就已经退出了,所以不会看见执行的结果
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
