package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name string `json:"name"`
	Age  int
}

func main_01() {
	f, err := os.Create("data.dat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := &Student{Name: "zhangsan", Age: 23}
	encoder := json.NewEncoder(f)
	encoder.Encode(s)

	f.Seek(0, os.SEEK_SET)
	decoder := json.NewDecoder(f)
	var sl Student
	decoder.Decode(&sl)
	fmt.Println(sl)
}

func main_02() {
	s := &Student{"zhangsan", 8}
	buff, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(buff))

	var sl Student
	json.Unmarshal(buff, &sl)
	fmt.Println(sl)
}

func main() {
	str := `{"name":"zhang", "Age":122}`
	var m map[string]interface{}
	json.Unmarshal([]byte(str), &m)
	for k, v := range m {
		switch v.(type) {
		case float64:
			fmt.Println(k, "-->", v)
		case string:
			fmt.Println(k, "-->", v)
		default:
			fmt.Println(k, "mistake")
		}
	}
}
