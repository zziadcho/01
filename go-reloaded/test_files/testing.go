package main

import (
	"fmt"
	"01/go-reloaded/common/functions"
	"01/go-reloaded/common/variables"
	"strconv"
	"strings"
)

func main() {
	content := "hello guy im gonna show you (up, 9)"
	contentSplit := strings.Split(content, " ")
	fmt.Println(len(contentSplit))
	for i, element := range contentSplit {
		if variables.UpFlagMulti.MatchString(element) && i+1 < len(contentSplit) {
			multipiler := contentSplit[i+1]
			multipiler = strings.TrimRight(multipiler, ")")
			multipilerInt, err := strconv.Atoi(multipiler)
			if err != nil {
				fmt.Println("there was a problem converting the Up multiplier to integer:", err)
				continue
			}
			if multipilerInt <= 0 || i-multipilerInt < 0 {
				fmt.Println("the Up multiplier is out of range, lower the number to fix the problem:", multipilerInt)
				continue
			}
			for j := i - multipilerInt; j < i; j++ {
				contentSplit[j] = functions.ToUpper(contentSplit[j])
			}
			convertedWord := functions.ToUpper(contentSplit[i-1])
			contentSplit[i-1] = convertedWord
		}
	}
	fmt.Println(contentSplit)
}
