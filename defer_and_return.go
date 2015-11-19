package main

import (
	"fmt"
)

func main() {
	// fmt.Println("func_1 return is :", func_1())
	// fmt.Println("func_2 return is :", func_2())
	c := func_3()
	fmt.Println("func_3 return is :", *c, c)
}

//返回值没有提前声明,是来自其他变量的赋值,defer中修改的是其他变量
func func_1() int {
	var i int
	defer func() {
		i++
		fmt.Println("func_1 defer 2 -> ", i)
	}()
	defer func() {
		i++
		fmt.Println("func_1 defer 1 -> ", i)
	}()
	return i
}

//返回值提前声明,返回的就是变量i,defer操作的也是i
func func_2() (i int) {
	defer func() {
		i++
		fmt.Println("func_2 defer 2 -> ", i)
	}()
	defer func() {
		i++
		fmt.Println("func_2 defer 1 -> ", i)
	}()
	return
}

func func_3() *int {
	var i int
	defer func() {
		i++
		fmt.Println("func_3 defer 2 ->", i, &i)
	}()
	defer func() {
		i++
		fmt.Println("func_3 defer 1 ->", i, &i)
	}()
	return &i
}
