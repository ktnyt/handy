package handy

import (
	"fmt"
	"testing"
)

type TestError string

func (s TestError) Error() string {
	return string(s)
}

func TestAs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		exp := TestError("error")
		in := fmt.Errorf("wrap: %w", exp)
		out, ok := As[TestError](in)
		if out != exp || !ok {
			t.Errorf("As[TestError](%v) = (%v, %t), want (%v, true)", in, out, ok, exp)
		}
	})

	t.Run("failure", func(t *testing.T) {
		exp := TestError("error")
		in := fmt.Errorf("wrap: %w", exp)
		out, ok := As[*TestError](in)
		if out != nil || ok {
			t.Errorf("As[*TestError](%v) = (%v, %t), want (nil, false)", in, out, ok)
		}
	})
}

func TestIsType(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		exp := TestError("error")
		in := fmt.Errorf("wrap: %w", exp)
		ok := IsType[TestError](in)
		if !ok {
			t.Errorf("IsType[TestError](%v) = %t, want true", in, ok)
		}
	})

	t.Run("failure", func(t *testing.T) {
		exp := TestError("error")
		in := fmt.Errorf("wrap: %w", exp)
		ok := IsType[*TestError](in)
		if ok {
			t.Errorf("IsType[*TestError](%v) = %t, want false", in, ok)
		}
	})
}
