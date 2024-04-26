package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello World")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	p := Person{"John", 25}
	fmt.Println(p)
	m, _ := json.Marshal(p)
	fmt.Println(string(m))
}
