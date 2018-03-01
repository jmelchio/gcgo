package main

import (
	"fmt"
	"sync"
)

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

	sc := make(chan int)
	go sendSix(sc, 25)
	receiveSix(sc)

	sc2 := make(chan int)
	go receiveSix(sc2)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		fmt.Println("Kicking of routine:", i)
		wg.Add(1)
		go sendSeven(sc2, 5, &wg)
	}
	wg.Wait()
	fmt.Println("WaitGroup is done, closing channel")
	close(sc2)

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

func sendSix(sc chan<- int, repeat int) {
	for i := 0; i < repeat; i++ {
		sc <- i
	}
	close(sc)
}

func receiveSix(rc <-chan int) {
	// for {
	// 	select {
	// 	case v, ok := <-rc:
	// 		if !ok {
	// 			fmt.Println("Received close signal")
	// 			return
	// 		}
	// 		fmt.Println("Received value:", v)
	// 	}
	// }
	// simpler ...
	for v := range rc {
		fmt.Println("Received value:", v)
	}
	fmt.Println("Done printing...")
}

func sendSeven(sc chan<- int, repeat int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < repeat; i++ {
		sc <- i
	}
}

// That's All Folks !!
