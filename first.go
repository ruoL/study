package main

import io "fmt"

func main1() {
	var a bool

	//var b int = 1
	//var c int
	//c = 2
	//d := 3

	io.Println(a)
}

func mainj() {
	var a, b, c, d int = 1, 2, 3, 4
	io.Println(a)
	io.Println(b)
	io.Println(c)
	io.Println(d)

	var cc int = 65
	bb := string(cc)
	io.Println(bb)
}

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
)

func maink() {
	io.Println(B)
	io.Println(KB)
	io.Println(MB)
	io.Println(GB)
}

func mainl() {

LABEL:
	for i := 0; i < 10; i++ {
		for {
			io.Println(i)
			continue LABEL
			io.Println(i + 10)
		}
	}

}

func maina() {
	//a := [2]int{1, 2}
	//b := [2]int{2, 3}
	aa := [...][3]int{
		{1, 2, 2: 10},
		{2, 3, 4}}
	io.Println(aa)

	cc := [...]int{3, 6, 2, 9, 0, -1, 22}

	var num int = len(cc)
	for i := 0; i < num; i++ {
		for j := 0; j < num-i-1; j++ {
			if cc[j] < cc[j+1] {
				temp := cc[j]
				cc[j] = cc[j+1]
				cc[j+1] = temp
			}
		}
	}
	io.Println(cc)
}

func main() {

	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	sb := a[3:5]
	io.Println(string(sb))
}
