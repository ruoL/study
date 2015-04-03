package main

import (
	"fmt"
	"time"
)

func main_01() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received : ", msg1)
		case msg2 := <-c2:
			fmt.Println("received : ", msg2)
		}
	}
}

//----------------------------------------------------------//
//超时机制
func main() {
	//在这个例子当中,假设我们执行了一个外部调用,2秒之后将结果写入c1
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result"
	}()

	select {
	//这里使用select来实现超时
	case res := <-c1: //等待通道结果
		fmt.Println(res)
	case <-time.After(time.Second * 1): //time.after()在等待1秒时候返回一个值
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
