package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	// This is how you read data from command line
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	text, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(text[:len(text)-1])

	if num < 5 {
		fmt.Println("num less than 5")
	} else if num < 10 {
		fmt.Println("Not lessthan ten")
	} else {
		fmt.Println("Greater than 10")
	}

	// This is switch case statments
	time_now := time.Now()

	switch time_now.Weekday() {
	case time.Monday, time.Tuesday:
		fmt.Println("week is starting")
	case time.Wednesday:
		fmt.Println("mid week")
	case time.Friday, time.Saturday:
		fmt.Println("yeyyyy weekend coming")
	default:
		fmt.Println("looks like sunday")
	}

}
