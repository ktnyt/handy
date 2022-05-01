package enums

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/ktnyt/handy"
)

func TestMap(t *testing.T) {
	in := handy.Slice(1, 2, 3)
	exp := handy.Slice("1", "2", "3")
	out := Map(in, func(elem int) string {
		return strconv.FormatInt(int64(elem), 10)
	})
	if !reflect.DeepEqual(out, exp) {
		t.Errorf("Map(...) = %v, want %v", out, exp)
	}
}
