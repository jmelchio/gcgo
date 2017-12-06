package main

import (
	"fmt"
)

const (
	a       = 42
	b int32 = 80
	c       = "ce"

	year1 = 2012 + iota
	year2 = 2012 + iota
	year3 = 2012 + iota
	year4 = 2012 + iota
)

func main() {
	number := 25
	fmt.Printf("number %d, %b, %#x\n", number, number, number)

	anotherNumber := 42
	fmt.Printf("anotherNumber %d, %b, %#x\n", anotherNumber, anotherNumber, anotherNumber)

	equals := number == anotherNumber
	lesEqaual := number <= anotherNumber
	greaterEqual := number >= anotherNumber
	notEqual := number != anotherNumber
	less := number < anotherNumber
	greater := number > anotherNumber

	fmt.Println("number equals anotherNumber? : ", equals)
	fmt.Println("number lesEqual anotherNumber? : ", lesEqaual)
	fmt.Println("number greaterEqual anotherNumber? : ", greaterEqual)
	fmt.Println("number notEqual anotherNumber? : ", notEqual)
	fmt.Println("number less anotherNumber? : ", less)
	fmt.Println("number greater anotherNumber? : ", greater)

	fmt.Printf("const %v %T\n", a, a)
	fmt.Printf("const %v %T\n", b, b)
	fmt.Printf("const %v %T\n", c, c)

	someNumber := 42
	fmt.Printf("someNumber %d %b %#x\n", someNumber, someNumber, someNumber)

	someShift := someNumber << 1
	fmt.Printf("someShift %d %b %#x\n", someShift, someShift, someShift)

	secondShift := someNumber >> 1
	fmt.Printf("secondShift %d %b %#x\n", secondShift, secondShift, secondShift)

	rawString := `And now for something completely different:
    "The Larch"`
	fmt.Println(rawString)

	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)

}
