package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	str := "You are doing great"
	ans := 4
	out := Count(str)
	if ans != out {
		t.Errorf("Expected %v, got %v", ans, out)
	}
}

func ExampleCount() {
	fmt.Println(Count("You are doing great"))
	// Output:
	// 4
}

func BenchmarkCount(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Count("You are doing great")
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		}
	}
}
