package reloaded

import "strings"

func Regulation(s string) bool {
	exceptions := "aeuioh"
	if len(s) > 0 {
		firstLetter := strings.ToLower(string(s[0]))
		return strings.ContainsAny(firstLetter, exceptions)
	}
	return false
}