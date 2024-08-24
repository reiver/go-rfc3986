package rfc3986

import (
	"github.com/reiver/go-rfc2234"
)

func IsDigit(r rune) bool {
	return rfc2234.IsDigit(r)
}
