package functions

import (
	"regexp"
	"strings"
)

func HandlePunctuation(text []string) []string {
	textJoined := strings.Join(text, " ")
	
	target := regexp.MustCompile(`\s*([.,!?:;])`)
	textJoined = target.ReplaceAllString(textJoined, "$1")

	target = regexp.MustCompile(`\s*([.,!?:;]+)\s*`)
	textJoined = target.ReplaceAllString(textJoined, "$1           ")

	result := strings.Split(textJoined, " ")
	return result
}
