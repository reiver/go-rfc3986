package rfc3986_test

import (
	"testing"

	"github.com/reiver/go-rfc3986"
)

func TestPeekPrefixScheme(t *testing.T) {

	tests := []struct{
		String string
		ExpectedScheme string
		ExpectedLength int
		ExpectedFound bool
	}{
		{

		},



		{
			String: "apple",
		},
		{
			String: "BANANA",
		},
		{
			String: "Cherry",
		},
		{
			String: "dAte",
		},



		{
			String:             "apple:",
			ExpectedScheme:     "apple",
			ExpectedLength: len("apple"),
			ExpectedFound: true,
		},
		{
			String:             "BANANA:",
			ExpectedScheme:     "banana",
			ExpectedLength: len("banana"),
			ExpectedFound: true,
		},
		{
			String:             "Cherry:",
			ExpectedScheme:     "cherry",
			ExpectedLength: len("cherry"),
			ExpectedFound: true,
		},
		{
			String:             "dAtE:",
			ExpectedScheme:     "date",
			ExpectedLength: len("date"),
			ExpectedFound: true,
		},



		{
			String:             "http://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "httP://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "htTp://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "htTP://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "hTtp://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "hTtP://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "hTTp://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "hTTP://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "Http://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "HttP://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "HtTp://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "HtTP://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "HTtp://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "HTtP://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "HTTp://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
		{
			String:             "HTTP://example.com/once/twice/thrice/fource.html",
			ExpectedScheme:     "http",
			ExpectedLength: len("http"),
			ExpectedFound: true,
		},
	}

	for testNumber, test := range tests {

		actualScheme, actualLength, actualFound := rfc3986.PeekPrefixScheme(test.String)

		{
			actual := actualFound
			expected := test.ExpectedFound

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				continue
			}
		}

		{
			actual := actualLength
			expected := test.ExpectedLength

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue
			}
		}

		{
			actual := actualScheme
			expected := test.ExpectedScheme

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}
	}
}
