package main

import (
	"01/go-reloaded/common/variables"
	"fmt"
)

func main() {
	text := "it was a (up, 5) beautiful day to begin!"
	text2 := []string{}
	text2 = append(text2, text)
	for _, element := range text2 {
		if variables.UpFlagMulti.MatchString(element) {
			fmt.Println("found!")
			match := variables.UpFlagMulti.FindStringSubmatch(element)
			if match != nil {
				numberString := match[1]
			fmt.Println(numberString)
		}
	}
}
}