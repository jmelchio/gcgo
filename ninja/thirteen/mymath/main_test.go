package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	// Output:
	// 5
}

func TestCenteredAvg(t *testing.T) {
	type tdata struct {
		data   []int
		answer float64
	}

	testing := []tdata{
		tdata{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, answer: 5.0},
		tdata{data: []int{10, 20, 30, 40, 50, 60, 70, 80, 90}, answer: 50.0},
		tdata{data: []int{5, 8, 10, 12, 15, 18, 20}, answer: 12.6},
	}

	for _, v := range testing {
		n := CenteredAvg(v.data)
		if n != v.answer {
			t.Error("Got:", n, "Want:", v.answer)
		}
	}
}
