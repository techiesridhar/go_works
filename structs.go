package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
}

// Methods inside structs like methods inside a class/objects

func (p Person) talk(lang string) string {
	temp := p.name + " is talking " + lang
	fmt.Println(temp)
	return temp
}

func (p Person) walk() string {
	return p.name + " is walking"
}

// Struct as argument to function
func just(human Person) {
	fmt.Println(human)
}

func main() {
	// Create a person
	var naren Person
	naren.name = "naren"
	naren.age = 12
	naren.gender = "male"

	fmt.Println(naren)

	// Create one more person
	var nami Person = Person{name: "namitha", age: 13, gender: "f"}
	just(nami)

	// Access struct attributes
	fmt.Println(nami.gender)

	// Now call method of struct jsut like obj.method()
	nami.talk("English")
}
