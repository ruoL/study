package main

import (
	"fmt"
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {

}

func create() {
	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}

func truncate() {
	err = os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}

func getinfo() {
	fileinfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file name :", fileinfo.Name())
	fmt.Println("size in bytes :", fileinfo.Size())
	fmt.Println("Permissions :", fileinfo.Mode())
	fmt.Println("Last modefied :", fileinfo.ModTime())
	fmt.Println("is Dir :", fileinfo.IsDir())
	fmt.Println("system interface type : %T\n", fileinfo.Sys())
	fmt.Println("system info : %+v\n\n", fileinfo.Sys())
}

func move() {
	path := "test.txt"
	newPath := "test2.txt"
	err := os.Rename(path, newPath)
	if err != nil {
		log.Fatal(err)
	}
}

func remove() {
	err := os.Remove("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}
