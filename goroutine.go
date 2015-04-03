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

func main() {
	c := make(chan int)
	select {
	case <-c:
		fmt.Println("input data ")
	case <-time.After(5 * time.Second):
		fmt.Println("chao shi tui chu")

	}
}
