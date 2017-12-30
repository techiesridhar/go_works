package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Self declared
	h := 1
	fmt.Println(reflect.TypeOf(h))

	// Declared
	var x string = "hey there"
	fmt.Println(reflect.TypeOf(x))
}
