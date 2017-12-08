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
	thisorthat("", "that")
	fmt.Println(simpleswitch("I don't know"))
	fmt.Println(stringswitch("sumtin"))
	truethat()
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

// exercise #6 and #7
func thisorthat(this, that string) {
	if len(this) == 0 && len(that) == 0 {
		fmt.Println("neither this nor that")
	} else if len(this) == 0 && len(that) != 0 {
		fmt.Println("not this but that")
	} else if len(this) != 0 && len(that) == 0 {
		fmt.Println("this but not that")
	} else {
		fmt.Println("this and that")
	}
}

// exercise #8
func simpleswitch(word string) string {
	switch {
	case word == "whoot":
		return "why whoot?"
	case word == "word":
		return "word dude!"
	case len(word) == 0:
		return "sup quiet one?"
	default:
		return "what you said man!"
	}
}

// exercise #9
func stringswitch(identifier string) string {
	switch identifier {
	case "word":
		return "word up"
	case "no way":
		return "why not?"
	default:
		return "whatever ..."
	}
}

// exercise #10
func truethat() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
}

// That's All Folks !!
