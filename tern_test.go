package handy

import "testing"

func TestTern(t *testing.T) {
	t.Run("truthy", func(t *testing.T) {
		a, b := "foo", "bar"
		out := Tern(true, a, b)
		if out != a {
			t.Errorf("Tern(true, %s, %s) = %s, want %s", a, b, out, a)
		}
	})

	t.Run("falsey", func(t *testing.T) {
		a, b := "foo", "bar"
		out := Tern(false, a, b)
		if out != b {
			t.Errorf("Tern(false, %s, %s) = %s, want %s", a, b, out, b)
		}
	})
}

func TestLazyTern(t *testing.T) {
	t.Run("truthy", func(t *testing.T) {
		a, b := "foo", "bar"
		out := LazyTern(
			true,
			func() string { return a },
			func() string { return b },
		)
		if out != a {
			t.Errorf("Tern(true, %s, %s) = %s, want %s", a, b, out, a)
		}
	})

	t.Run("falsey", func(t *testing.T) {
		a, b := "foo", "bar"
		out := LazyTern(
			false,
			func() string { return a },
			func() string { return b },
		)
		if out != b {
			t.Errorf("Tern(false, %s, %s) = %s, want %s", a, b, out, b)
		}
	})
}
