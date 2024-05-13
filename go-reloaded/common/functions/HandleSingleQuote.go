package functions

import (
	"regexp"
	"strings"
)

func HandleSingleQuote(text []string) []string {
	textJoined := strings.Join(text, " ")

	replaceSingleQuotes := regexp.MustCompile(`([a-zA-Z]+)'([a-zA-Z]{1}\s)(\w*)`)
	textJoined = replaceSingleQuotes.ReplaceAllString(textJoined, "$1\\PLACEHOLDER\\$2$3")

	replaceStandaloneQuotes := regexp.MustCompile(`\s*'\s*(.*?)\s*'\s*`)
	textJoined = replaceStandaloneQuotes.ReplaceAllString(textJoined, " '$1' ")

	restoreSingleQuotes := regexp.MustCompile(`\\PLACEHOLDER\\`)
	textJoined = restoreSingleQuotes.ReplaceAllString(textJoined, "'")

	result := strings.Fields(textJoined)
	return result
}
