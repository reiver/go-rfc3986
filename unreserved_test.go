package rfc3986_test

import (
	"testing"

	"github.com/reiver/go-rfc2234"

	"github.com/reiver/go-rfc3986"
)

func TestIsUnreserved(t *testing.T) {

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

		if rfc2234.IsAlpha(r) ||
		   rfc2234.IsDigit(r) ||
		   '-' == r           ||
		   '.' == r           ||
		   '_' == r           ||
		   '~' == r           {
			test.Expected = true
		}

		tests = append(tests, test)
	}

	for testNumber, test := range tests {
		actual := rfc3986.IsUnreserved(test.Rune)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for rfc3986.Unreserved() is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("RUNE: (%U) %q", test.Rune, string(test.Rune))
			continue
		}
	}
}
