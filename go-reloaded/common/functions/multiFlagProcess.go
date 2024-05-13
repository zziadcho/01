package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func processMultiplier(contentSplit []string, i int, multiplierFlag *regexp.Regexp, converter func(string) string, errMsg string) {
	multiplier := contentSplit[i+1]
	multiplier = strings.TrimRight(multiplier, ")")
	multiplierInt, err := strconv.Atoi(multiplier)
	if err != nil {
		fmt.Println("there was a problem converting the multiplier to integer:", err)
		return
	}
	if multiplierInt <= 0 || i-multiplierInt < 0 {
		fmt.Println(errMsg)
		return
	}
	for j := i - multiplierInt; j < i; j++ {
		contentSplit[j] = converter(contentSplit[j])
	}
	convertedWord := converter(contentSplit[i-1])
	contentSplit[i-1] = convertedWord
}
