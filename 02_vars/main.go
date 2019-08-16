package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for uint32
	// float32 float64
	// complex64 complex128

	// Using var/const
	var name = "Gioacchino"
	var age = 30

	// Using shorthand
	// name2 := "Terry"
	// email := "terry@gmail.com"
	name2, email := "Terry", "terry@gmail.com"
	size := 1.2

	// fmt.Println() is like console.log() in JavaScript
	fmt.Println("This is the var name:", name, "This is age", age, name2, email)

	// This tells you the type of the variable
	fmt.Printf("%T\n", size)
}
