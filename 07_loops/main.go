package main

import "fmt"

func main() {
	//long Method
	i := 0
	for i <= 10 {
		i++
		fmt.Println(i)
	}
	// short Method
	for i := 1; i <= 10; i++ {
		fmt.Printf("number %d\n", i)
	}

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
