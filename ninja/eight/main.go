package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	simpleMarshal()
	fmt.Println("That's all for Ninja level eight folks !!")
}

func simpleMarshal() {
	x := "someString"
	b, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err)
	} else {
		os.Stdout.Write(b)
		os.Stdout.WriteString("\n")
	}

}
