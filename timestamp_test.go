package golib_test

import (
	"testing"
	"time"

	"github.com/rchampourlier/golib"
)

func TestTimestamp(t *testing.T) {
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	aTime, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	result := golib.Timestamp(aTime)
	expected := "20130203T195400000PST"
	if result != expected {
		t.Errorf("Unexpected timestamp, got %s, expected %s\n", result, expected)
	}
}
