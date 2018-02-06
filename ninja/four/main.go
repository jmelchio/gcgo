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
		fmt.Printf("The age of %s is: %d", k, v)
	}
}

// That's All Folks !!
