package main

import (
	"fmt"
	"strconv"
)

func main() {
	// String to integer (err: anything could go wrong)
	val, _ := strconv.Atoi("123")
	fmt.Println(val)

	// Integer to string, nothing can go wrong
	value := strconv.Itoa(1111111111)
	fmt.Println(value)

	// Simillarly there are many
	// b, err := strconv.ParseBool("true")
	// f, err := strconv.ParseFloat("3.1415", 64)
	// i, err := strconv.ParseInt("-42", 10, 64)
	// u, err := strconv.ParseUint("42", 10, 64)
	// s := strconv.FormatBool(true)
	// s := strconv.FormatFloat(3.1415, 'E', -1, 64)
	// s := strconv.FormatInt(-42, 16)
	// s := strconv.FormatUint(42, 16)
	// -------------------------

	q := strconv.Quote("Hello, 世界")
	fmt.Println(q)

	qq := strconv.QuoteToASCII("Hello, 世界")
	fmt.Println(qq)

}
