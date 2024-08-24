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

	if rfc2234.IsDigit(r) {
		return true
	}

	switch r {
	case                          '-' , '.' , '_' , '~':
		return true
	}

	return false
}
