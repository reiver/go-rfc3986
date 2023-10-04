package rfc3986

// IsReserved returns true if value of 'r' matches 'reserved' as defined in IETF RFC-3986:
//
//	reserved    = gen-delims / sub-delims
//	
//	gen-delims  = ":" / "/" / "?" / "#" / "[" / "]" / "@"
//	
//	sub-delims  = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
func IsReserved(r rune) bool {
	return IsGenDelim(r) || IsSubDelim(r)
}
