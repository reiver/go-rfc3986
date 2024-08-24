package rfc3986_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-rfc3986"
)

func TestWritePctEncodedByte(t *testing.T) {

	tests := []struct{
		Value byte
		Expected string
	}{
		{
			Value:       0,
			Expected: "%00",
		},
		{
			Value:       1,
			Expected: "%01",
		},
		{
			Value:       2,
			Expected: "%02",
		},
		{
			Value:       3,
			Expected: "%03",
		},
		{
			Value:       4,
			Expected: "%04",
		},
		{
			Value:       5,
			Expected: "%05",
		},
		{
			Value:       6,
			Expected: "%06",
		},
		{
			Value:       7,
			Expected: "%07",
		},
		{
			Value:       8,
			Expected: "%08",
		},
		{
			Value:       9,
			Expected: "%09",
		},
		{
			Value:      10,
			Expected: "%0A",
		},
		{
			Value:      11,
			Expected: "%0B",
		},
		{
			Value:      12,
			Expected: "%0C",
		},
		{
			Value:      13,
			Expected: "%0D",
		},
		{
			Value:      14,
			Expected: "%0E",
		},
		{
			Value:      15,
			Expected: "%0F",
		},
		{
			Value:      16,
			Expected: "%10",
		},
		{
			Value:      17,
			Expected: "%11",
		},
		{
			Value:      18,
			Expected: "%12",
		},
		{
			Value:      19,
			Expected: "%13",
		},
		{
			Value:      20,
			Expected: "%14",
		},
		{
			Value:      21,
			Expected: "%15",
		},
		{
			Value:      22,
			Expected: "%16",
		},
		{
			Value:      23,
			Expected: "%17",
		},
		{
			Value:      24,
			Expected: "%18",
		},
		{
			Value:      25,
			Expected: "%19",
		},
		{
			Value:      26,
			Expected: "%1A",
		},
		{
			Value:      27,
			Expected: "%1B",
		},
		{
			Value:      28,
			Expected: "%1C",
		},
		{
			Value:      29,
			Expected: "%1D",
		},
		{
			Value:      30,
			Expected: "%1E",
		},
		{
			Value:      31,
			Expected: "%1F",
		},
		{
			Value:      32,
			Expected: "%20",
		},
		{
			Value:      33,
			Expected: "%21",
		},
		{
			Value:      34,
			Expected: "%22",
		},
		{
			Value:      35,
			Expected: "%23",
		},
		{
			Value:      36,
			Expected: "%24",
		},
		{
			Value:      37,
			Expected: "%25",
		},
		{
			Value:      38,
			Expected: "%26",
		},
		{
			Value:      39,
			Expected: "%27",
		},
		{
			Value:      40,
			Expected: "%28",
		},
		{
			Value:      41,
			Expected: "%29",
		},
		{
			Value:      42,
			Expected: "%2A",
		},
		{
			Value:      43,
			Expected: "%2B",
		},
		{
			Value:      44,
			Expected: "%2C",
		},
		{
			Value:      45,
			Expected: "%2D",
		},
		{
			Value:      46,
			Expected: "%2E",
		},
		{
			Value:      47,
			Expected: "%2F",
		},
		{
			Value:      48,
			Expected: "%30",
		},



		{
			Value:     143,
			Expected: "%8F",
		},
		{
			Value:     144,
			Expected: "%90",
		},
		{
			Value:     145,
			Expected: "%91",
		},
		{
			Value:     146,
			Expected: "%92",
		},
		{
			Value:     147,
			Expected: "%93",
		},
		{
			Value:     148,
			Expected: "%94",
		},
		{
			Value:     149,
			Expected: "%95",
		},
		{
			Value:     150,
			Expected: "%96",
		},
		{
			Value:     151,
			Expected: "%97",
		},
		{
			Value:     152,
			Expected: "%98",
		},
		{
			Value:     153,
			Expected: "%99",
		},
		{
			Value:     154,
			Expected: "%9A",
		},
		{
			Value:     155,
			Expected: "%9B",
		},
		{
			Value:     156,
			Expected: "%9C",
		},
		{
			Value:     157,
			Expected: "%9D",
		},
		{
			Value:     158,
			Expected: "%9E",
		},
		{
			Value:     159,
			Expected: "%9F",
		},
		{
			Value:     160,
			Expected: "%A0",
		},
		{
			Value:     161,
			Expected: "%A1",
		},
		{
			Value:     162,
			Expected: "%A2",
		},
		{
			Value:     163,
			Expected: "%A3",
		},
		{
			Value:     164,
			Expected: "%A4",
		},
		{
			Value:     165,
			Expected: "%A5",
		},
		{
			Value:     166,
			Expected: "%A6",
		},
		{
			Value:     167,
			Expected: "%A7",
		},
		{
			Value:     168,
			Expected: "%A8",
		},
		{
			Value:     169,
			Expected: "%A9",
		},



		{
			Value:     240,
			Expected: "%F0",
		},
		{
			Value:     241,
			Expected: "%F1",
		},
		{
			Value:     242,
			Expected: "%F2",
		},
		{
			Value:     243,
			Expected: "%F3",
		},
		{
			Value:     244,
			Expected: "%F4",
		},
		{
			Value:     245,
			Expected: "%F5",
		},
		{
			Value:     246,
			Expected: "%F6",
		},
		{
			Value:     247,
			Expected: "%F7",
		},
		{
			Value:     248,
			Expected: "%F8",
		},
		{
			Value:     249,
			Expected: "%F9",
		},
		{
			Value:     250,
			Expected: "%FA",
		},
		{
			Value:     251,
			Expected: "%FB",
		},
		{
			Value:     252,
			Expected: "%FC",
		},
		{
			Value:     253,
			Expected: "%FD",
		},
		{
			Value:     254,
			Expected: "%FE",
		},
		{
			Value:     255,
			Expected: "%FF",
		},
	}

	for testNumber, test := range tests {

		var buffer strings.Builder

		err := rfc3986.WritePctEncodedByte(&buffer, test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q (%d) (%X)", test.Value, test.Value, test.Value)
			t.Logf("EXPECTED: %q", test.Expected)
			continue
		}

		{
			actual := buffer.String()
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual pct-encoded (i.e., percent-encoded) value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %q (%d) (%X)", test.Value, test.Value, test.Value)
				continue
			}
		}
	}
}
