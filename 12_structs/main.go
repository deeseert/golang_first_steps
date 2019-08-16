package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// METHODS

// Define Greeting method here (value receiver)
// p is an identifier like this in JavaScript
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Define hasBirthday method here (Pointer Receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	// Initialize the person struct here
	person1 := Person{firstName: "Gio", lastName: "Di Nardo", city: "London", gender: "x", age: 30}
	// person1 := Person{"Gio", "Di Nardo", "London", "m", 30}

	// Can also access one only filed
	fmt.Println(person1)
	fmt.Println(person1.firstName)

	// Or change one thing
	// person1.age++

	// Call the Pointer Receiver here
	person1.hasBirthday()
	person1.getMarried("Tua madre")

	// Call the Value Receiver here
	fmt.Println(person1.greet())
}
