// Package toggle implements condition variables usable in selects
//
// A toggle consists of two channels, r and w. A read from r will succeed iff
// there was a send to w since the last read. A send to w will never block.
// When w is closed, r is closed and all associated resources are released,
// but only after a potential value written to w was read.
package toggle // import "merovius.de/go-misc/toggle"

func create(last bool) (chan interface{}, chan interface{}) {
	r, w := make(chan interface{}), make(chan interface{})

	go func() {
		var v interface{}
		var ok bool
		for {
			v, ok = <-w
			if !ok {
				w = nil
			}
		L:
			for {
				select {
				case V, ok := <-w:
					if last {
						v = V
					}
					if !ok {
						w = nil
					}
				case r <- v:
					v = nil
					break L
				}
			}

			if w == nil {
				close(r)
				return
			}
		}
	}()

	return r, w
}

// First returns a new toggle, that yields the first value written to w since
// the last read.
func First() (r <-chan interface{}, w chan<- interface{}) {
	return create(false)
}

// Last returns a new toggle, that yields the last value written to w since the
// last read.
func Last() (r <-chan interface{}, w chan<- interface{}) {
	return create(true)
}
