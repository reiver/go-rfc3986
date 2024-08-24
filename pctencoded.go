package rfc3986

import (
	"io"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-rfc2234"
)

const pctencodedprefix = '%'

// IsPctEncodedPrefix returns true if what is at the beginning of the string is 'pct-encoded' as defined by IETF RFC-3986.
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

	if !rfc2234.IsHexDig(str1) {
		return false
	}

	if !rfc2234.IsHexDig(str2) {
		return false
	}

	return true
}

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
