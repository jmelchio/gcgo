package main

import (
	"fmt"

	"github.com/jmelchio/gcgo/ninja/thirteen/dog"
	"github.com/jmelchio/gcgo/ninja/thirteen/mymath"
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

	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}

	fmt.Println("That's all for Ninja level thirteen folks !!")
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
