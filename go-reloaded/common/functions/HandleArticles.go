package functions

import "strings"

func AdjustArticles(text []string) []string {
	for i, word := range text {
		if word == "an" {
			text[i] = "a"
		}
		if word == "An" || word == "AN" {
			text[i] = "A"
		}
	}

	for i := 0; i < len(text)-1; i++ {
		if text[i] == "a" && NeedsAnAdjustment(text[i+1]) {
			text[i] = "an"
		}
		if text[i] == "A" && NeedsAnAdjustment(text[i+1]) {
			text[i] = "An"
		}
	}
	return text
}

func NeedsAnAdjustment(nextWord string) bool {
	vowels := "aeiouAEIOU"
	if len(nextWord) > 0 && strings.Contains(vowels, string(nextWord[0])) {
		return true
	}
	return false
}
