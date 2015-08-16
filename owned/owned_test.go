package owned

import "fmt"

func Example() {
	ch := New(nil)

	// Set an arbitrary value.
	ch.Set(42)
	// Equivalent:
	ch <- func(_ interface{}) interface{} {
		return 42
	}

	// Print current value.
	fmt.Println("Current value:", ch.Get())
	// Equivalent:
	ch <- func(v interface{}) interface{} {
		fmt.Println("Current value:", v)
		return v
	}

	// Atomic compare and swap
	if ch.CAS(42, 23) {
		fmt.Println("Swapped")
	} else {
		fmt.Println("Didn't swap")
	}
	// Equivalent (but doesn't swap, as the comparison fails):
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
	// Current value: 42
	// Swapped
	// Didn't swap
}
