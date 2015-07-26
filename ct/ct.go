// Package ct provides a wrapper around time.Ticker that closes the channel on Stop.
package ct // import "merovius.de/go-misc/ct"

import "time"

// A Ticker holds a channel that delivers `ticks' of a clock at intervals.
type Ticker struct {
	c    chan time.Time
	done chan bool
	t    *time.Ticker
}

// NewTicker returns a new Ticker containing a channel that will send the time
// with a period specified by the duration argument. It adjusts the intervals
// or drops ticks to make up for slow receivers. The duration d must be greater
// than zero; if not, NewTicker will panic. Stop the ticker to release
// associated resources.
func NewTicker(d time.Duration) *Ticker {
	t := &Ticker{
		c:    make(chan time.Time),
		done: make(chan bool),
		t:    time.NewTicker(d),
	}
	go func() {
		for {
			select {
			case ti := <-t.t.C:
				t.c <- ti
			case <-t.done:
				close(t.c)
				return
			}
		}
	}()
	return t
}

// C returns the channel where ticks are delivered.
func (t *Ticker) C() <-chan time.Time {
	return t.c
}

// Stop turns off a ticker. After Stop, no more ticks will be sent. Stop closes
// the channel, so you can range over it.
func (t *Ticker) Stop() {
	t.t.Stop()
	close(t.done)
}
