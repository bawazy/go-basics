package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// // assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"
	// declare map and assign
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com", "Mike": "mike@gmail.com"}
	fmt.Println(len(emails))
	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	// Delete
	delete(emails, "Bob")
	emails["hayatu"] = "hayatu@gmail.com"
	fmt.Println(emails)

}
