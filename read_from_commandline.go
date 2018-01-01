package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://golang.org/pkg/strings/

func main() {
	// This is how you read data from command line
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter something: ")
	text, _ := reader.ReadString('')

	fmt.Println(strings.Repeat(text, 10))

	// One more way to scan
	var a string
	fmt.Scan(&a)
	fmt.Println(a)
}
