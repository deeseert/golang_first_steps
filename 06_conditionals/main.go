package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(50)
	y := rand.Intn(50)

	// If else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "red"
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is NOT red or blue")
	}

	// Switch Statement
	car := "Audi"
	switch car {
	case "Audi":
		fmt.Println("Your car is an Audi")
	case "BMW":
		fmt.Println("Your car is a BMW")
	default:
		fmt.Println("You don't have a car!")
	}

}
