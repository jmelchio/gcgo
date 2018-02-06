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
	exerciseTwo()
}

func exerciseOne() {
	intArray := [5]int{1, 3, 5, 7, 9}

	for i, v := range intArray {
		fmt.Printf("Position: %d - Value: %d, Type: %T\n", i, v, v)
	}
	fmt.Printf("Type of variable is: %T\n", intArray)
}

func exerciseTwo() {
	intSlice := []int{12, 13, 14, 16, 28, 17, 38, 56, 23, 77}

	for i, v := range intSlice {
		fmt.Printf("Position: %d, Value: %d, Type: %T\n", i, v, v)
	}
	fmt.Printf("Type of variable is: %T\n", intSlice)

	fmt.Println(intSlice)
	fmt.Println(intSlice[3:8])
	fmt.Println(intSlice[5:])
	fmt.Println(intSlice[:5])
	fmt.Println(intSlice[2:7])
	fmt.Println(intSlice[1:6])
}

// That's All Folks !!
