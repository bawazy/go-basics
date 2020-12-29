package main

import "fmt"

func main() {
	name := "hayatu"
	var sizes float32 = 4.9
	var age int32 = 24
	const isCool = true
	size := 1.3
	firstName, lastName := "Hayatu", "Bayero"
	fmt.Println(name, age, isCool, size, firstName, lastName)

	// finding the type of a variable
	fmt.Printf("%T\n", sizes)

}
