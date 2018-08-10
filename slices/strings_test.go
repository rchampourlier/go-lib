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

func TestStringsMatch(t *testing.T) {
	trueCases := [][][]string{
		[][]string{[]string{}, []string{}},
		[][]string{[]string{"1"}, []string{"1"}},
		[][]string{[]string{"a", "b"}, []string{"b", "a"}},
	}
	falseCases := [][][]string{
		[][]string{[]string{}, []string{"A"}},
		[][]string{[]string{"1"}, []string{"2"}},
		[][]string{[]string{"c", "b"}, []string{"b", "a"}},
	}
	for _, s := range trueCases {
		if !slices.StringsMatch(s[0], s[1]) {
			t.Errorf("expected to return true (`%s` matches `%s`)", s[0], s[1])
		}
	}
	for _, s := range falseCases {
		if slices.StringsMatch(s[0], s[1]) {
			t.Errorf("expected to return false (`%s` matches `%s`)", s[0], s[1])
		}
	}
}
