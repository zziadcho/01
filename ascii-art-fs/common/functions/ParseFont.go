package functions

import (
	"strings"
	"fmt"
)

// function responsible for parsong the font file
func ParseFont(data string, flag bool) map[rune][]string {
	startChar := ' '
	var blocks []string
	if flag {
		data = data[1:]
		blocks = strings.Split(data,"\r\n\r\n")
	} else {
		blocks = strings.Split(data, "\n\n")
	}
	fontMap := make(map[rune][]string)

	for i, block := range blocks {
		var lines []string
		if flag{
			lines = strings.Split(block,"\r\n")
		} else {
			lines = strings.Split(block, "\n")
		}
		if len(lines) > 0 {
			char := rune(startChar + rune(i))
			fontMap[char] = lines
		} else {
			fmt.Printf("warning: empty or malformed block at index %d", i)
		}

	}

	return fontMap
}
