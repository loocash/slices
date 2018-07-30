package slices

import (
	"fmt"
	"testing"
)

func ExamplePartition() {
	xs := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	pivot := Partition(xs, 5, func(i, j int) bool {
		return xs[i] < xs[j]
	})
	fmt.Println(pivot, 4)
}

func getLess(xs []int) lessFunc {
	return func(i, j int) bool {
		return xs[i] < xs[j]
	}
}

func TestPartition(t *testing.T) {
	var tests = []struct {
		in   []int
		k    int
		want int
	}{
		{[]int{5, 4, 3, 2, 1}, 0, 4},
		{[]int{5, 4, 3, 2, 1}, 1, 3},
		{[]int{5, 4, 3, 2, 1}, 2, 2},
		{[]int{5, 4, 3, 2, 1}, 3, 1},
		{[]int{5, 4, 3, 2, 1}, 4, 0},
	}

	for _, test := range tests {
		if got := Partition(test.in, test.k, getLess(test.in)); got != test.want {
			t.Errorf("Partition(%v, %v): got: %v, want: %v\n", test.in, test.k, got, test.want)
		}
	}
}
