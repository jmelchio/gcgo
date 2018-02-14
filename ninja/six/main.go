package main

import "fmt"

func main() {
	fmt.Printf("Factorial %d = %d\n", 5, factorial(5))
	fmt.Printf("Factorial loop %d = %d\n", 5, factorialLoop(5))

	fmt.Println("That's all folks !!")
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialLoop(n int) int {
	result := 1
	for ; n > 0; n-- {
		result *= n
	}
	return result
}
