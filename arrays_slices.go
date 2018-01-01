package main

import "fmt"

func main() {
	// https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html
	// Arrays are fixed-length sequences of items of the same type.
	// Arrays in Go can be created using the following syntaxes:

	// [N]Type
	// [N]Type{value1, value2, ..., valueN}
	// [...]Type{value1, value2, ..., valueN}

	// Array with five positions
	var an_array [5]int
	an_array[0] = 1
	an_array[1] = 2
	fmt.Println(len(an_array))
	fmt.Println(an_array)

	// Slices, on the other hand, are much more flexible, powerful, and convenient than arrays. Unlike arrays, slices
	// can be resized using the built-in append function.
	var s []int
	fmt.Println(s)

	s = append(s, 0)
	fmt.Println(s)

	s = append(s, 23, 233, 344, 23)
	fmt.Println(s)

	//Slices can be created using the following syntaxes:

	//make([]Type, length, capacity)
	//make([]Type, length)
	//[]Type{}
	//[]Type{value1, value2, ..., valueN}

	str_slice := make([]string, 3)
	str_slice = append(str_slice, "narendra", "nami")
	fmt.Println(str_slice)

	// Watch the lenght of str_slice is not 3, as apended two new values
	// so the length of str_slice is 5 now
	fmt.Println(len(str_slice))

	str_two_slice := make([]string, len(str_slice))
	copy(str_two_slice, str_slice)
	fmt.Println(str_two_slice)

	// Create slice with value run time
	t := []string{"AA", "BB", "CC"}
	fmt.Println(t)

}
