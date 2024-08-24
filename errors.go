package rfc3986

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilReader = erorr.Error("rfc3986: nil reader")
	errNilWriter = erorr.Error("rfc3986: nil writer")
)
