package slices

import "reflect"

func kthElem(xs interface{}, k int, start int, end int, less lessFunc) int {
	i := (start + end) / 2
	pivot := partition(xs, i, start, end, less)

	if end-start == 1 {
		return start
	}
	if pivot == k {
		return start + k
	}
	if pivot < k {
		return kthElem(xs, k-pivot-1, start+pivot+1, end, less)
	}
	return kthElem(xs, k, start, start+pivot, less)
}

// KthElement finds kth lowest element just like the slice
// was sorted using less function and swaps it to the kth
// place in the slice.
// Complexity: Time - O(n), Memory - O(1)
func KthElement(slice interface{}, k int, less lessFunc) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)

	i := kthElem(slice, k, 0, rv.Len(), less)
	swap(k, i)
}
