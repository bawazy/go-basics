package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// greeting Method(value Receiver)
func (p Person) greet() string {
	return " Hello my Name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// has birthday method (Pointer Receiver)
func (p *Person) hasBirthday() {
	p.age++
}

//gets married (Pointer Receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else if p.gender == "f" {
		p.lastName = spouseLastName
	}

}

func main() {
	person1 := Person{"Samantha", "Smith", "Boston", "f", 25}
	Fahad := Person{"Fahad", "Abdullahi", "Yola", "m", 25}

	person1.hasBirthday()
	person1.getMarried("Waziri")
	Fahad.getMarried("Aisha")
	fmt.Println(person1.greet())
	fmt.Println(Fahad.greet())

}
