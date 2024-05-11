package functions

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}

	firstChar, size := utf8.DecodeRuneInString(s)

	if unicode.IsLower(firstChar) {
		result := strings.ToUpper(string(firstChar)) + s[size:]
		return result
	}
	return s
}
