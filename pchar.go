package rfc3986

import (
	"strings"
)

// PeekPrefixPChar peeks a string, checks to see if it starts with a 'pchar' (as defined in IETF RFC-3986),
// and returns it if it does.
//
// 'pchar' is defined in IETF RFC-3986 as:
//
//	pchar = unreserved / pct-encoded / sub-delims / ":" / "@"
func PeekPrefixPChar(str string) (rune, int, bool) {
	if len(str) < 1 {
		var nada rune
		return nada, 0, false
	}

	{
		r, n, found := PeekPrefixUnreserved(str)
		if found {
			return r, n, true
		}
	}

	{
		r, n, found := PeekPrefixPctEncoded(str)
		if found {
			return r, n, true
		}
	}

	{
		r, n, found := PeekPrefixSubDelims(str)
		if found {
			return r, n, true
		}
	}

	{
		const r rune = ':'
		const s string = string(r)

		if strings.HasPrefix(str, s) {
			return r, len(s), true
		}
	}

	{
		const r rune = '@'
		const s string = string(r)

		if strings.HasPrefix(str, s) {
			return r, len(s), true
		}
	}

	{
		var nada rune
		return nada, 0, false
	}
}
