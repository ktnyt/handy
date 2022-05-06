package handy

import (
	"errors"
	"testing"
)

func TestMust(t *testing.T) {
	in := 42
	out := Must(in, nil)
	if out != in {
		t.Errorf("Must(%d, nil) = %d, want %d", in, out, in)
	}

	defer func() {
		t.Helper()
		err := recover()
		if err == nil {
			t.Error("expected function to panic")
		}
	}()
	Must(in, errors.New("error"))
}
