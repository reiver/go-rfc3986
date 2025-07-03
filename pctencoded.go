package rfc3986

import (
	"io"

	"github.com/reiver/go-erorr"
)

const pctencodedprefix = '%'

// DecodePctEncoded decodes the 2nd and 3rd byte of an IETF RFC-3986 'pct-encoded'.
// (The 1st byte of an IETF RFC-3986 'pct-encoded' is a percent-symbol (i.e., '%'), so we don't include it in the parameters.)
//
// For example, this:
//
//	b, err := rfc3986.DecodePctEncoded('2', 'F')
//
// Would return:
//
//	'/'
func DecodePctEncoded(b2 byte, b3 byte) (byte, error) {

	var result byte

	{
		var b byte = b2

		switch {
		case '0' <= b && b <= '9':
			b -= '0'
		case 'A' <= b && b <= 'F':
			b -= 'A'
			b += 10
		case 'a' <= b && b <= 'f':
			b -= 'a'
			b += 10
		default:
			return 0, erorr.Errorf("rfc3986: the 2nd byte in the pct-encoded string is not an IETF RFC-2234 'HEXDIG'")
		}

		result |= (b << 4)
	}

	{
		var b byte = b3

		switch {
		case '0' <= b && b <= '9':
			b -= '0'
		case 'A' <= b && b <= 'F':
			b -= 'A'
			b += 10
		case 'a' <= b && b <= 'f':
			b -= 'a'
			b += 10
		default:
			return 0, erorr.Errorf("rfc3986: the 3nd byte in the pct-encoded string is not an IETF RFC-2234 'HEXDIG'")
		}

		result |= b
	}

	return result, nil
}

// HasPrefixPctEncoded returns true if what is at the beginning of the string is 'pct-encoded' as defined by IETF RFC-3986.
//
// For example, any of these would return 'true':
//
//	"%2Fdir"
//
//	"%7Ejoeblow"
//
// And, for example, any of these would return 'false':
//
//	"Hello world!"
//
//	"dir%2F"
//
//	"joeblow%40example.com"
//
// 'pct-encoded' is defined by IETF RFC-3986 as follows:
//
//	pct-encoded = "%" HEXDIG HEXDIG
//
// Where 'HEXDIG' is defined in IETF RFC-2234 as:
//
//	HEXDIG = DIGIT / "A" / "B" / "C" / "D" / "E" / "F"
//	
//	DIGIT  = %x30-39
//	       ; 0-9
func HasPrefixPctEncoded(str string) bool {
	var length int = len(str)

	if length < 3 {
		return false
	}

	var str0 rune = rune(str[0])
	var str1 rune = rune(str[1])
	var str2 rune = rune(str[2])

	if !IsPctEncodedPrefix(str0) {
		return false
	}

	if !IsHexDig(str1) {
		return false
	}

	if !IsHexDig(str2) {
		return false
	}

	return true
}

// IsPctEncodedPrefix return 'true' if the rune is the 'pct-encoded' prefix character (i.e., if it is '%'),
// and returns 'false' otherwise.
//
// Note, to determine if a string begins with a 'pct-encoded' instead use [HasPrefixPctEncoded].
func IsPctEncodedPrefix(r rune) bool {
	return pctencodedprefix == r
}

func ReadPctEncodedByte(reader io.Reader) (byte, error) {
	if nil == reader {
		return 0, errNilReader
	}

	var buffer [1]byte
	var p []byte = buffer[:]

	{
		n, err := reader.Read(p)
		if nil != err {
			return 0, erorr.Errorf("rfc3986: could not read the 1st byte in pct-encoded string because: %w", err)
		}
		if 1 != n {
			return 0, erorr.Errorf("rfc3986: problem reading the 1st byte in pct-encoded string — expected to have read 1 byte but actually read %d bytes.", n)
		}

		b0 := buffer[0]

		if !IsPctEncodedPrefix(rune(b0)) {
			return 0, erorr.Errorf("rfc3986:the 1st byte in pct-encoded string is not a percent-sign (%U) (%q) is actually %U (%q)", pctencodedprefix, pctencodedprefix, b0, b0)
		}
	}

	var result byte = 0

	{
		n, err := reader.Read(p)
		if nil != err {
			return 0, erorr.Errorf("rfc3986: could not read the 2nd byte in pct-encoded string because: %w", err)
		}
		if 1 != n {
			return 0, erorr.Errorf("rfc3986: problem reading the 2nd byte in pct-encoded string — expected to have read 1 byte but actually read %d bytes.", n)
		}

		b0 := buffer[0]

		switch {
		case '0' <= b0 && b0 <= '9':
			b0 -= '0'
		case 'A' <= b0 && b0 <= 'F':
			b0 -= 'A'
			b0 += 10
		case 'a' <= b0 && b0 <= 'f':
			b0 -= 'a'
			b0 += 10
		default:
			return 0, erorr.Errorf("rfc3986: the 2nd byte in the pct-encoded string is not an IETF RFC-2234 'HEXDIG'")
		}

		result |= (b0 << 4)
	}

	{
		n, err := reader.Read(p)
		if nil != err {
			return 0, erorr.Errorf("rfc3986: could not read the 3rd byte in pct-encoded string because: %w", err)
		}
		if 1 != n {
			return 0, erorr.Errorf("rfc3986: problem reading the 3rd byte in pct-encoded string — expected to have read 1 byte but actually read %d bytes.", n)
		}

		b0 := buffer[0]

		switch {
		case '0' <= b0 && b0 <= '9':
			b0 -= '0'
		case 'A' <= b0 && b0 <= 'F':
			b0 -= 'A'
			b0 += 10
		case 'a' <= b0 && b0 <= 'f':
			b0 -= 'a'
			b0 += 10
		default:
			return 0, erorr.Errorf("rfc3986: the 2nd byte in the pct-encoded string is not an IETF RFC-2234 'HEXDIG'")
		}

		result |= b0
	}

	return result, nil
}

func WritePctEncodedByte(writer io.Writer, b byte) error {
	if nil == writer {
		return errNilWriter
	}

	var buffer [3]byte = [3]byte{'%','.','.'}

	{
		var x byte = ((b>>4) % 16)

		if x < 10 {
			x += '0'
		} else {
			x += ('A' - 10)
		}

		buffer[1] = x
	}

	{
		var x byte = (b & 0x0f)

		if x < 10 {
			x += '0'
		} else {
			x += ('A' - 10)
		}

		buffer[2] = x
	}

	_, err := writer.Write(buffer[:])
	return err
}
