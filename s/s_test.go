package s_test

import (
	"fmt"
	"sort"

	"merovius.de/go-misc/s"
)

func ExampleSortStruct() {
	a := []int{2, 4, 42, 23, 1337, 1}

	sort.Sort(s.SortStruct{
		Length: len(a),
		SwapF:  func(i, j int) { a[i], a[j] = a[j], a[i] },
		LessF:  func(i, j int) bool { return a[i] < a[j] },
	})

	fmt.Println(a)

	// Output: [1 2 4 23 42 1337]
}
