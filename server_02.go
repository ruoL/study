package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

const (
	LS  = "LS"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
//在7076端口监听
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":7076")
	checkError(err)
	listener, err1 := net.ListenTCP("tcp", tcpAddr)
	checkError(err1)
	for {
//监听客户端,等待链接
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println(err2)
//通常服务端为一个服务,当出错的时候不会退出,出错后继续等待下一个链接
			continue
		}
		fmt.Println("收到客户端请求")
		go ServerClient(conn)//开启一个线程处理客户端的请求
	}
}

//处理客户端的请求
func ServerClient(conn net.Conn) {
	defer conn.Close()
	str := ReadData(conn)//从conn读取数据
	if str == "" {
		SendData(conn, "接受数据时出错")
		return
	}
	fmt.Println("收到命令", str)
	switch str {
	case LS:
		ListDir(conn)
	case PWD:
		Pwd(conn)
	default:
		if str[0:2] == CD {
			Chdir(conn, str[3:])
		} else {
			SendData(conn, "命令错误")
		}
	}
}

func Chdir(conn net.Conn, s string) {
	err := os.Chdir(s)
	if err != nil {
		SendData(conn, err.Error())
	} else {
		SendData(conn, "ok")
	}
}

func ListDir(conn net.Conn) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		SendData(conn, err.Error())
		return
	}
	var str string
	for i, j := 0, len(files); i < j; i++ {
		f := files[i]
		str += f.Name() + "\t"
		if f.IsDir() {
			str += "dir\r\n"
		} else {
			str += "file\r\n"
		}
	}
	SendData(conn, str)
}
//读取数据
func ReadData(conn net.Conn) string {
	var data bytes.Buffer
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return ""
		}

		if buf[n-1] == 0 {
			data.Write(buf[0 : n-1])
			break
		} else {
			data.Write(buf[0:n])
		}
	}
	return string(data.Bytes())
}

func SendData(conn net.Conn, data string) {
	buf := []byte(data)
	buf = append(buf, 0)
	_, err := conn.Write(buf)
	if err != nil {
		fmt.Println(err)
	}
}

func Pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		SendData(conn, err.Error())
	} else {
		SendData(conn, s)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
