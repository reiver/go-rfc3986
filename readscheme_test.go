package rfc3986_test

import (
	"testing"

	"io"
	"strings"

	"github.com/reiver/go-rfc3986"
)

func TestReadScheme(t *testing.T) {

	tests := []struct{
		Data string
		Expected string
	}{
		{
			Data:     "http:",
			Expected: "http",
		},
		{
			Data:     "http://",
			Expected: "http",
		},
		{
			Data:     "http://example.com",
			Expected: "http",
		},
		{
			Data:     "http://example.com/once/twice/thrice/fource.html",
			Expected: "http",
		},
		{
			Data:     "http://example.com/once/twice/thrice/fource.html#middle",
			Expected: "http",
		},



		{
			Data:     "HTTP:",
			Expected: "HTTP",
		},
		{
			Data:     "HTTP://",
			Expected: "HTTP",
		},
		{
			Data:     "HTTP://example.com",
			Expected: "HTTP",
		},
		{
			Data:     "HTTP://example.com/once/twice/thrice/fource.html",
			Expected: "HTTP",
		},
		{
			Data:     "HTTP://example.com/once/twice/thrice/fource.html#middle",
			Expected: "HTTP",
		},



		{
			Data:     "wxyz123:",
			Expected: "wxyz123",
		},
		{
			Data:     "wxyz123:/apple/banana/cherry",
			Expected: "wxyz123",
		},
		{
			Data:     "wxyz123:/apple/banana/cherry?aaa=22",
			Expected: "wxyz123",
		},



		{
			Data:     "WXYZ123:",
			Expected: "WXYZ123",
		},
		{
			Data:     "WXYZ123:/apple/banana/cherry",
			Expected: "WXYZ123",
		},
		{
			Data:     "WXYZ123:/apple/banana/cherry?aaa=22",
			Expected: "WXYZ123",
		},



		{
			Data:     "example+123:",
			Expected: "example+123",
		},
		{
			Data:     "example-123:",
			Expected: "example-123",
		},
		{
			Data:     "example.123:",
			Expected: "example.123",
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Data)

		actual, err := rfc3986.ReadScheme(reader)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("DATA: %q", test.Data)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual read-scheme is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("DATA:     %q", test.Data)
			continue
		}
	}
}

func TestReadScheme_fail(t *testing.T) {

	tests := []struct{
		Data string
		ExpectedError string
	}{
		{
			Data: "",
			ExpectedError: `rfc3986: bad scheme — no data`,
		},



		{
			Data: ":",
			ExpectedError: `rfc3986: bad scheme — scheme must be at least one character long`,
		},



		{
			Data: "http",
			ExpectedError: `rfc3986: bad scheme — scheme terminator character (0x3A) (U+003A) (':') not found`,
		},
		{
			Data: "HTTP",
			ExpectedError: `rfc3986: bad scheme — scheme terminator character (0x3A) (U+003A) (':') not found`,
		},



		{
			Data: "a😈2:",
			ExpectedError: `rfc3986: bad scheme — byte №2 has value 0xF0 is not allowed in scheme`,
		},



		{
			Data: "+example:",
			ExpectedError: `rfc3986: bad scheme — byte №1 has value 0x2B is not allowed for the first character of scheme`,
		},
		{
			Data: "-example:",
			ExpectedError: `rfc3986: bad scheme — byte №1 has value 0x2D is not allowed for the first character of scheme`,
		},
		{
			Data: ".example:",
			ExpectedError: `rfc3986: bad scheme — byte №1 has value 0x2E is not allowed for the first character of scheme`,
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Data)

		_, err := rfc3986.ReadScheme(reader)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("DATA: %q", test.Data)
			continue
		}

		actual := err.Error()
		expected := test.ExpectedError

		if expected != actual {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("DATA:     %q", test.Data)
			continue
		}
	}
}
