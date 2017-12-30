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

	s := "self declare"
	fmt.Println(s)
	I := 12
	fmt.Println(I)

	var d = 10
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))
}
