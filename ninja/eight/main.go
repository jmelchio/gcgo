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

type Address struct {
	StreetOne  string `json:"streetone"`
	StreetTwo  string `json:"streettwo,omitempty"`
	PostalCode string `json:"postalcode,omitempty"`
	City       string `json:"city"`
	Province   string `json:"province,omitempty"`
	Country    string `json:"country,omitempty"`
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
	marshalPerson()
	unmarshalToStruct()
	simpleSort()
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

func marshalPerson() {
	p1 := Person{
		First: "John",
		Last:  "Doe",
		Age:   19,
	}

	p2 := Person{
		First: "Jane",
		Last:  "Doe",
		Age:   21,
	}

	jp, err := json.Marshal([]Person{p1, p2})
	if err != nil {
		fmt.Println("Marshal person did not work:", err)
		return
	}
	fmt.Println(string(jp))
}

func unmarshalToStruct() {
	address := `{"streetone": "51 Braywin Drive", "city": "Etobicoke", "province": "Ontario", "country": "Canada"}`
	fmt.Println(address)
	aba := []byte(address)
	anAddress := Address{}
	err := json.Unmarshal(aba, &anAddress)
	if err != nil {
		fmt.Println("Unable to unmarshal: ", err)
		return
	}
	fmt.Println("\t decode it ...")
	fmt.Println(anAddress)
	fmt.Println("\t and encode back again ...")
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(anAddress)
}

func simpleSort() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
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

	fmt.Println("\t unsorted ...")
	fmt.Println(ps)
	sort.Sort(ByFirst(ps))
	fmt.Println("\t sorted by first name ...")
	fmt.Println(ps)
	fmt.Println("\t sorted by last name ...")
	sort.Sort(ByLast(ps))
	fmt.Println(ps)
}
