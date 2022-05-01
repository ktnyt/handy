package enums

import (
	"testing"

	"github.com/ktnyt/handy"
)

func TestIndexFunc(t *testing.T) {
	in := handy.Slice("foo", "bar", "baz")
	for i, v := range in {
		out := IndexFunc(in, func(elem string) bool { return elem == v })
		if out != i {
			t.Errorf("IndexFunc(...) = %d, want %d", out, i)
		}
	}

	out := IndexFunc(in, func(elem string) bool { return false })
	if out >= 0 {
		t.Errorf("IndexFunc(...) = %d, want -1", out)
	}
}

func TestIndex(t *testing.T) {
	in := handy.Slice("foo", "bar", "baz")
	for i, v := range in {
		out := Index(in, v)
		if out != i {
			t.Errorf("Index(...) = %d, want %d", out, i)
		}
	}

	out := Index(in, "foobar")
	if out >= 0 {
		t.Errorf("Index(...) = %d, want -1", out)
	}
}

func TestFind(t *testing.T) {
	in := handy.Slice("foo", "bar", "baz")
	for _, v := range in {
		out := Find(in, func(elem string) bool { return elem == v })
		if out != v {
			t.Errorf("Find(...) = %s, want %s", out, v)
		}
	}

	defer func() {
		t.Helper()
		err := recover()
		if err == nil {
			t.Error("expected function to panic")
		}
	}()
	Find(in, func(elem string) bool { return false })
}

func TestContainsFunc(t *testing.T) {
	in := handy.Slice("foo", "bar", "baz")
	for _, v := range in {
		out := ContainsFunc(in, func(elem string) bool { return elem == v })
		if !out {
			t.Errorf("ContainsFunc(...) = %t, want true", out)
		}
	}

	out := ContainsFunc(in, func(elem string) bool { return false })
	if out {
		t.Errorf("ContainsFunc(...) = %t, want false", out)
	}
}

func TestContains(t *testing.T) {
	in := handy.Slice("foo", "bar", "baz")
	for _, v := range in {
		out := Contains(in, v)
		if !out {
			t.Errorf("Contains(...) = %t, want true", out)
		}
	}

	out := Contains(in, "foobar")
	if out {
		t.Errorf("Contains(...) = %t, want false", out)
	}
}
