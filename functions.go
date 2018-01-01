package main

import "fmt"

// multiple arguments
func add(a int, b int) int {
	return a + b
}

// signature with one declerations
func mul(a, b int) int {
	return a * b
}

// Multiple return
func hello(h string, w string) (string, int) {
	return h + w, 123
}

// functin with variable number of arguments
func reduce(numbers ...int) int {
	var sum int = 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Closures AKA python lamda functions
// Ananymous function

func supply_func(start int) func() int {
	return func() int {
		start = start + 1
		return start
	}
}

// Lets try recursion
func fib(num int) int {
	if num < 3 {
		return num - 1
	} else {
		return fib(num-2) + fib(num-1)
	}
}

func main() {
	a := add(1, 4)
	m := mul(4, 8)

	var hw string
	var status int

	hw, status = hello("hello ", "world")

	s := reduce(1, 2, 3, 4, 5, 23)

	fmt.Println("1+4 =", a)
	fmt.Println("4*8 =", m)
	fmt.Println("Hw = ", hw, status)
	fmt.Println("Reduce = ", s)

	// Get func from supplyfunc
	anon := supply_func(5)
	fmt.Println("anon =", anon())
	fmt.Println("anon =", anon())
	fmt.Println("anon =", anon())
	fmt.Println("anon =", anon())

	// once again supply func
	anon2 := supply_func(5)
	fmt.Println("anon2 =", anon2())
	fmt.Println("anon2 =", anon2())
	fmt.Println("anon2 =", anon2())

	// Lets calculate the fibonacci
	fmt.Println("fibooo 10 = ", fib(10))
	fmt.Println("fibooo 15 = ", fib(15))
	fmt.Println("fibooo 30 = ", fib(30))
}
