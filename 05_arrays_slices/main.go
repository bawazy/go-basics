package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	//assign values
	// fruitArr[0] = "Apples"
	// fruitArr[1] = "Orange"

	//declare and assign

	// fruitArr := [2]string{"Apples", "Oranges"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	// Slices

	fruitSlices := []string{"Apples", "Oranges", "Grapes"}
	fmt.Println(len(fruitSlices))
	fmt.Println(fruitSlices[1:2])

}
