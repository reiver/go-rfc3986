package rfc3986_test

import (
	"testing"

	"github.com/reiver/go-rfc3986"
)

func TestPeekPrefixPctEncoded(t *testing.T) {

	tests := []struct{
		String string
		ExpectedRune rune
		ExpectedLength int
		ExpectedFound bool
	}{
		{

		},



		{
			String: "",
		},
		{
			String: "!",
		},
		{
			String: "A",
		},
		{
			String: "a",
		},
		{
			String: "AB",
		},
		{
			String: "Ab",
		},
		{
			String: "aB",
		},
		{
			String: "ab",
		},
		{
			String: "%",
		},
		{
			String: "%A",
		},
		{
			String: "%a",
		},



		{
			String: "%HE",
		},
		{
			String: "%HELLO WORLD",
		},



		{
			String: "&%AB",
		},
		{
			String: "&%AB",
		},



		{
			String:       "%00",
			ExpectedRune: 0x00,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%01",
			ExpectedRune: 0x01,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%02",
			ExpectedRune: 0x02,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%03",
			ExpectedRune: 0x03,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%04",
			ExpectedRune: 0x04,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%05",
			ExpectedRune: 0x05,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%06",
			ExpectedRune: 0x06,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%07",
			ExpectedRune: 0x07,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%08",
			ExpectedRune: 0x08,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%09",
			ExpectedRune: 0x09,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0A",
			ExpectedRune: 0x0A,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0a",
		},
		{
			String:       "%0B",
			ExpectedRune: 0x0B,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0b",
		},
		{
			String:       "%0C",
			ExpectedRune: 0x0C,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0c",
		},
		{
			String:       "%0D",
			ExpectedRune: 0x0D,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0d",
		},
		{
			String:       "%0E",
			ExpectedRune: 0x0E,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0e",
		},
		{
			String:       "%0F",
			ExpectedRune: 0x0F,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0f",
		},
		{
			String:       "%10",
			ExpectedRune: 0x10,
			ExpectedLength:3,
			ExpectedFound:true,
		},

		{
			String:       "%FF",
			ExpectedRune: 0xFF,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%fF",
		},
		{
			String:       "%Ff",
		},
		{
			String:       "%ff",
		},



		{
			String:       "%DE%AD%C0%DE",
			ExpectedRune: 0xDE,
			ExpectedLength:3,
			ExpectedFound:true,
		},
	}

	for testNumber, test := range tests {

		actualRune, actualLength, actualFound := rfc3986.PeekPrefixPctEncoded(test.String)

		{
			expected := test.ExpectedFound
			actual := actualFound

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:  %t", actual)
				t.Logf("STRING: %q", test.String )
				continue
			}
		}

		{
			expected := test.ExpectedLength
			actual := actualLength

			if expected != actual {
				t.Errorf("For test #%d, the actual 'length' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:  %d", actual)
				t.Logf("STRING: %q", test.String )
				continue
			}
		}

		{
			expected := test.ExpectedRune
			actual := actualRune

			if expected != actual {
				t.Errorf("For test #%d, the actual 'rune' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:  %d", actual)
				t.Logf("STRING: %q", test.String )
				continue
			}
		}
	}
}
