package rfc3986

import (
	"testing"

	"github.com/reiver/go-ascii"
)

func TestASCIIToLower(t *testing.T) {
	tests := []struct{
		Rune rune
		Expected rune
	}{
		{
			Rune:     ascii.NUL,
			Expected: ascii.NUL,
		},
		{
			Rune:     ascii.SOH,
			Expected: ascii.SOH,
		},
		{
			Rune:     ascii.STX,
			Expected: ascii.STX,
		},
		{
			Rune:     ascii.ETX,
			Expected: ascii.ETX,
		},
		{
			Rune:     ascii.EOT,
			Expected: ascii.EOT,
		},
		{
			Rune:     ascii.ENQ,
			Expected: ascii.ENQ,
		},
		{
			Rune:     ascii.ACK,
			Expected: ascii.ACK,
		},
		{
			Rune:     ascii.BEL,
			Expected: ascii.BEL,
		},
		{
			Rune:     ascii.BS,
			Expected: ascii.BS,
		},
		{
			Rune:     ascii.HT,
			Expected: ascii.HT,
		},
		{
			Rune:     ascii.LF,
			Expected: ascii.LF,
		},
		{
			Rune:     ascii.VT,
			Expected: ascii.VT,
		},
		{
			Rune:     ascii.FF,
			Expected: ascii.FF,
		},
		{
			Rune:     ascii.CR,
			Expected: ascii.CR,
		},
		{
			Rune:     ascii.SO,
			Expected: ascii.SO,
		},
		{
			Rune:     ascii.SI,
			Expected: ascii.SI,
		},
		{
			Rune:     ascii.DLE,
			Expected: ascii.DLE,
		},
		{
			Rune:     ascii.DC1,
			Expected: ascii.DC1,
		},
		{
			Rune:     ascii.DC2,
			Expected: ascii.DC2,
		},
		{
			Rune:     ascii.DC3,
			Expected: ascii.DC3,
		},
		{
			Rune:     ascii.DC4,
			Expected: ascii.DC4,
		},
		{
			Rune:     ascii.NAK,
			Expected: ascii.NAK,
		},
		{
			Rune:     ascii.SYN,
			Expected: ascii.SYN,
		},
		{
			Rune:     ascii.ETB,
			Expected: ascii.ETB,
		},
		{
			Rune:     ascii.CAN,
			Expected: ascii.CAN,
		},
		{
			Rune:     ascii.EM,
			Expected: ascii.EM,
		},
		{
			Rune:     ascii.SUB,
			Expected: ascii.SUB,
		},
		{
			Rune:     ascii.ESC,
			Expected: ascii.ESC,
		},
		{
			Rune:     ascii.FS,
			Expected: ascii.FS,
		},
		{
			Rune:     ascii.GS,
			Expected: ascii.GS,
		},
		{
			Rune:     ascii.RS,
			Expected: ascii.RS,
		},
		{
			Rune:     ascii.US,
			Expected: ascii.US,
		},
		{
			Rune:     ascii.SP,
			Expected: ascii.SP,
		},
		{
			Rune:     '!',
			Expected: '!',
		},
		{
			Rune:     '"',
			Expected: '"',
		},
		{
			Rune:     '#',
			Expected: '#',
		},
		{
			Rune:     '$',
			Expected: '$',
		},
		{
			Rune:     '%',
			Expected: '%',
		},
		{
			Rune:     '&',
			Expected: '&',
		},
		{
			Rune:     '\'',
			Expected: '\'',
		},
		{
			Rune:     '(',
			Expected: '(',
		},
		{
			Rune:     ')',
			Expected: ')',
		},
		{
			Rune:     '*',
			Expected: '*',
		},
		{
			Rune:     '+',
			Expected: '+',
		},
		{
			Rune:     ',',
			Expected: ',',
		},
		{
			Rune:     '-',
			Expected: '-',
		},
		{
			Rune:     '.',
			Expected: '.',
		},
		{
			Rune:     '/',
			Expected: '/',
		},
		{
			Rune:     '0',
			Expected: '0',
		},
		{
			Rune:     '1',
			Expected: '1',
		},
		{
			Rune:     '2',
			Expected: '2',
		},
		{
			Rune:     '3',
			Expected: '3',
		},
		{
			Rune:     '4',
			Expected: '4',
		},
		{
			Rune:     '5',
			Expected: '5',
		},
		{
			Rune:     '6',
			Expected: '6',
		},
		{
			Rune:     '7',
			Expected: '7',
		},
		{
			Rune:     '8',
			Expected: '8',
		},
		{
			Rune:     '9',
			Expected: '9',
		},
		{
			Rune:     ':',
			Expected: ':',
		},
		{
			Rune:     ';',
			Expected: ';',
		},
		{
			Rune:     '<',
			Expected: '<',
		},
		{
			Rune:     '=',
			Expected: '=',
		},
		{
			Rune:     '>',
			Expected: '>',
		},
		{
			Rune:     '?',
			Expected: '?',
		},
		{
			Rune:     '@',
			Expected: '@',
		},
		{
			Rune:     'A',
			Expected: 'a',
		},
		{
			Rune:     'B',
			Expected: 'b',
		},
		{
			Rune:     'C',
			Expected: 'c',
		},
		{
			Rune:     'D',
			Expected: 'd',
		},
		{
			Rune:     'E',
			Expected: 'e',
		},
		{
			Rune:     'F',
			Expected: 'f',
		},
		{
			Rune:     'G',
			Expected: 'g',
		},
		{
			Rune:     'H',
			Expected: 'h',
		},
		{
			Rune:     'I',
			Expected: 'i',
		},
		{
			Rune:     'J',
			Expected: 'j',
		},
		{
			Rune:     'K',
			Expected: 'k',
		},
		{
			Rune:     'L',
			Expected: 'l',
		},
		{
			Rune:     'M',
			Expected: 'm',
		},
		{
			Rune:     'N',
			Expected: 'n',
		},
		{
			Rune:     'O',
			Expected: 'o',
		},
		{
			Rune:     'P',
			Expected: 'p',
		},
		{
			Rune:     'Q',
			Expected: 'q',
		},
		{
			Rune:     'R',
			Expected: 'r',
		},
		{
			Rune:     'S',
			Expected: 's',
		},
		{
			Rune:     'T',
			Expected: 't',
		},
		{
			Rune:     'U',
			Expected: 'u',
		},
		{
			Rune:     'V',
			Expected: 'v',
		},
		{
			Rune:     'W',
			Expected: 'w',
		},
		{
			Rune:     'X',
			Expected: 'x',
		},
		{
			Rune:     'Y',
			Expected: 'y',
		},
		{
			Rune:     'Z',
			Expected: 'z',
		},
		{
			Rune:     '[',
			Expected: '[',
		},
		{
			Rune:     '\\',
			Expected: '\\',
		},
		{
			Rune:     ']',
			Expected: ']',
		},
		{
			Rune:     '^',
			Expected: '^',
		},
		{
			Rune:     '_',
			Expected: '_',
		},
		{
			Rune:     '`',
			Expected: '`',
		},
		{
			Rune:     'a',
			Expected: 'a',
		},
		{
			Rune:     'b',
			Expected: 'b',
		},
		{
			Rune:     'c',
			Expected: 'c',
		},
		{
			Rune:     'd',
			Expected: 'd',
		},
		{
			Rune:     'e',
			Expected: 'e',
		},
		{
			Rune:     'f',
			Expected: 'f',
		},
		{
			Rune:     'g',
			Expected: 'g',
		},
		{
			Rune:     'h',
			Expected: 'h',
		},
		{
			Rune:     'i',
			Expected: 'i',
		},
		{
			Rune:     'j',
			Expected: 'j',
		},
		{
			Rune:     'k',
			Expected: 'k',
		},
		{
			Rune:     'l',
			Expected: 'l',
		},
		{
			Rune:     'm',
			Expected: 'm',
		},
		{
			Rune:     'n',
			Expected: 'n',
		},
		{
			Rune:     'o',
			Expected: 'o',
		},
		{
			Rune:     'p',
			Expected: 'p',
		},
		{
			Rune:     'q',
			Expected: 'q',
		},
		{
			Rune:     'r',
			Expected: 'r',
		},
		{
			Rune:     's',
			Expected: 's',
		},
		{
			Rune:     't',
			Expected: 't',
		},
		{
			Rune:     'u',
			Expected: 'u',
		},
		{
			Rune:     'v',
			Expected: 'v',
		},
		{
			Rune:     'w',
			Expected: 'w',
		},
		{
			Rune:     'x',
			Expected: 'x',
		},
		{
			Rune:     'y',
			Expected: 'y',
		},
		{
			Rune:     'z',
			Expected: 'z',
		},
		{
			Rune:     '{',
			Expected: '{',
		},
		{
			Rune:     '|',
			Expected: '|',
		},
		{
			Rune:     '}',
			Expected: '}',
		},
		{
			Rune:     '~',
			Expected: '~',
		},
		{
			Rune:     ascii.DEL,
			Expected: ascii.DEL,
		},


		{
			Rune:     '۰',
			Expected: '۰',
		},
		{
			Rune:     '۱',
			Expected: '۱',
		},
		{
			Rune:     '۲',
			Expected: '۲',
		},
		{
			Rune:     '۳',
			Expected: '۳',
		},
		{
			Rune:     '۴',
			Expected: '۴',
		},
		{
			Rune:     '۵',
			Expected: '۵',
		},
		{
			Rune:     '۶',
			Expected: '۶',
		},
		{
			Rune:     '۷',
			Expected: '۷',
		},
		{
			Rune:     '۸',
			Expected: '۸',
		},
		{
			Rune:     '۹',
			Expected: '۹',
		},



		{
			Rune:     '😈',
			Expected: '😈',
		},



		{
			Rune:     '🙂',
			Expected: '🙂',
		},
	}

	for testNumber, test := range tests {

		actual := asciiToLower(test.Rune)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual to-lowered rune it not what was expected", testNumber)
			t.Logf("EXPECTED: %q (%U)", expected, expected)
			t.Logf("ACTUAL:   %q (%U)", actual, actual)
			continue
		}
	}
}
