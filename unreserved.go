package rfc3986

import (
	"github.com/reiver/go-rfc2234"
)

// IsUnreserved returns true if the value of 'r' matches 'unreserved' as defined in IETF RFC-3986:
//
//	unreserved  = ALPHA / DIGIT / "-" / "." / "_" / "~"
func IsUnreserved(r rune) bool {
	if rfc2234.IsAlpha(r) {
		return true
	}

	if IsDigit(r) {
		return true
	}

	switch r {
	case                          '-' , '.' , '_' , '~':
		return true
	}

	return false
}

func PeekPrefixUnreserved(str string) (rune, int, bool) {

	r, n, found := firstCharacter(str)
	if !found {
		var nada rune
		return nada, 0, false
	}

	if !IsUnreserved(r) {
		var nada rune
		return nada, 0, false
	}

	return r, n, true
}
