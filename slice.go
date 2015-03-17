package main

import "fmt"

func main() {
	//先定义一个数组
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//基于数组创建一个切片, 0-4
	var mySlice []int = arr[:5]

	//遍历
	fmt.Println("Elements of myArray :")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

/*
	创建切片的方法
	1. slice = arr[:]  //表示数组的所有元素
	2. slice = arr[:5] //表示数组的前5个元素, 0-4, 不包含arr[5]
	3. slice = arr[5:] //表示从第五个元素开始的所有元素
*/
