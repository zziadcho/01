package functions

import (
	"strings"
	"fmt"
)

// function responsible for parsong the font file
func ParseFont(data string) map[rune][]string {
	startChar := ' '
	blocks := strings.Split(data, "\n\n")
	fontMap := make(map[rune][]string)

	for i, block := range blocks {
		lines := strings.Split(block, "\n")
		if len(lines) > 0 {
			char := rune(startChar + rune(i))
			fontMap[char] = lines
		} else {
			fmt.Printf("warning: empty or malformed block at index %d", i)
		}

	}

	return fontMap
}