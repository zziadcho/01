package functions

import (
	"01/go-reloaded/common/variables"
	"regexp"
)

func RemoveFlagSuffixes(text string) string {
	flags := []*regexp.Regexp{variables.HexFlag, variables.BinFlag, variables.UpFlag, variables.LowFlag, variables.CapFlag, variables.UpFlagMulti, variables.LowFlagMulti, variables.CapFlagMulti}

	for _, flagRegex := range flags {
		text = flagRegex.ReplaceAllString(text, "")
	}
	return text
}
