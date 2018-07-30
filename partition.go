package slices

import "reflect"

type lessFunc func(i, j int) bool

func partition(slice interface{}, k, a, b int, less lessFunc) int {
	swap := reflect.Swapper(slice)

	swap(b-1, k)
	last := -1
	for i := a; i < b; i++ {
		if less(i, b-1) {
			last++
			swap(i, a+last)
		}
	}
	last++
	swap(b-1, a+last)
	return last
}

// Partition rearranges elements in the slice so that
// every element lower that slice[k] is placed before
// and slice[k] is placed right after them.
// Complexity: Time - O(n), Memory - O(1)
func Partition(slice interface{}, k int, less lessFunc) int {
	rv := reflect.ValueOf(slice)
	length := rv.Len()

	return partition(slice, k, 0, length, less)
}
