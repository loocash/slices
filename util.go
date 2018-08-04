package slices

import "reflect"

type equalsFunc func(i, j int) bool

// Uniq returns a slice with adjacent duplicates removed
func Uniq(slice interface{}, eq equalsFunc) interface{} {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()

	last := 0
	for i := 1; i < length; i++ {
		if !eq(i, last) {
			last++
			swap(last, i)
		}
	}

	return rv.Slice(0, last+1)
}

// Equals return true if two slices have the same length
// and corresponding elements are equal to each other
func Equals(a, b interface{}, eq equalsFunc) bool {
	ra := reflect.ValueOf(a)
	rb := reflect.ValueOf(b)
	alen := ra.Len()
	blen := rb.Len()
	if alen != blen {
		return false
	}
	for i := 0; i < alen; i++ {
		if !eq(i, i) {
			return false
		}
	}
	return true
}
