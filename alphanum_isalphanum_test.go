package rfc3986_test

import (
	"testing"

	"github.com/reiver/go-ascii"

	"github.com/reiver/go-rfc3986"
)

func TestIsAlphaNum(t *testing.T) {

	tests := []struct{
		Rune rune
		Expected bool
	}{
		{
			Rune:     ascii.NUL,
			Expected: false,
		},
		{
			Rune:     ascii.SOH,
			Expected: false,
		},
		{
			Rune:     ascii.STX,
			Expected: false,
		},
		{
			Rune:     ascii.ETX,
			Expected: false,
		},
		{
			Rune:     ascii.EOT,
			Expected: false,
		},
		{
			Rune:     ascii.ENQ,
			Expected: false,
		},
		{
			Rune:     ascii.ACK,
			Expected: false,
		},
		{
			Rune:     ascii.BEL,
			Expected: false,
		},
		{
			Rune:     ascii.BS,
			Expected: false,
		},
		{
			Rune:     ascii.HT,
			Expected: false,
		},
		{
			Rune:     ascii.LF,
			Expected: false,
		},
		{
			Rune:     ascii.VT,
			Expected: false,
		},
		{
			Rune:     ascii.FF,
			Expected: false,
		},
		{
			Rune:     ascii.CR,
			Expected: false,
		},
		{
			Rune:     ascii.SO,
			Expected: false,
		},
		{
			Rune:     ascii.SI,
			Expected: false,
		},
		{
			Rune:     ascii.DLE,
			Expected: false,
		},
		{
			Rune:     ascii.DC1,
			Expected: false,
		},
		{
			Rune:     ascii.DC2,
			Expected: false,
		},
		{
			Rune:     ascii.DC3,
			Expected: false,
		},
		{
			Rune:     ascii.DC4,
			Expected: false,
		},
		{
			Rune:     ascii.NAK,
			Expected: false,
		},
		{
			Rune:     ascii.SYN,
			Expected: false,
		},
		{
			Rune:     ascii.ETB,
			Expected: false,
		},
		{
			Rune:     ascii.CAN,
			Expected: false,
		},
		{
			Rune:     ascii.EM,
			Expected: false,
		},
		{
			Rune:     ascii.SUB,
			Expected: false,
		},
		{
			Rune:     ascii.ESC,
			Expected: false,
		},
		{
			Rune:     ascii.FS,
			Expected: false,
		},
		{
			Rune:     ascii.GS,
			Expected: false,
		},
		{
			Rune:     ascii.RS,
			Expected: false,
		},
		{
			Rune:     ascii.US,
			Expected: false,
		},
		{
			Rune:     ascii.SP,
			Expected: false,
		},
		{
			Rune:     '!',
			Expected: false,
		},
		{
			Rune:     '"',
			Expected: false,
		},
		{
			Rune:     '#',
			Expected: false,
		},
		{
			Rune:     '$',
			Expected: false,
		},
		{
			Rune:     '%',
			Expected: false,
		},
		{
			Rune:     '&',
			Expected: false,
		},
		{
			Rune:     '\'',
			Expected: false,
		},
		{
			Rune:     '(',
			Expected: false,
		},
		{
			Rune:     ')',
			Expected: false,
		},
		{
			Rune:     '*',
			Expected: false,
		},
		{
			Rune:     '+',
			Expected: false,
		},
		{
			Rune:     ',',
			Expected: false,
		},
		{
			Rune:     '-',
			Expected: false,
		},
		{
			Rune:     '.',
			Expected: false,
		},
		{
			Rune:     '/',
			Expected: false,
		},
		{
			Rune:     '0',
			Expected: true,
		},
		{
			Rune:     '1',
			Expected: true,
		},
		{
			Rune:     '2',
			Expected: true,
		},
		{
			Rune:     '3',
			Expected: true,
		},
		{
			Rune:     '4',
			Expected: true,
		},
		{
			Rune:     '5',
			Expected: true,
		},
		{
			Rune:     '6',
			Expected: true,
		},
		{
			Rune:     '7',
			Expected: true,
		},
		{
			Rune:     '8',
			Expected: true,
		},
		{
			Rune:     '9',
			Expected: true,
		},
		{
			Rune:     ':',
			Expected: false,
		},
		{
			Rune:     ';',
			Expected: false,
		},
		{
			Rune:     '<',
			Expected: false,
		},
		{
			Rune:     '=',
			Expected: false,
		},
		{
			Rune:     '>',
			Expected: false,
		},
		{
			Rune:     '?',
			Expected: false,
		},
		{
			Rune:     '@',
			Expected: false,
		},
		{
			Rune:     'A',
			Expected: true,
		},
		{
			Rune:     'B',
			Expected: true,
		},
		{
			Rune:     'C',
			Expected: true,
		},
		{
			Rune:     'D',
			Expected: true,
		},
		{
			Rune:     'E',
			Expected: true,
		},
		{
			Rune:     'F',
			Expected: true,
		},
		{
			Rune:     'G',
			Expected: true,
		},
		{
			Rune:     'H',
			Expected: true,
		},
		{
			Rune:     'I',
			Expected: true,
		},
		{
			Rune:     'J',
			Expected: true,
		},
		{
			Rune:     'K',
			Expected: true,
		},
		{
			Rune:     'L',
			Expected: true,
		},
		{
			Rune:     'M',
			Expected: true,
		},
		{
			Rune:     'N',
			Expected: true,
		},
		{
			Rune:     'O',
			Expected: true,
		},
		{
			Rune:     'P',
			Expected: true,
		},
		{
			Rune:     'Q',
			Expected: true,
		},
		{
			Rune:     'R',
			Expected: true,
		},
		{
			Rune:     'S',
			Expected: true,
		},
		{
			Rune:     'T',
			Expected: true,
		},
		{
			Rune:     'U',
			Expected: true,
		},
		{
			Rune:     'V',
			Expected: true,
		},
		{
			Rune:     'W',
			Expected: true,
		},
		{
			Rune:     'X',
			Expected: true,
		},
		{
			Rune:     'Y',
			Expected: true,
		},
		{
			Rune:     'Z',
			Expected: true,
		},
		{
			Rune:     '[',
			Expected: false,
		},
		{
			Rune:     '\\',
			Expected: false,
		},
		{
			Rune:     ']',
			Expected: false,
		},
		{
			Rune:     '^',
			Expected: false,
		},
		{
			Rune:     '_',
			Expected: false,
		},
		{
			Rune:     '`',
			Expected: false,
		},
		{
			Rune:     'a',
			Expected: true,
		},
		{
			Rune:     'b',
			Expected: true,
		},
		{
			Rune:     'c',
			Expected: true,
		},
		{
			Rune:     'd',
			Expected: true,
		},
		{
			Rune:     'e',
			Expected: true,
		},
		{
			Rune:     'f',
			Expected: true,
		},
		{
			Rune:     'g',
			Expected: true,
		},
		{
			Rune:     'h',
			Expected: true,
		},
		{
			Rune:     'i',
			Expected: true,
		},
		{
			Rune:     'j',
			Expected: true,
		},
		{
			Rune:     'k',
			Expected: true,
		},
		{
			Rune:     'l',
			Expected: true,
		},
		{
			Rune:     'm',
			Expected: true,
		},
		{
			Rune:     'n',
			Expected: true,
		},
		{
			Rune:     'o',
			Expected: true,
		},
		{
			Rune:     'p',
			Expected: true,
		},
		{
			Rune:     'q',
			Expected: true,
		},
		{
			Rune:     'r',
			Expected: true,
		},
		{
			Rune:     's',
			Expected: true,
		},
		{
			Rune:     't',
			Expected: true,
		},
		{
			Rune:     'u',
			Expected: true,
		},
		{
			Rune:     'v',
			Expected: true,
		},
		{
			Rune:     'w',
			Expected: true,
		},
		{
			Rune:     'x',
			Expected: true,
		},
		{
			Rune:     'y',
			Expected: true,
		},
		{
			Rune:     'z',
			Expected: true,
		},
		{
			Rune:     '{',
			Expected: false,
		},
		{
			Rune:     '|',
			Expected: false,
		},
		{
			Rune:     '}',
			Expected: false,
		},
		{
			Rune:     '~',
			Expected: false,
		},
		{
			Rune:     ascii.DEL,
			Expected: false,
		},


		{
			Rune:     '۰',
			Expected: false,
		},
		{
			Rune:     '۱',
			Expected: false,
		},
		{
			Rune:     '۲',
			Expected: false,
		},
		{
			Rune:     '۳',
			Expected: false,
		},
		{
			Rune:     '۴',
			Expected: false,
		},
		{
			Rune:     '۵',
			Expected: false,
		},
		{
			Rune:     '۶',
			Expected: false,
		},
		{
			Rune:     '۷',
			Expected: false,
		},
		{
			Rune:     '۸',
			Expected: false,
		},
		{
			Rune:     '۹',
			Expected: false,
		},



		{
			Rune:     '😈',
			Expected: false,
		},



		{
			Rune:     '🙂',
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := rfc3986.IsAlphaNum(test.Rune)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'is-alphanum' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("RUNE: %q (%U)", test.Rune, test.Rune)
			continue
		}
	}
}
