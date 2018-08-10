package matchers

import (
	"regexp"
	"testing"
	"time"

	"github.com/rchampourlier/golib/slices"
)

// MatchString matches the `expected` and `matched` strings and
// raise a test error on `t` if they don't match.
func MatchString(t *testing.T, label, expected, matched string, context interface{}) {
	t.Helper()
	if expected != matched {
		t.Errorf("expected %s to be `%s`, got `%s` (%s)", label, expected, matched, context)
	}
}

// MatchStringWithRegex verifies if the `matched` string against
// the `expected` regexp and raises a test error on `t` if they
// don't.
func MatchStringWithRegex(t *testing.T, label, expected, matched string, context interface{}) {
	t.Helper()
	result, err := regexp.MatchString(expected, matched)
	if !result {
		t.Errorf("expected %s to match `%s`, got `%s` (%s)", label, expected, matched, context)
	}
	if err != nil {
		t.Fatal(err)
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

// MatchTimeApprox matches the `expected` and `matched` times and
// raise a test error on `t` if they don't match. The match is
// performed approximately, allowing a difference of `tolerance`
// milliseconds.
func MatchTimeApprox(t *testing.T, label string, expected, matched time.Time, tolerance int, context interface{}) {
	t.Helper()
	diff := expected.Sub(matched).Nanoseconds() / 1000000
	if diff < 0 {
		diff = -diff
	}
	if diff > int64(tolerance) {
		t.Errorf("expected %s to be `%s` +/- %d ms, got `%s` (%v)", label, expected, tolerance, matched, context)
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
