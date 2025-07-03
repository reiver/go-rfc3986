package rfc3986_test

import (
	"testing"

	"github.com/reiver/go-rfc3986"
)

func TestDecodePctEncoded(t *testing.T) {

	tests := []struct{
		Byte2 byte
		Byte3 byte
		Expected byte
	}{
		{
			Byte2:     '0',
			Byte3:      '0',
			Expected: 0x00,
		},
		{
			Byte2:     '0',
			Byte3:      '1',
			Expected: 0x01,
		},
		{
			Byte2:     '0',
			Byte3:      '2',
			Expected: 0x02,
		},
		{
			Byte2:     '0',
			Byte3:      '3',
			Expected: 0x03,
		},
		{
			Byte2:     '0',
			Byte3:      '4',
			Expected: 0x04,
		},
		{
			Byte2:     '0',
			Byte3:      '5',
			Expected: 0x05,
		},
		{
			Byte2:     '0',
			Byte3:      '6',
			Expected: 0x06,
		},
		{
			Byte2:     '0',
			Byte3:      '7',
			Expected: 0x07,
		},
		{
			Byte2:     '0',
			Byte3:      '8',
			Expected: 0x08,
		},
		{
			Byte2:     '0',
			Byte3:      '9',
			Expected: 0x09,
		},
		{
			Byte2:     '0',
			Byte3:      'A',
			Expected: 0x0A,
		},
		{
			Byte2:     '0',
			Byte3:      'a',
			Expected: 0x0A,
		},
		{
			Byte2:     '0',
			Byte3:      'B',
			Expected: 0x0B,
		},
		{
			Byte2:     '0',
			Byte3:      'b',
			Expected: 0x0B,
		},
		{
			Byte2:     '0',
			Byte3:      'C',
			Expected: 0x0C,
		},
		{
			Byte2:     '0',
			Byte3:      'c',
			Expected: 0x0C,
		},
		{
			Byte2:     '0',
			Byte3:      'D',
			Expected: 0x0D,
		},
		{
			Byte2:     '0',
			Byte3:      'd',
			Expected: 0x0D,
		},
		{
			Byte2:     '0',
			Byte3:      'E',
			Expected: 0x0E,
		},
		{
			Byte2:     '0',
			Byte3:      'e',
			Expected: 0x0E,
		},
		{
			Byte2:     '0',
			Byte3:      'F',
			Expected: 0x0F,
		},
		{
			Byte2:     '0',
			Byte3:      'f',
			Expected: 0x0F,
		},
		{
			Byte2:     '1',
			Byte3:      '0',
			Expected: 0x10,
		},
		{
			Byte2:     '1',
			Byte3:      '1',
			Expected: 0x11,
		},
		{
			Byte2:     '1',
			Byte3:      '2',
			Expected: 0x12,
		},



		{
			Byte2:     '2',
			Byte3:      'E',
			Expected: 0x2E,
		},
		{
			Byte2:     '2',
			Byte3:      'e',
			Expected: 0x2E,
		},
		{
			Byte2:     '2',
			Byte3:      'F',
			Expected: 0x2F,
		},
		{
			Byte2:     '2',
			Byte3:      'f',
			Expected: 0x2F,
		},
		{
			Byte2:     '3',
			Byte3:      '0',
			Expected: 0x30,
		},
		{
			Byte2:     '3',
			Byte3:      '1',
			Expected: 0x31,
		},
		{
			Byte2:     '3',
			Byte3:      '2',
			Expected: 0x32,
		},



		{
			Byte2:     'A',
			Byte3:      '1',
			Expected: 0xA1,
		},
		{
			Byte2:     'a',
			Byte3:      '1',
			Expected: 0xA1,
		},



		{
			Byte2:     'A',
			Byte3:      'B',
			Expected: 0xAB,
		},
		{
			Byte2:     'A',
			Byte3:      'b',
			Expected: 0xAB,
		},
		{
			Byte2:     'a',
			Byte3:      'B',
			Expected: 0xAB,
		},
		{
			Byte2:     'a',
			Byte3:      'b',
			Expected: 0xAB,
		},



		{
			Byte2:     'F',
			Byte3:      'F',
			Expected: 0xFF,
		},
		{
			Byte2:     'F',
			Byte3:      'f',
			Expected: 0xFF,
		},
	}

	for testNumber, test := range tests {

		actual, err := rfc3986.DecodePctEncoded(test.Byte2, test.Byte3)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("BYTE-2: %q (0x%02X)", test.Byte2, test.Byte2)
			t.Logf("BYTE-3: %q (0x%02X)", test.Byte3, test.Byte3)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual decoded 'pct-encoded' is not what was expected", testNumber)
			t.Logf("EXPECTED: %q (0x%02X)", expected, expected)
			t.Logf("ACTUAL:   %q (0x%02X)", actual, actual)
			t.Logf("BYTE-2: %q (0x%02X)", test.Byte2, test.Byte2)
			t.Logf("BYTE-3: %q (0x%02X)", test.Byte3, test.Byte3)
			continue
		}
	}
}
