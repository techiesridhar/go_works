package main

import (
	"fmt"
)

func main() {
	fmt.Println("Print all the elements in an array")
	array := []int{10, 20, 30, 40, 50, 60}

	for i, val := range array {
		fmt.Println(i, val)
	}

	// print map
	fmt.Println("Print all the elements in a map")
	m := map[string]string{"a": "aaa", "b": "bbb", "c": "ccc"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("print Only keys in map")
	for key := range m {
		fmt.Println(key)
	}

	// print elements in a string
	fmt.Println("Print all the elements in a string")
	for index, ele := range "this is a long string" {
		fmt.Println(index)
		fmt.Printf("%c ", ele)
	}

}
