package rfc3986_test

import (
	"testing"

	"github.com/reiver/go-rfc3986"
)

func TestPeekPrefixUnreserved(t *testing.T) {

	tests := []struct{
		String string
		ExpectedRune rune
		ExpectedLength int
		ExpectedFound bool
	}{
		{
		},



		{
			String:       "!",
			ExpectedRune: 0,
			ExpectedLength: 0,
			ExpectedFound: false,
		},
		{
			String:       "@",
			ExpectedRune: 0,
			ExpectedLength: 0,
			ExpectedFound: false,
		},



		{
			String:       "A",
			ExpectedRune: 'A',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "B",
			ExpectedRune: 'B',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "C",
			ExpectedRune: 'C',
			ExpectedLength: 1,
			ExpectedFound: true,
		},



		{
			String:       "a",
			ExpectedRune: 'a',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "b",
			ExpectedRune: 'b',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "c",
			ExpectedRune: 'c',
			ExpectedLength: 1,
			ExpectedFound: true,
		},



		{
			String:       "0",
			ExpectedRune: '0',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "1",
			ExpectedRune: '1',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "2",
			ExpectedRune: '2',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "3",
			ExpectedRune: '3',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "4",
			ExpectedRune: '4',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "5",
			ExpectedRune: '5',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "6",
			ExpectedRune: '6',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "7",
			ExpectedRune: '7',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "8",
			ExpectedRune: '8',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "9",
			ExpectedRune: '9',
			ExpectedLength: 1,
			ExpectedFound: true,
		},



		{
			String:       "-",
			ExpectedRune: '-',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       ".",
			ExpectedRune: '.',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "_",
			ExpectedRune: '_',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "~",
			ExpectedRune: '~',
			ExpectedLength: 1,
			ExpectedFound: true,
		},



		{
			String:       "apple",
			ExpectedRune: 'a',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "BANANA",
			ExpectedRune: 'B',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
		{
			String:       "Cherry",
			ExpectedRune: 'C',
			ExpectedLength: 1,
			ExpectedFound: true,
		},



		{
			String:       "~joeblow",
			ExpectedRune: '~',
			ExpectedLength: 1,
			ExpectedFound: true,
		},



		{
			String:       "Hello world!",
			ExpectedRune: 'H',
			ExpectedLength: 1,
			ExpectedFound: true,
		},
	}

	for testNumber, test := range tests {

		actualRune, actualLength, actualFound := rfc3986.PeekPrefixUnreserved(test.String)

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
