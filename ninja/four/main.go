package main

import (
	"fmt"
)

func main() {
	fmt.Println("What's up Jackie?")
	mymap := map[string]int{
		"Joris":  53,
		"Andrea": 52,
		"Jasper": 18,
		"Alanna": 15,
	}

	for k, v := range mymap {
		fmt.Printf("The age of %s is: %d\n", k, v)
	}

	exerciseOne()
}

func exerciseOne() {
	intArray := [5]int{1, 3, 5, 7, 9}

	for i, v := range intArray {
		fmt.Printf("Position: %d - Value: %d, Type: %T\n", i, v, v)
	}
	fmt.Printf("Type of variable is: %T\n", intArray)
}

// That's All Folks !!
