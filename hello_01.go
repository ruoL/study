package main

import (
	"fmt"
	"os"
)

func main() {
	str := "World"
	//os.Args 表示传入的参数, 当传入的参数大于1时进行下面的操作
	if len(os.Args) > 1 {
		str = os.Args[1]
	}
	fmt.Println("Hello", str)
}
