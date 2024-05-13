package functions

import "strings"

func correctArticles(text string) string {
	// Split the text into words
	words := strings.Fields(text)

	// Define lists of definite and indefinite articles
	definiteArticles := []string{"the"}
	indefiniteArticles := []string{"a", "an"}

	// Loop through words and correct articles
	for i, word := range words {
		// Check if the word is a definite article and correct it
		if contains(definiteArticles, strings.ToLower(word)) {
			words[i] = "the"
		}

		// Check if the word is an indefinite article and correct it
		if contains(indefiniteArticles, strings.ToLower(word)) {
			nextWord := ""
			if i+1 < len(words) {
				nextWord = words[i+1]
			}
			if nextWord != "" && strings.ToLower(nextWord)[0] == 'a' {
				words[i] = "an"
			} else {
				words[i] = "a"
			}
		}
	}

	// Join the corrected words back into a sentence
	correctedText := strings.Join(words, " ")
	return correctedText
}

// Helper function to check if a slice contains a specific string
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}