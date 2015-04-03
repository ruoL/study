package main

import (
	"errors"
	"fmt"
)

//go语言中约定错误代码是函数的最后一个返回值
//并且类型是error类型,这是一个内置的接口
func f1(i int) (int, error) {
	if i == 42 {
		//errors.New()使用错误信息作为参数,构建一个基本的错误
		return -1, errors.New("can't work with 42")
	}
	//返回nil表示没有错误
	return i + 3, nil
}

//可以通过实现error接口的方法Error()来自定义错误
//下面我们自定义了一个错误类型来表示上面例子中的参数错误
type myError struct {
	arg  int
	prob string
}

//实现了接口
func (e *myError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(i int) (int, error) {
	if i == 42 {
		//这里我们使用&myError的语法来创建一个新的结构体对象,并且给他的成员赋值
		return -1, &myError{i, "can't work with it"}
	}
	return i + 3, nil
}

func main() {
	//下面的两个循环例子用来测试我们带有错误返回值的函数,在for循环语句里面,使用了if来判断函数返回值是否为nil
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed : ", e)
		} else {
			fmt.Println("f1 worked : ", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed : ", e)
		} else {
			fmt.Println("f2 worked : ", r)
		}
	}
	//如果你需要使用自定义的错误类型返回的错误数据,你需要使用类型断言来获得一个自定义的错误类型实体才行
	_, e := f2(42)
	if ae, ok := e.(*myError); ok {
		fmt.Println(ae.arg, " ", ae.prob)
	}
}
