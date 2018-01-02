package main

import "fmt"

// This is interface, which can be used in places where
// you want some structs to be passed with talk() and run()
// must present
// For eg: func animal_speak(a animal)
// Means animal_speak function accepts any struct which implements
// Interface animal

type animal interface {
	talk() string
	run() string
}

func animal_speak(a animal) {
	fmt.Println(a.talk())
}

type dog struct {
	sound string
	speed int
}

type cat struct {
	sound string
}

func (d dog) talk() string {
	return "talking -> " + d.sound
}

func (d dog) run() string {
	return "Can run @ speed " + string(d.speed)
}

func (c cat) talk() string {
	return "talking -> " + c.sound
}

// If we comment below method and run this code
// We get error: cat does not implement animal (missing run method)
func (c cat) run() string {
	return "cant run"
}

// ------------------------ main ----------------------
func main() {
	var c cat = cat{sound: "MEOW"}
	animal_speak(c)
}
