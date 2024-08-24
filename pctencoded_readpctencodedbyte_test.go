package rfc3986_test

import (
	"testing"

	"io"
	"strings"

	"github.com/reiver/go-rfc3986"
)

func TestReadPctEncodedByte(t *testing.T) {

	tests := []struct{
		Value string
		Expected byte
	}{
		{
			Value:    "%00",
			Expected:    0,
		},
		{
			Value:    "%01",
			Expected:    1,
		},
		{
			Value:    "%02",
			Expected:    2,
		},
		{
			Value:    "%03",
			Expected:    3,
		},
		{
			Value:    "%04",
			Expected:    4,
		},
		{
			Value:    "%05",
			Expected:    5,
		},
		{
			Value:    "%06",
			Expected:    6,
		},
		{
			Value:    "%07",
			Expected:    7,
		},
		{
			Value:    "%08",
			Expected:    8,
		},
		{
			Value:    "%09",
			Expected:    9,
		},
		{
			Value:    "%0A",
			Expected:   10,
		},
		{
			Value:    "%0B",
			Expected:   11,
		},
		{
			Value:    "%0C",
			Expected:   12,
		},
		{
			Value:    "%0D",
			Expected:   13,
		},
		{
			Value:    "%0E",
			Expected:   14,
		},
		{
			Value:    "%0F",
			Expected:   15,
		},



		{
			Value:    "%0a",
			Expected:   10,
		},
		{
			Value:    "%0b",
			Expected:   11,
		},
		{
			Value:    "%0c",
			Expected:   12,
		},
		{
			Value:    "%0d",
			Expected:   13,
		},
		{
			Value:    "%0e",
			Expected:   14,
		},
		{
			Value:    "%0f",
			Expected:   15,
		},



		{
			Value:    "%10",
			Expected:   16,
		},
		{
			Value:    "%11",
			Expected:   17,
		},
		{
			Value:    "%12",
			Expected:   18,
		},
		{
			Value:    "%13",
			Expected:   19,
		},
		{
			Value:    "%14",
			Expected:   20,
		},
		{
			Value:    "%15",
			Expected:   21,
		},
		{
			Value:    "%16",
			Expected:   22,
		},
		{
			Value:    "%17",
			Expected:   23,
		},
		{
			Value:    "%18",
			Expected:   24,
		},
		{
			Value:    "%19",
			Expected:   25,
		},
		{
			Value:    "%1A",
			Expected:   26,
		},
		{
			Value:    "%1B",
			Expected:   27,
		},
		{
			Value:    "%1C",
			Expected:   28,
		},
		{
			Value:    "%1D",
			Expected:   29,
		},
		{
			Value:    "%1E",
			Expected:   30,
		},
		{
			Value:    "%1F",
			Expected:   31,
		},



		{
			Value:    "%1a",
			Expected:   26,
		},
		{
			Value:    "%1b",
			Expected:   27,
		},
		{
			Value:    "%1c",
			Expected:   28,
		},
		{
			Value:    "%1d",
			Expected:   29,
		},
		{
			Value:    "%1e",
			Expected:   30,
		},
		{
			Value:    "%1f",
			Expected:   31,
		},



		{
			Value:    "%20F",
			Expected:   32,
		},



		{
			Value:    "%FA",
			Expected:  250,
		},
		{
			Value:    "%FB",
			Expected:  251,
		},
		{
			Value:    "%FC",
			Expected:  252,
		},
		{
			Value:    "%FD",
			Expected:  253,
		},
		{
			Value:    "%FE",
			Expected:  254,
		},
		{
			Value:    "%FF",
			Expected:  255,
		},



		{
			Value:    "%Fa",
			Expected:  250,
		},
		{
			Value:    "%Fb",
			Expected:  251,
		},
		{
			Value:    "%Fc",
			Expected:  252,
		},
		{
			Value:    "%Fd",
			Expected:  253,
		},
		{
			Value:    "%Fe",
			Expected:  254,
		},
		{
			Value:    "%Ff",
			Expected:  255,
		},



		{
			Value:    "%fA",
			Expected:  250,
		},
		{
			Value:    "%fB",
			Expected:  251,
		},
		{
			Value:    "%fC",
			Expected:  252,
		},
		{
			Value:    "%fD",
			Expected:  253,
		},
		{
			Value:    "%fE",
			Expected:  254,
		},
		{
			Value:    "%fF",
			Expected:  255,
		},



		{
			Value:    "%fa",
			Expected:  250,
		},
		{
			Value:    "%fb",
			Expected:  251,
		},
		{
			Value:    "%fc",
			Expected:  252,
		},
		{
			Value:    "%fd",
			Expected:  253,
		},
		{
			Value:    "%fe",
			Expected:  254,
		},
		{
			Value:    "%ff",
			Expected:  255,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Value)

		actual, err := rfc3986.ReadPctEncodedByte(reader)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("EXPECTED: %d (0x%X)", test.Expected, test.Expected)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual byte returned is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d (0x%X)", expected, expected)
				t.Logf("ACTUAL:   %d (0x%X)", actual, actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
