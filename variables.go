package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x string = "Test"
	fmt.Println(x)

	var y int = 100
	fmt.Println(y)

	var z, j = 9, 30
	fmt.Println(z, j)

	var booo = true
	fmt.Println(booo)

	// shortcut to declare a variable as string (var f string = "self declare")
	s := "self declare"
	fmt.Println(s)
	I := 12
	fmt.Println(I)

	// Check what type of variable is this
	var d = 10
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))
}
