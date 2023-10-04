package rfc3986

// IsGenDelim returns true if the value of 'r' matches 'sub-delims' as defined in IETF RFC-3986:
//
//	gen-delims  = ":" / "/" / "?" / "#" / "[" / "]" / "@"
func IsGenDelim(r rune) bool {
	switch r {
	case          ':' , '/' , '?' , '#' , '[' , ']' , '@':
		return true
	default:
		return false
	}
}
