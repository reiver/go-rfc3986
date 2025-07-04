package rfc3986

// IsSubDelim returns true if the value of 'r' matches 'sub-delims' as defined in IETF RFC-3986:
//
//	sub-delims  = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
func IsSubDelim(r rune) bool {
	switch r {
	case          '!' , '$' , '&' ,'\'' , '(' , ')' , '*' , '+' , ',' , ';' , '=':
		return true
	default:
		return false
	}
}

func PeekPrefixSubDelims(str string) (rune, int, bool) {
	r, n, found := firstCharacter(str)
	if !found {
		var nada rune
		return nada, 0, false
	}

	if !IsSubDelim(r) {
		var nada rune
		return nada, 0, false
	}

	return r, n, true
}
