package main

import "fmt"

func main() {
	simpleChannel()
	bufferedChannel()
	exerciseTwoOne()

	c := gen()
	receive(c)

	fmt.Println("about to exit one")

	q := make(chan int)
	c2 := genFour(q)
	receiveFour(c2, q)

	fmt.Println("about to exit two")

	showOkayIdiom()

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

func genFour(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receiveFour(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func showOkayIdiom() {
	c := make(chan int)

	go func() {
		c <- 5
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

// That's All Folks !!
