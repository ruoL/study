package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	outfile := "infile.data"
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 1000000; i++ {
		file.WriteString(strconv.Itoa(r.Intn(1000000)) + "\n")
	}
}
