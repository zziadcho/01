package functions

import(
	"strings"
)

func ArgSplitter(arg string) []string{
	return strings.Split(arg, `\n`)
}