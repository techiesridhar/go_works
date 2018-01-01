package main

import (
	"fmt"
)

// Go supports constants of character, string, boolean, and numeric values.
const a int = 30

func main() {

	// do some read only actions
	const b int = 10
	fmt.Println(a / b)

	// Update the constant
	// salary = age * 90
	// you will get error "cannot assign to salary"
}
