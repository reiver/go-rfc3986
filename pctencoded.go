package rfc3986

import (
	"io"
)

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
