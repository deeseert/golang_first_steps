package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)

	// Use * to read value from memory address
	fmt.Println(*b)  // 5
	fmt.Println(*&a) // 5

	// Use * to change the value using the pointer to the memory address
	*b = 10
	fmt.Println("This is the new value of a: ", a) // a := 10
}
