package matchers

import (
	"testing"
	"time"

	"github.com/rchampourlier/go-lib/slices"
)

// MatchString matches the `expected` and `matched` strings and
// raise a test error on `t` if they don't match.
func MatchString(t *testing.T, label, expected, matched string, context interface{}) {
	t.Helper()
	if expected != matched {
		t.Errorf("expected %s to be `%s`, got `%s` (%s)", label, expected, matched, context)
	}
}

// MatchStringSlices matches the `expected` and `matched` string slices
// and raise a test error on `t` if they don't.
func MatchStringSlices(t *testing.T, label string, expected, matched []string, context interface{}) {
	t.Helper()
	if !slices.StringsMatch(expected, matched) {
		t.Errorf("expected %s to be `%s`, got `%s` (%s)", label, expected, matched, context)
	}
}

// MatchStringPtr matches the string values for the `expected` and `matched` pointers and
// raise a test error on `t` if they don't match.
func MatchStringPtr(t *testing.T, label string, expected, matched *string, context interface{}) {
	t.Helper()
	if expected == nil && matched == nil {
		return
	}
	if expected == nil && matched != nil {
		t.Errorf("expected %s to be nil, got `%s` (%s)", label, *matched, context)
	} else if expected != nil && matched == nil {
		t.Errorf("expected %s to be `%s`, got nil (%s)", label, *expected, context)
	} else if *expected != *matched {
		t.Errorf("expected %s to be `%s`, got `%s` (%s)", label, *expected, *matched, context)
	}
}

// MatchTime matches the `expected` and `matched` times and
// raise a test error on `t` if they don't match.
func MatchTime(t *testing.T, label string, expected, matched time.Time, context interface{}) {
	t.Helper()
	if !expected.Equal(matched) {
		t.Errorf("expected %s to be `%s`, got `%s` (%v)", label, expected, matched, context)
	}
}

// MatchInt matches the `expected` and `matched` int values and
// raise a test error on `t` if they don't match.
func MatchInt(t *testing.T, label string, expected, matched int, context interface{}) {
	t.Helper()
	if expected != matched {
		t.Errorf("expected %s to be `%d`, got `%d` (%v)", label, expected, matched, context)
	}
}
