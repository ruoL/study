package main

import (
	"fmt"
)

type CallBack func(x, y int) int

func main() {
	x, y := 1, 2
	fmt.Println(test(x, y, add))
}

func test(x, y int, callback CallBack) int {
	return callback(x, y)
}

func add(x, y int) int {
	return x + y
}
