package main

import (
	"fmt"

	"github.com/jmelchio/gcgo/ninja/thirteen/dog"
	"github.com/jmelchio/gcgo/ninja/thirteen/quote"
	"github.com/jmelchio/gcgo/ninja/thirteen/word"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  10,
	}
	fmt.Printf("The dog %s is %d dog years and %d in human years\n", fido.name, fido.age, dog.HumanToDogYears(fido.age))

	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
	fmt.Println("That's all for Ninja level thirteen folks !!")
}
