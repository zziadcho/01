package functions

import (
	"01/morse-interpeter/common/variables"
	"strings"
)

func MorseCodeInterpeter(code string) string {
	var result strings.Builder
	for _, char := range code {
		if char == ' '{
			result.WriteString(" / ") 
		} else {
			if morse, found := variables.MorseMap[string(char)]; found {
				result.WriteString(morse + " ")
			}
		}
	}
	return result.String()
}
