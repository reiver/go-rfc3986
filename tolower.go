package rfc3986

func asciiToLower(r rune) rune {

	if 'A' <= r && r <= 'Z' {
		r = r + ('a' - 'A')
	}

	return r
}
