package main

import "fmt"

func loop(ref string) {
	for i := 0; i < 10; i++ {
		fmt.Println(ref, " -> ", i)
	}
}

func main() {
	go loop("one")
	go loop("two")
	go loop("three")

	var input string
	fmt.Scanln(&input)
}
