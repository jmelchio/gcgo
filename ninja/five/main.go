package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favoriteIcFlavors []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	exerciseOne()
	exerciseThree()
	exerciseFour()

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

func exerciseThree() {
	aTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	aSedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Printf("A vehicle of type %T, with %d doors %s color and fourWheelDrive: %v\n",
		aTruck, aTruck.doors, aTruck.color, aTruck.fourWheel)
	fmt.Printf("A vehicle of type %T, with %d doors %s color and luxury equipped: %v\n",
		aSedan, aSedan.doors, aSedan.color, aSedan.luxury)

}

func exerciseFour() {
	anonymous := struct {
		somestuff  string
		otherstuff string
		really     bool
	}{
		somestuff:  "some",
		otherstuff: "other",
		really:     true,
	}

	fmt.Printf("Sometimes %s is going on\nand then other times it's %s\n\t it's really %v\n",
		anonymous.somestuff, anonymous.otherstuff, anonymous.really)
}
