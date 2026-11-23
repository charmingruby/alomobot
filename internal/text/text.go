package text

import (
	"unicode"
	"unicode/utf8"
)

func Capitalize(s string) string {
	if s == "" {
		return s
	}

	first, size := utf8.DecodeRuneInString(s)

	return string(unicode.ToUpper(first)) + s[size:]
}
