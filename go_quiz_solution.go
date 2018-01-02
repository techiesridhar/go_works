package main

import "os"
import "fmt"
import "strings"
import "time"
import "io/ioutil"

func timer() {

}

func main() {
	args := os.Args[1:]
	var questions_file string = "quiz_questions.txt"

	if len(args) > 0 {
		questions_file = args[0]
	}

	fmt.Println("File name = ", questions_file)

	// read problems file
	dat, _ := ioutil.ReadFile(questions_file)

	ticker := time.NewTicker(time.Second)
	var timer int

	go func() {
		for range ticker.C {
			timer++

			if timer == 20 {
				fmt.Println("you took ", timer, " seconds, Game exiting")
				os.Exit(1)
			}
		}
	}()

	var score int
	var input string
	for i, str := range strings.Split(string(dat), "\n") {
		operations := strings.Split(str, ",")
		if len(operations) == 2 {
			fmt.Println("Your ", i, "question is", operations[0])

			fmt.Scanln(&input)
			if input == operations[1] {
				score++
			}
		}
	}
	fmt.Println("Your score is", score)

}
