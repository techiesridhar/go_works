package main

import "fmt"

func main() {

	// this is like while
	fmt.Println("This is for (like while)")
	i := 1
	for i < 4 {
		fmt.Println(i)
		i++
	}

	fmt.Println("this is c like for loop")
	for i := 1; i < 8; i++ {
		fmt.Println(i)
	}

	// Continue and break, Print odd numbers
	fmt.Println("this is continue demo")
	for i := 1; i < 8888888; i++ {

		// Continue if its even number
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)

		// break if its > 10
		if i > 10 {
			break
		}

	}

}
