package main

import "fmt"

func main() {
	// Define map
	// make(map[]) takes is 2 arguments. First one is the key and second one is the value
	// and is written this way:
	emails := make(map[string]string)

	// Assign key/value
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)

	// Delete
	delete(emails, "Bob")
	fmt.Println(emails)

	// Declare map and assign kay/value pair
	otherEmails := map[string]string{
		"Terry":   "terry@gmail.com",
		"Antonio": "antonio@gmail.com"}

	otherEmails["Giuseppe"] = "giuseppe@gmail.com"

	fmt.Println(otherEmails)

}
