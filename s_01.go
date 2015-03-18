package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//获取随机数
	num := create_random_num()
	fmt.Printf("Please input num :")
	//游戏是否结束
	flag := true
	reader := bufio.NewReader(os.Stdin)
	for flag {
		data, _, _ := reader.ReadLine()
		command, err := strconv.Atoi(string(data))
		if err != nil {
			fmt.Println("Please input num :")
		} else {
			fmt.Println("Your num is :", command)
			if command == num {
				flag = false
				fmt.Println("Yes !")
			} else if command < num {
				fmt.Println("Min try again")
			} else if command > num {
				fmt.Println("Max try again")
			}
		}
	}
}

func create_random_num() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(100)
}
