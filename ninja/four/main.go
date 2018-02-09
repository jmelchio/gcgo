package main

import (
	"fmt"
)

func main() {
	fmt.Println("What's up Jackie?")
	mymap := map[string]int{
		"Joris":  53,
		"Andrea": 52,
		"Jasper": 18,
		"Alanna": 15,
	}

	for k, v := range mymap {
		fmt.Printf("The age of %s is: %d\n", k, v)
	}

	exerciseOne()
	exerciseTwo()
	exerciseFour()
	exerciseFive()
	exerciseSix()
	exerciseSeven()
	exerciseEight()
}

func exerciseOne() {
	intArray := [5]int{1, 3, 5, 7, 9}

	for i, v := range intArray {
		fmt.Printf("Position: %d - Value: %d, Type: %T\n", i, v, v)
	}
	fmt.Printf("Type of variable is: %T\n", intArray)
}

func exerciseTwo() {
	intSlice := []int{12, 13, 14, 16, 28, 17, 38, 56, 23, 77}

	for i, v := range intSlice {
		fmt.Printf("Position: %d, Value: %d, Type: %T\n", i, v, v)
	}
	fmt.Printf("Type of variable is: %T\n", intSlice)

	fmt.Println(intSlice)
	fmt.Println(intSlice[3:8])
	fmt.Println(intSlice[5:])
	fmt.Println(intSlice[:5])
	fmt.Println(intSlice[2:7])
	fmt.Println(intSlice[1:6])
}

func exerciseFour() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func exerciseFive() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func exerciseSix() {
	// not really according to exercise but it's a good experiment
	x := make([]string, 50, 55)
	fmt.Printf("Slice capacity is %d and length is %d\n", cap(x), len(x))
	x = append(x, `Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`,
		`Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`,
		`Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`,
		`Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`,
		`Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`,
		`North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`,
		`Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`,
		`Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`)

	for i := 0; i < len(x); i++ {
		fmt.Printf("State %d is %s\n", i, x[i])
	}
	fmt.Printf("Slice capacity is %d and length is %d\n", cap(x), len(x))
}

func exerciseSeven() {
	x := [][]string{{"James", "Bond", "Shaken not stirred"}, {"Miss", "Monypenny", "Hellooo, James"}}

	for i, ss := range x {
		fmt.Println("record: ", i)
		for j, s := range ss {
			fmt.Printf("Index %d, is Value: %s\n", j, s)
		}
	}
}

func exerciseEight() {
	x := make(map[string][]string)
	x["bond_james"] = []string{
		"Shaken, not stirred",
		"Martinis",
		"Women"}
	x["moneypenny_miss"] = []string{
		"James Bond",
		"Literature",
		"Computer Science",
	}
	x["no_dr"] = []string{
		"Being evil",
		"Ice cream",
		"Sunsets",
	}
	x["M_MI5"] = []string{
		"Gadgets",
		"Lab",
		"Explosive stuff",
	}

	delete(x, "no_dr")

	for k, v := range x {
		fmt.Printf("Favorite things for: %s\n", k)
		for i, s := range v {
			fmt.Printf("\tItem %d: %s\n", i, s)
		}
	}

}

// That's All Folks !!
