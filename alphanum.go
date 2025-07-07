package rfc3986

func IsAlphaNum(r rune) bool {
	return IsAlpha(r) || IsDigit(r)
}
