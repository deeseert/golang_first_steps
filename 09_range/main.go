package main

import "fmt"

func main() {
	ids := []int{12, 454, 65, 87, 32, 9, 54, 32, 87, 543, 5}
	// This loop has always index and the element you want to do something with
	// If you don't use it, put _
	for i, ids := range ids {
		fmt.Printf("at the index %d you have ID %d\n", i, ids)
	}

	// Without using index
	for id := range ids {
		fmt.Printf("you have ID %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Terry": "terry@gmail.com", "Antonio": "antonio@gmail.com"}

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}

	for key := range emails {
		fmt.Println("Name: ", key)
	}
}
