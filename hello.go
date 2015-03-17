/*
可见性规则:
	变量名称开头小写表示私有
	变量名称开头大写表示公有
*/

//当前程序包名
package main

//导入其他包
//import io "fmt"

//省略调用
import . "fmt"

//常量的定义
const PI = 3.14

//常量的定义第二种方法
const (
	PI     = 3.14
	const1 = 1
	const2 = 2
	const3 = 3
)

//全局变量的声明和赋值
var name = "zhang"

//全局变量的声明和赋值, 这种方法只能在这里使用
var (
	name  = "zhang"
	name1 = "tian"
	name2 = 1
	name3 = 3
)

//一般类型的声明
type newType int

type (
	newType int
	type1   float32
	type2   string
	type3   byte
)

//结构类型
type student struct{}

//接口类型
type golang interface{}

//函数
func main() {
	io.Println("Hello World!")
}
