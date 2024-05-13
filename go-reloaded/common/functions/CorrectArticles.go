package functions

import (
	"strings"
)

func CorrectArticles(text []string) []string {

	definiteArticles := []string{"the"}
	indefiniteArticles := []string{"a", "an"}

	for i, word := range text {
		if contains(definiteArticles, strings.ToLower(word)) {
			text[i] = "the"
		}

		if contains(indefiniteArticles, strings.ToLower(word)) {
			nextWord := ""
			if i+1 < len(text) {
				nextWord = text[i+1]
			}
			if nextWord != "" && (strings.ToLower(nextWord) == "a" || strings.ToLower(nextWord) == "an") {
				text[i] = "an"
			} else {
				text[i] = "a"
			}
		}
	}

	return text
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
