package slices

import (
	"fmt"
	"testing"
)

func ExampleKthElement() {
	elems := []int{8, 3, 6, 9, 2, 1, 4, 0, 5, 7}
	k := 4
	KthElement(elems, k, func(i, j int) bool {
		return elems[i] < elems[j]
	})
	fmt.Println(elems[k])
	// Output:
	// 4
}

func TestKthElement(t *testing.T) {
	var tests = []struct {
		in   []int
		k    int
		want int
	}{
		{[]int{5, 4, 3, 2, 1}, 0, 1},
		{[]int{5, 4, 3, 2, 1}, 1, 2},
		{[]int{5, 4, 3, 2, 1}, 2, 3},
		{[]int{5, 4, 3, 2, 1}, 3, 4},
		{[]int{5, 4, 3, 2, 1}, 4, 5},
	}

	for _, test := range tests {
		KthElement(test.in, test.k, getLess(test.in))
		if got := test.in[test.k]; got != test.want {
			t.Errorf("KthElement(%v, %v): got: %v, want: %v\n", test.in, test.k, got, test.want)
		}
	}
}
