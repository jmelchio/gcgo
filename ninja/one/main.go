package main

import (
	"fmt"
)

var a int
var b string
var c bool

type myInt int

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	abc()
	def()
	myFunc()
}

func abc() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func def() {
	x := 42
	y := "James Blond"
	z := false

	s := fmt.Sprintf("%d, %s, %v", x, y, z)
	fmt.Println(s)
}

func myFunc() {
	var x myInt
	var y int

	fmt.Printf("var x: %v, %T\n", x, x)

	x = 42
	y = int(x)

	fmt.Printf("var x: %v\n", x)
	fmt.Printf("var y: %v, %T\n", y, y)
}
