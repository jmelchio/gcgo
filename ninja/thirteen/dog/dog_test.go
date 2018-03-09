package dog

import (
	"fmt"
	"testing"
)

func TestHumanToDogYears(t *testing.T) {
	n := HumanToDogYears(5)
	if n != 35 {
		t.Error("want:", 35, "got:", n)
	}
}

func TestHumanToDogYearsTwo(t *testing.T) {
	n := HumanToDogYearsTwo(5)
	if n != 35 {
		t.Error("want:", 35, "got:", n)
	}
}

func ExampleHumanToDogYears() {
	fmt.Println(HumanToDogYears(10))
	// Output:
	// 70
}

func ExampleHumanToDogYearsTwo() {
	fmt.Println(HumanToDogYearsTwo(10))
	// Output:
	// 70
}

func BenchmarkHumanToDogYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HumanToDogYears(5)
	}
}

func BenchmarkHumanToDogYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HumanToDogYearsTwo(5)
	}
}
