package word

import (
	"fmt"
	"testing"

	"github.com/jmelchio/gcgo/ninja/thirteen/quote"
)

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("three simple words"))
	// Output:
	// 3
}

func TestUseCount(t *testing.T) {
	wordgroup := "some of of these words words are double"
	wordmap := UseCount(wordgroup)
	for k, v := range wordmap {
		switch k {
		case "some":
			if v != 1 {
				t.Error("Got:", v, "Wanted:", 1)
			}
		case "of":
			if v != 2 {
				t.Error("Got:", v, "Wanted:", 1)
			}
		case "these":
			if v != 1 {
				t.Error("Got:", v, "Wanted:", 1)
			}
		case "words":
			if v != 2 {
				t.Error("Got:", v, "Wanted:", 1)
			}
		case "are":
			if v != 1 {
				t.Error("Got:", v, "Wanted:", 1)
			}
		case "double":
			if v != 1 {
				t.Error("Got:", v, "Wanted:", 1)
			}
		}
	}
}

func TestCount(t *testing.T) {
	somewords := "here are six very interesting words"
	n := Count(somewords)
	if n != 6 {
		t.Error("got:", n, "want:", 6)
	}
}
