package golib_test

import (
	"strings"
	"testing"

	"github.com/rchampourlier/golib"
)

type ParameterizeTest struct {
	input  string
	output string
	sep    rune
}

var parameterizeTests = []ParameterizeTest{
	{"a small label", "a-small-label", '-'},
	{"a label with UpperCase", "a-label-with-upper-case", '-'},
	{"downCaseUp_underscore", "down-case-up-underscore", '-'},
	{"multiple  Spaces", "multiple-spaces", '-'},
}

func TestParameterize(t *testing.T) {
	for _, test := range parameterizeTests {
		if golib.Parameterize(test.input, test.sep) != test.output {
			t.Errorf("Parameterize(%q) -> %q, want %q", test.input, golib.Parameterize(test.input, '-'), test.output)
		}
	}
}

func BenchmarkParameterize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range parameterizeTests {
			golib.Parameterize(test.input, '-')
		}
	}
}

func BenchmarkToLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range parameterizeTests {
			strings.ToLower(test.input)
		}
	}
}
