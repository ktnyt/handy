package handy

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	in := []string{"foo", "bar", "baz"}
	out := Slice(in...)
	if !reflect.DeepEqual(in, out) {
		t.Errorf("Slice(%v...) = %v, want %v", in, out, in)
	}
}
