package dog

import (
	"fmt"
	"testing"
)

type data struct {
	input  int
	output int
}

func TestYears(t *testing.T) {
	xdata := []data{
		data{1, 7},
		data{2, 14},
		data{5, 35},
	}

	for _, v := range xdata {
		ans := Years(v.input)
		if ans != v.output {
			t.Errorf("Expected %v, got %v", v.output, ans)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	xdata := []data{
		data{1, 7},
		data{2, 14},
		data{5, 35},
	}

	for _, v := range xdata {
		ans := YearsTwo(v.input)
		if ans != v.output {
			t.Errorf("Expected %v, got %v", v.output, ans)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(2))
	// Output:
	// 14
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(2))
	// Output:
	// 14
}

func BenchmarkYears(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for index := 0; index < b.N; index++ {
		YearsTwo(10)
	}
}
