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
	exerciseFour()
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

func exerciseFour() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

// That's All Folks !!
