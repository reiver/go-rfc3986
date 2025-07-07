package rfc3986

import (
	"io"
	"strings"

	"github.com/reiver/go-erorr"
)

const SchemeTerminator = ':'

// ReadScheme reads and return the 'scheme' as defined by IETF RFC-3986:
//
//	scheme = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
//	
//	ALPHA  = %41-%5A / %61-%7A
//	
//	DIGIT  = %30-%39
//
// Although ReadScheme reads the colon (":") it does not return it as part of the scheme.
func ReadScheme(reader io.Reader) (string, error) {

	if nil == reader {
		return "", errNilReader
	}

	var buffer [512]byte
	var scheme []byte = buffer[0:0]

	{
		var byteNumber int

		var b [1]byte

		loop: for {
			byteNumber++

			n, err := reader.Read(b[:])
			if io.EOF == err {
				switch {
				case len(scheme) <= 0:
					return "", erorr.Errorf("rfc3986: bad scheme — no data")
				default:
					return "", erorr.Errorf("rfc3986: bad scheme — scheme terminator character (0x%02X) (%U) (%q) not found", SchemeTerminator, SchemeTerminator, SchemeTerminator)
				}
			}
			if nil != err {
				return "", erorr.Errorf("rfc3986: could not read the byte №%d in scheme string because: %w", byteNumber, err)
			}
			if 1 != n {
				return "", erorr.Errorf("rfc3986: problem reading the byte №%d in scheme string — expected to have read 1 byte but actually read %d bytes.", byteNumber, n)
			}

			b0 := b[0]
			r0 := rune(b0)

			switch {
			case SchemeTerminator == b0:
		/////////////// BREAK
				break loop
			case IsAlpha(r0) || IsDigit(r0) || '+' == r0 || '-' == r0 || '.' == r0:
				scheme = append(scheme, b0)
			default:
				return "", erorr.Errorf("rfc3986: bad scheme — byte №%d has value 0x%02X is not allowed in scheme", byteNumber, b0)
			}
		}

	}

	if len(scheme) <= 0 {
				return "", erorr.Error("rfc3986: bad scheme — scheme must be at least one character long")
	}
	if b0 := scheme[0]; !IsAlpha(rune(b0)) {
				return "", erorr.Errorf("rfc3986: bad scheme — byte №1 has value 0x%02X is not allowed for the first character of scheme", b0)
	}

	return string(scheme), nil
}

// PeekPrefixScheme peeks and returns the 'scheme' as defined by IETF RFC-3986:
//
//	scheme = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
//	
//	ALPHA  = %41-%5A / %61-%7A
//	
//	DIGIT  = %30-%39
func PeekPrefixScheme(str string) (scheme string, n int, found bool) {

	if len(str) <= 0 {
		return "", 0, false
	}

	var buffer [512]byte
	var bytes []byte = buffer[0:0]

	var sawColon bool

	loop: for index, r := range str {
		switch {
		case index <= 0:
			if !IsAlpha(r) {
				return "", 0, false
			}
			r = asciiToLower(r)
			bytes = append(bytes, string(r)...)
		case ':' == r:
			sawColon = true
			break loop
		default:
			if !IsAlpha(r) && !IsDigit(r) && '+' != r && '-' != r && '.' != r {
				return "", 0, false
			}
			r = asciiToLower(r)
			bytes = append(bytes, string(r)...)
		}
	}

	if !sawColon {
		return "", 0, false
	}

	return string(bytes), len(bytes), true
}

func ValidateScheme(scheme string) error {
	var reader io.Reader = strings.NewReader(scheme)

	_, err := ReadScheme(reader)
	return err
}
