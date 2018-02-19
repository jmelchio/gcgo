package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	showPointer()
	somePerson := person{
		first: "John",
		last:  "Doe",
		age:   12,
	}
	fmt.Println(somePerson)
	changeMe(&somePerson)
	fmt.Println(somePerson)
	fmt.Println("That's All Folks!!")
}

func showPointer() {
	someVar := "This is a var"
	fmt.Printf("%s, at address %v\n", someVar, &someVar)
}

func changeMe(p *person) {
	p.age = 53
	p.first = "Joris"
	p.last = "Melchior"
}
