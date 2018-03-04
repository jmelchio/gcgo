package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

type customErr struct {
	Stuff          string
	Morestuff      string
	Evenmorestuff  string
	Erroneousstuff string
}

func main() {
	exerciseOne()
	ce := customErr{
		Stuff:          "stuffy",
		Morestuff:      "stuffier",
		Evenmorestuff:  "stuff to the max",
		Erroneousstuff: "and dead wrong too",
	}

	err := exerciseThree(ce)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("That's all for Ninja eleven folks !!")
}

func (es customErr) Error() string {
	return fmt.Sprintf("This is extremely wrong %s because %s and %s and also %s", es.Erroneousstuff, es.Stuff, es.Morestuff, es.Evenmorestuff)
}
func exerciseOne() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}

func exerciseThree(ce customErr) error {
	if ce.Erroneousstuff != "" {
		return ce
	}
	return nil
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("Some error occurred converting %v into JSON: %v", a, err)
	}
	return bs, nil
}
