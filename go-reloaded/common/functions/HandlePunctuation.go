package functions

import (
	"regexp"
	"strings"
)

func HandlePunctuation(text []string) []string {
	textJoin := strings.Join(text, " ")
	target := regexp.MustCompile(`\s*([.,!?:;])`)
	textJoin = target.ReplaceAllString(textJoin, "$1")
	target = regexp.MustCompile(`\s*([.,!?:;]+)\s*`)
	textJoin = target.ReplaceAllString(textJoin, "$1           ")
	last := strings.Fields(textJoin)
	return last
}
