package owned

import "fmt"

func Example() {
	ch := New(nil)

	// Set an arbitrary value.
	ch <- func(_ interface{}) interface{} {
		return 42
	}

	// Print current value.
	ch <- func(v interface{}) interface{} {
		fmt.Println("Current value:", v)
		return v
	}

	// Atomic compare and swap
	swapped := make(chan bool)
	ch <- func(v interface{}) interface{} {
		if v.(int) == 42 {
			swapped <- true
			return 23
		}
		swapped <- false
		return v
	}
	if <-swapped {
		fmt.Println("Swapped")
	} else {
		fmt.Println("Didn't swap")
	}

	// Output:
	// Current value: 42
	// Swapped
}
