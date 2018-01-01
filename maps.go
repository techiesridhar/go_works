package main

import "fmt"

func main() {

	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)

	// How to delete
	delete(m, "one")
	fmt.Println(m)

	// how to check if key exists
	_, ok := m["key doesnt exist"]
	fmt.Println(ok)

	m["test"] = 11111111

	// how to loop
	for key, value := range m {
		fmt.Println("Key", key, "  -  Value:", value)
	}
}
