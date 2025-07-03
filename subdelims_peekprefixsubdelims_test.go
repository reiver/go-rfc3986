package rfc3986_test

import (
	"testing"

	"github.com/reiver/go-rfc3986"
)

func TestPeekPrefixSubDelims(t *testing.T) {

	tests := []struct{
		String string
		ExpectedRune rune
		ExpectedLength int
		ExpectedFound bool
	}{
		{
		},



		{
			String:       "A",
			ExpectedRune: 0,
			ExpectedLength: 0,
			ExpectedFound: false,
		},
		{
			String:       "a",
			ExpectedRune: 0,
			ExpectedLength: 0,
			ExpectedFound: false,
		},
		{
			String:       "1",
			ExpectedRune: 0,
			ExpectedLength: 0,
			ExpectedFound: false,
		},



		{
			String:       "!",
			ExpectedRune: '!',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "$",
			ExpectedRune: '$',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "&",
			ExpectedRune: '&',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "'",
			ExpectedRune: '\'',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "(",
			ExpectedRune: '(',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       ")",
			ExpectedRune: ')',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "*",
			ExpectedRune: '*',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "+",
			ExpectedRune: '+',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       ",",
			ExpectedRune: ',',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       ";",
			ExpectedRune: ';',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "=",
			ExpectedRune: '=',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
	}

	for testNumber, test := range tests {

		actualRune, actualLength, actualFound := rfc3986.PeekPrefixSubDelims(test.String)

		{
			expected := test.ExpectedFound
			actual := actualFound

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			expected := test.ExpectedLength
			actual := actualLength

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			expected := test.ExpectedRune
			actual := actualRune

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q (%U)", expected, expected)
				t.Logf("ACTUAL:   %q (%U)", actual, actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}
