package main

import (
	"fmt"
	"sync"
	//"time"
)

var c chan int

func ready(w string, sec int64) {
	//time.Sleep()
	fmt.Println(w, "is ready !")
	c <- 1
}

func main1() {
	c = make(chan int)
	go ready("Tee", 2)
	go ready("Coffee", 1)
	fmt.Println("I'am waiting, but not too lang")
	<-c
	<-c
}

type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

func handle(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

var s string
var once sync.Once

func setUp() {
	s = "Hello World"
}

func doprint() {
	once.Do(setUp)
	fmt.Println(s)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	twoprint()
}
