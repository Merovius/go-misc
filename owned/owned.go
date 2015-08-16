// Package owned gives synchronized access (via channel) to a value.
//
// The package creates a container (in the form of a closures) that gives you
// exactly one operation: You can pass it a filter (via a channel), that it
// will apply atomically to the held value, i.e. your function executes with an
// exclusive lock on that value. This is a surprisingly powerful concept. See
// the examples (and the code of this package) of how to use this for common
// operations.
//
// The value can be anything, but if you use a pointer, you should never retain
// it outside of the filter-function. Otherwise races are still possible. This
// means, in particular, that you shouldn't use the Get method on
// pointer-values and ignore the return value of Set.
//
// The code heavily uses channels and closures and is thus a bit
// allocation-heavy and probably significantly slower than relying on a mutex.
// It is mainly meant as an example of how communication can elegantly express
// shared-memory patterns.
package owned

// Value represents a value owned by a separate goroutine.
type Value chan<- func(interface{}) interface{}

// New returns an owned value, initialized to v.
func New(v interface{}) Value {
	ch := make(chan func(interface{}) interface{})
	go func() {
		for {
			for f := range ch {
				v = f(v)
			}
			return
		}
	}()
	return ch
}

// Get is a shorthand to atomically load the current value. If the held value
// is of a pointer type, you shouldn't use this method (as only the pointer is
// synchronized, not the value pointed to).
func (ch Value) Get() interface{} {
	ret := make(chan interface{})
	ch <- func(v interface{}) interface{} {
		ret <- v
		close(ret)
		return v
	}
	return <-ret
}

// Set is a shorthand to atomically set the current value. The value before the
// change will be returned. If the held value is of a pointer type, you
// shouldn't use the return value (as only the pointer is synchronized, not the
// value pointed to).
func (ch Value) Set(to interface{}) (old interface{}) {
	ret := make(chan interface{})
	ch <- func(v interface{}) interface{} {
		ret <- v
		close(ret)
		return to
	}
	return <-ret
}

// CAS is a shorthand to atomically compare-and-swap the current value. It will
// return, whether the swap took place.
func (ch Value) CAS(cmp, set interface{}) bool {
	ret := make(chan bool)
	ch <- func(v interface{}) interface{} {
		if v == cmp {
			ret <- true
			return set
		}
		ret <- false
		return v
	}
	return <-ret
}
