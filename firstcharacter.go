package rfc3986

// firstCharacter to try to future-proof this code a bit, to deal with
// differences between ASCII having single-byte characters, and UTF-8 having variable length characters.
//
// firstCharacter does it the ASCII way, but returns the information in a way that could handle variable length characters.
func firstCharacter(str string) (r rune, n int, found bool) {
	if len(str) < 1 {
		var nada rune
		return nada, 0, false
	}

	var b byte = str[0] // We assume ASCII characters (which are 1 byte long)

	r = rune(b)
	n = 1               // We assume ASCII characters (which are 1 byte long)

	return r, n, true
}
