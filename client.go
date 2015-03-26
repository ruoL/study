/************************************************
*作者：张天其
*时间：2015年3月26日上午9点46分
*描述：通过go语言的net包实现了简单的客户端连接服务器，以下是客户端
************************************************/

package main

import (
	"bufio" //用于读写的包bufio
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

//链接服务器
func connectServer() {
	//链接
	conn, err := net.Dial("tcp", "localhost:7777")
	checkError(err)
	fmt.Println("client success !")
	//接受屏幕的输入
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Who are you !")
	name, _ := inputReader.ReadString('\n')

	trimName := strings.Trim(name, "\r\n")
	conn.Write([]byte(trimName + "client \n"))
	for {
		fmt.Println("let's wo chat, if you want to quit,input quit")
		input, _ := inputReader.ReadString('\n')
		trimInput := strings.Trim(input, "\r\n")
		//当用户输入quit的时候退出
		if trimInput == "quit" {
			fmt.Println("bye")
			conn.Write([]byte(trimName + " exit"))
			return
		}

		_, err := conn.Write([]byte(trimName + "says" + trimInput))
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("an error !", err.Error())
	}
}

//主函数
func main() {
	//链接server
	connectServer()
}
