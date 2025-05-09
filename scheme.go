package rfc3986

import (
	"io"

	"github.com/reiver/go-erorr"
)

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
			if nil != err {
				return "", erorr.Errorf("rfc3986: could not read the byte №%d in scheme string because: %w", byteNumber, err)
			}
			if 1 != n {
				return "", erorr.Errorf("rfc3986: problem reading the byte №%d in scheme string — expected to have read 1 byte but actually read %d bytes.", byteNumber, n)
			}

			b0 := b[0]
			r0 := rune(b0)

			switch {
			case ':' == b0:
		/////////////// BREAK
				break loop
			case IsAlpha(r0) || IsDigit(r0) || '+' == r0 || '-' == r0 || '.' == r0:
				scheme = append(scheme, b0)
			default:
				return "", erorr.Errorf("rfc3986: problem reading the byte №%d in scheme string — byte 0x%02X is not allowed in scheme", byteNumber, b0)
			}
		}

	}

	return string(scheme), nil
}
