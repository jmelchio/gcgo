package main

import "fmt"

func main() {
	fmt.Printf("Factorial %d = %d\n", 5, factorial(5))
	fmt.Printf("Factorial loop %d = %d\n", 5, factorialLoop(5))

	fmt.Printf("The answer to life and everything: %d\n", foo())
	answer, question := bar()
	fmt.Printf("%s: %d\n", question, answer)

	numbers := []int{2, 3, 5, 6, 8, 9}
	fmt.Printf("Using foo2 the sum of %v is: %d\n", numbers, foo2(numbers...))
	fmt.Printf("Using bar2 the sum of %v is: %d\n", numbers, bar2(numbers))

	deferit()

	fmt.Println("\nThat's all folks !!")
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

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "How much is seven times six?"
}

func foo2(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func bar2(numbers []int) int {
	return foo2(numbers...)
}

func deferit() {
	defer defered()
	fmt.Println("This should show up before defered ....")
}

func defered() {
	fmt.Println("This is defered ...")
}
