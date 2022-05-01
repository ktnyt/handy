package enums

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/ktnyt/handy"
)

func TestReduce(t *testing.T) {
	in := handy.Slice(1, 2, 3)
	exp := "122333"
	out := Reduce(in, "", func(acc string, elem int) string {
		return acc + strings.Repeat(strconv.FormatInt(int64(elem), 10), elem)
	})
	if out != exp {
		t.Errorf("Reduce(...) = %s, want %s", out, exp)
	}
}
func TestMapReduce(t *testing.T) {
	in := handy.Slice(1, 2, 3)
	exp := handy.Slice("", "1", "122", "122333")
	out := MapReduce(in, "", func(acc string, elem int) string {
		return acc + strings.Repeat(strconv.FormatInt(int64(elem), 10), elem)
	})
	if !reflect.DeepEqual(out, exp) {
		t.Errorf("Reduce(...) = %v, want %v", out, exp)
	}
}
