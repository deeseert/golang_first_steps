package main

import "fmt"

// FIRST WAY
// Delcare array
var fruitArr [2]string

func main() {
	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// SECOND WAY
	// Declare and Assign alltogether
	fruitSlice := [2]string{"Banana", "Melon"}
	fruitSlice2 := []string{"Banana", "Melon", "more fruit", "all", "elements", "You", "what"}

	fmt.Println(fruitArr)
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice2)
}
