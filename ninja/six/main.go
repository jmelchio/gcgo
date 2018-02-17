package main

import (
	"fmt"
	"math"
)

type person struct {
	first string
	last  string
	age   int
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type shape interface {
	area() float64
}

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

	nobody := person{
		first: "John",
		last:  "Doe",
		age:   42,
	}
	fmt.Println(nobody.speak())

	rondje := circle{radius: 20.0}
	fmt.Printf("The surface area of the circle with r %.2f is : %.2f\n",
		rondje.radius, info(rondje))
	vierkant := square{side: 10.0}
	fmt.Printf("The surface area of the square with length/width %.2f is : %.2f\n",
		vierkant.side, info(vierkant))

	defer func(x int) {
		fmt.Printf("Ever wonder what the answer to life and everything is? It's %d\n", x)
	}(42)

	funky := func(x int) int {
		return x
	}

	otherfunc := somefunc()
	someInt := 10
	fmt.Printf("This func %T returns the %T you pass into it. For example: %d\n", otherfunc, someInt, otherfunc(someInt))

	fmt.Println("It's", funky(42))

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
	fmt.Println("This should show up before deferred ....")
}

func defered() {
	fmt.Println("This is deferred ...")
}

func (p *person) speak() string {
	return fmt.Sprintf("Hi, I am %s %s and I'm %d years old.\n", p.first, p.last, p.age)
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2.0)
}

func info(s shape) float64 {
	return s.area()
}

func somefunc() func(int) int {
	return func(x int) int {
		return x
	}
}
