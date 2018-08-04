package slices

import (
	"testing"
)

func getEquals(xs []int) equalsFunc {
	var eq = func(i, j int) bool {
		return xs[i] == xs[j]
	}
	return eq
}

func getEquals2(xs, ys []int) equalsFunc {
	var eq = func(i, j int) bool {
		return xs[i] == ys[j]
	}
	return eq
}

func TestUniq(t *testing.T) {
	var tests = []struct {
		before, after []int
	}{
		{[]int{1, 1, 1, 1, 1}, []int{1}},
		{[]int{1, 1, 2, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1, 1, 2}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		got := Uniq(test.before, getEquals(test.before))
		assertSlice(t, test.after, got)
	}
}

func TestEquals(t *testing.T) {
	var tests = []struct {
		a, b []int
		want bool
	}{
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}, true},
		{[]int{1, 1, 2, 2, 3}, []int{1, 2, 3}, false},
		{[]int{1, 1, 1, 1, 2}, []int{1, 1, 2, 1, 2}, false},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, true},
	}

	for _, test := range tests {
		got := Equals(test.a, test.b, getEquals2(test.a, test.b))
		if got != test.want {
			t.Errorf("Equals(%v, %v): want: %v, got: %v\n", test.a, test.b, test.want, got)
		}
	}
}
