package main

import (
	"fmt"

	"github.com/jmelchio/gcgo/ninja/thirteen/dog"
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
	fmt.Printf("The dog %s is %d dog years and %d in human years", fido.name, fido.age, dog.HumanToDogYears(fido.age))
	fmt.Println("That's all for Ninja level thirteen folks !!")
}
