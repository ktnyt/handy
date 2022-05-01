package handy

import "testing"

func TestNew(t *testing.T) {
	i := 42
	p := New(i)
	if *p != i {
		t.Errorf("*New(%d) = %d, want 42", i, *p)
	}

	if p == &i {
		t.Errorf("New(%d) = %p, want not %p", i, p, &i)
	}
}
