package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favoriteIcFlavors []string
}

func main() {
	exerciseOne()

	fmt.Println("That's all folks!!")
}

func exerciseOne() {
	joris := person{
		firstName:         "Joris",
		lastName:          "Melchior",
		favoriteIcFlavors: []string{"strawberry", "blackberry", "mango"},
	}

	andrea := person{
		firstName:         "Andrea",
		lastName:          "Melchior",
		favoriteIcFlavors: []string{"chocolate", "pistachio", "mango"},
	}

	people := map[string]person{"joris": joris, "andrea": andrea}

	for _, v := range people {
		fmt.Printf("%s %s likes icecream flavors: \n", v.firstName, v.lastName)
		for _, flavor := range v.favoriteIcFlavors {
			fmt.Printf("\t%s\n", flavor)
		}
	}
}
