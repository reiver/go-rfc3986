package rfc3986

import (
	"github.com/reiver/go-rfc2234"
)

func IsHexDig(r rune) bool {
	return rfc2234.IsHexDig(r)
}
