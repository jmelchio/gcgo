package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type ByFirst []Person

func (n ByFirst) Len() int {
	return len(n)
}

func (n ByFirst) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n ByFirst) Less(i, j int) bool {
	return n[i].First < n[j].First
}

type ByLast []Person

func (n ByLast) Len() int {
	return len(n)
}

func (n ByLast) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n ByLast) Less(i, j int) bool {
	return n[i].Last < n[j].Last
}

func main() {
	simpleMarshal()
	sortDemo()
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

func sortDemo() {
	p1 := Person{
		First: "Joris",
		Last:  "Melchior",
		Age:   53,
	}
	p2 := Person{
		First: "Andrea",
		Last:  "Domize",
		Age:   52,
	}
	p3 := Person{
		First: "Tom",
		Last:  "Boduch",
		Age:   58,
	}
	p4 := Person{
		First: "Hiram",
		Last:  "Lee",
		Age:   48,
	}

	ps := []Person{p1, p2, p3, p4}

	fmt.Println(ps)
	sort.Sort(ByFirst(ps))
	fmt.Println(ps)
	sort.Sort(ByLast(ps))
	fmt.Println(ps)
}
