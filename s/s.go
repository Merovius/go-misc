// Package s provides a closure-based wrapper around sort.Interface
//
// SortStruct wraps closures that implement sort.Interface in a struct. That
// way, there is no need to define top-level types for each []T in your program
// and you save a tiny amount of boiler plate. It's probably the most
// convenient implementation of sorting that doesn't try to circumvent the type
// system. It is a completely trivial package and neither it's implementation
// nor it's interface will ever change.
package s

// SortStruct implements sort.Interface by wrapping func values.
type SortStruct struct {
	// Length is the number of elements in the collection.
	Length int

	// SwapF swaps the elements with indixes i and j.
	SwapF func(i, j int)

	// LessF reports whether the element with index i should sort before the
	// element with index j.
	LessF func(i, j int) bool
}

// Len implements sort.Interface.
func (s SortStruct) Len() int {
	return s.Length
}

// Swap implements sort.Interface.
func (s SortStruct) Swap(i, j int) {
	s.SwapF(i, j)
}

// Less implements sort.Interface.
func (s SortStruct) Less(i, j int) bool {
	return s.LessF(i, j)
}
