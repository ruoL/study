package main

import (
	"fmt"
)

/*
go语言支持匿名函数,匿名函数可以形成闭包,闭包函数可以访问定义闭包的函数定义的内部变量
*/
//-----------------------------------------------------------//
//这个intSeq函数返回另外一个在intSeq内部定义的匿名函数,这个返回的匿名函数包住了变量i,从而形成了闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main_01() {
	//我们调用intSeq函数,并且把结果赋值给一个函数nextInt,这个函数拥有自己的变量i这个变量每次调用都被更新
	//这里的i的初始值是由intSeq调用的时候决定的
	nextInt := intSeq()
	//调用几次nextInt看看闭包的结果
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	//为了确认闭包的状态是独立于intSeq函数,在创建一个
	newInts := intSeq()
	fmt.Println(newInts())
}
//1
//2
//3
//1

//-----------------------------------------------------------//
func main_02() {
	add10 := closure(10) //其实是构造了一个加10的函数
	fmt.Println(add10(5))
	fmt.Println(add10(6))
	add20 := closure(20)
	fmt.Println(add20(5))
}

func closure(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}
//15
//16
//25

//-----------------------------------------------------------//
func main() {
	var fs []func() int
	for i := 0; i < 3; i++ {
		fs = append(fs, func() int {
			return i
		})
	}
	for _, f := range fs {
		fmt.Printf("%p = %v\n", f, f())
	}
}
//3
//3
//3

//-----------------------------------------------------------//
func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main_04() {
	result := addr()
	for i := 0; i < 10; i++ {
		fmt.Println(result(i))
	}
}
