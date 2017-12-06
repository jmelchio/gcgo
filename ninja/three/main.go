package main

import (
	"fmt"
)

func main() {
	printem(10001)
	printalfarange(65, 90)
	printalive(1964)
	printaliveagain(1964)
	printremainder(10, 100, 4)
}

// exercise #1
func printem(max int) {
	for i := 0; i < max; i++ {
		fmt.Println(i)
	}
}

// exercise #2
func printalfarange(start, end int) {
	for i := start; i <= end; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

// exercise #3
func printalive(birthyear int) {
	i := 2017
	for i >= birthyear {
		fmt.Println("I'm alive!: ", birthyear)
		birthyear++
	}
}

// exercise #4
func printaliveagain(birthyear int) {
	i := 2017
	for {
		if i < birthyear {
			break
		}
		fmt.Println("I'm alive!: ", birthyear)
		birthyear++
	}
}

// exercise $5
func printremainder(from, to, divisor int) {
	for i := from; i <= to; i++ {
		fmt.Printf("%d mod %d = %d\n", i, divisor, i%divisor)
	}
}

// That's All Folks !!
