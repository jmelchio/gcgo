// Package dog helps us understand the age of dogs
package dog

// HumanToDogYears calculates how many dog years you have for each human year
func HumanToDogYears(humanYears int) int {
	return humanYears * 7
}

// HumanToDogYearsTwo calculates how many dog years there are for each human years
// but in a slightly awkward fashion
func HumanToDogYearsTwo(humanYears int) int {
	count := 0
	for i := 0; i < humanYears; i++ {
		count += 7
	}
	return count
}
