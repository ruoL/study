/************************************************
*作者：张天其
*时间：2015年3月26日上午9点46分
*描述：通过go语言的net包实现了简单的客户端连接服务器，以下是服务器端
************************************************/

package main

import (
	"fmt"
	"log"
	"net" //支持通信的包
)

func startServer() {
	//连接主机，端口，采用tcp连接，监听7777端口
	listener, err := net.Listen("tcp", "localhost:7777")
	checkError(err)
	fmt.Println("server success !")
	//循环等待客户端接入
	for {
		conn, err := listener.Accept()
		checkError(err)
		//开启一个goroutines处理客户端消息，实线并发只要go一下就好
		go doServerStuff(conn)
	}
}

//处理客户端的消息
func doServerStuff(conn net.Conn) {
	//生成一个缓存名字数组
	nameInfo := make([]byte, 512)
	//将消息读取到nameInfo
	_, err := conn.Read(nameInfo)
	checkError(err)
	for {
		//创建一个存放消息的切片
		buff := make([]byte, 512)
		//将客户端发来的消息读取到buff里面
		_, err := conn.Read(buff)
		//错误检查
		flag := checkError(err)
		if flag == 0 {
			break
		}
		fmt.Println(string(buff))
	}
}

//错误检查函数
func checkError(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			return 0
		}
		log.Fatal("an error !", err.Error())
		return -1
	}
	return 1
}
func main() {
	//开启服务
	startServer()
}
