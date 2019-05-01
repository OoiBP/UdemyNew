package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type data struct {
		input  []int
		output float64
	}

	xdata := []data{
		data{
			input:  []int{2, 3, 4},
			output: 3,
		},
		data{
			input:  []int{2, 3, 4, 5, 6},
			output: 4,
		},
		data{
			input:  []int{10, 20, 30, 50},
			output: 25,
		},
	}

	for _, v := range xdata {
		ans := CenteredAvg(v.input)
		if ans != v.output {
			t.Error("Expected", v.output, "got", ans)
		}
	}
}

func ExampleCenteredAvg() {
	xi := []int{2, 3, 4, 5, 6}
	fmt.Println(CenteredAvg(xi))
	// Output:
	// 4
}

func BenchmarkCenteredAvg(b *testing.B) {
	for index := 0; index < b.N; index++ {
		xi := []int{2, 3, 4, 5, 6}
		CenteredAvg(xi)
	}
}
