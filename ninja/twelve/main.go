package main

import (
	"fmt"

	"github.com/jmelchio/gcgo/ninja/twelve/dog"
)

func main() {
	dogYears := 5
	humanYears := dog.Years(dogYears)
	fmt.Printf("A dog of %d years old is like a human that is %d years old\n", dogYears, humanYears)
}
