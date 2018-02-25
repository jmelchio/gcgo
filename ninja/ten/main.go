package main

import "fmt"

func main() {
	simpleChannel()
	bufferedChannel()
	exerciseTwoOne()

	c := gen()
	receive(c)

	fmt.Println("about to exit")

	fmt.Println("That's all for Ninja level ten folks !!")
}

func simpleChannel() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println("Received", <-c, "from simpleChannel")
}

func bufferedChannel() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println("Received", <-c, "from bufferedChannel")
}

func exerciseTwoOne() {
	c := make(chan int)

	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	fmt.Printf("------\n")
	fmt.Printf("c\t%T\n", c)
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for i := range c {
		fmt.Println("Int from channel:", i)
	}
}

// That's All Folks !!
