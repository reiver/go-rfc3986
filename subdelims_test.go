package rfc3986_test

import (
	"testing"

	"github.com/reiver/go-rfc3986"
)

func TestIsSubDelim(t *testing.T) {

	tests := []struct{
		Rune rune
		Expected bool
	}{
	}

	for r:=rune(0); r < rune(8192); r++ {
		test := struct{
			Rune rune
			Expected bool
		}{
			Rune: r,
			Expected: false,
		}

		if '!' == r ||
		   '$' == r ||
		   '&' == r ||
		  '\'' == r ||
		   '(' == r ||
		   ')' == r ||
		   '*' == r ||
		   '+' == r ||
		   ',' == r ||
		   ';' == r ||
		   '=' == r {
			test.Expected = true
		}

		tests = append(tests, test)
	}

	for testNumber, test := range tests {
		actual := rfc3986.IsSubDelim(test.Rune)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for rfc3986.IsSubDelim() is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("RUNE: (%U) %q", test.Rune, string(test.Rune))
			continue
		}
	}
}
