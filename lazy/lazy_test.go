package lazy

import "testing"

func TestBool(t *testing.T) {
	for _, expect := range []bool{true, false} {
		called := false
		b := Bool(func() bool {
			if called {
				t.Errorf("Bool func evaluated twice")
			}
			called = true
			return expect
		})

		if got := b(); got != expect {
			t.Errorf("b() == %v, expected %v", got, expect)
		}
		if got := b(); got != expect {
			t.Errorf("b() == %v, expected %v", got, expect)
		}
	}
}
