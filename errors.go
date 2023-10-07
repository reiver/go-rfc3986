package rfc3986

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilReader = erorr.Error("rfc3986: nil reader")
	errNilWriter = erorr.Error("rfc3986: nil writer")
)
