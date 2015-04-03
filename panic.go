package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}
}
