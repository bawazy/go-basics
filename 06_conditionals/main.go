package main

import "fmt"

func main() {
	x := 10
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n ", x, y)
	} else {
		fmt.Printf("%d is less than %d\n ", y, x)
	}

	color := "Red"

	// if color == "Red" {
	// 	fmt.Println("Color is Red")
	// } else if color == "Blue" {
	// 	fmt.Println("Color is Blue")
	// } else {
	// 	fmt.Println("Color is neither Blue nor Red")
	// }

	//Switch
	switch color {
	case "Red":
		fmt.Println("Color is Red")
	case "Blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color is Neither Red Nor Blue")
	}
}
