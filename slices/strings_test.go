package slices_test

import (
	"testing"

	"github.com/rchampourlier/golib/slices"
)

func TestStringsContain(t *testing.T) {
	trueCases := map[string][]string{
		"in": {"in", "other"},
	}
	falseCases := map[string][]string{
		"missing": {"string", "other"},
	}
	for str, slice := range trueCases {
		if slices.StringsContain(slice, str) != true {
			t.Errorf("expected to return true (`%s` in `%s`)", str, slice)
		}
	}
	for str, slice := range falseCases {
		if slices.StringsContain(slice, str) != false {
			t.Errorf("expected to return false (`%s` not in `%s`)", str, slice)
		}
	}
}
