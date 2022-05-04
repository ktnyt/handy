package handy

import (
	"errors"
	"testing"
)

func TestUnpackError(t *testing.T) {
	exp := errors.New("error")
	f := func() (int, error) {
		return 0, exp
	}
	out := UnpackError(f())
	if !errors.Is(out, exp) {
		t.Errorf("UnpackError(f()) = %v, want %v", out, exp)
	}
}
