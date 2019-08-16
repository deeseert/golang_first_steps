package main

import "fmt"

// Create a function.
// You need to specify the input type and the output type.
// In this case, you put it a string and you want to get a string back
const myName = "Terry"

func greeting(name string) string {
	return "Hello " + name
}

// When the input arguments are of the same type, you could also 
// avoid to specify the first one
func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(getSum(3, 4))
}
