package enums

import (
	"reflect"
	"testing"

	"github.com/ktnyt/handy"
)

func TestFilter(t *testing.T) {
	in := handy.Slice(1, 2, 3)
	exp := handy.Slice(1, 3)
	out := Filter(in, func(elem int) bool {
		return elem%2 == 1
	})
	if !reflect.DeepEqual(out, exp) {
		t.Errorf("Filter(...) = %v, want %v", out, exp)
	}
}
