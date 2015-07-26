package toggle_test

import (
	"fmt"

	"merovius.de/go-misc/toggle"
)

func Example() {
	r, w := toggle.First()

	select {
	case v := <-r:
		fmt.Println("Read", v)
	default:
		fmt.Println("Nothing to read")
	}

	w <- "First"

	s, ok := <-r
	fmt.Println(s, ok)

	w <- "Second"
	w <- "Third"

	fmt.Println(<-r)

	w <- "Fourth"
	close(w)
	fmt.Println(<-r)

	// Output:
	// Nothing to read
	// First true
	// Second
	// Fourth
}
