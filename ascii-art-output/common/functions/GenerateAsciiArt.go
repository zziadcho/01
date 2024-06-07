package functions

import (
	"fmt"
	"os"
	"strings"
)

// ascii art generator function
func GenerateAsciiArt(text string, fontMap map[rune][]string) string {
	var result []string

	for i := 0; i <= 8; i++ {
		result = append(result, "")
	}

	for _, char := range text {
		if char >= ' ' && char <= '~' {
			if charArt, found := fontMap[char]; found {
				for i, line := range charArt {
					result[i] += line
				}
			}
		} else {
			fmt.Println("your input has a non printable character")
			os.Exit(0)
		}

	}
	return strings.Join(result[:len(result)-1], "\n")
}

func GeneratorLoop(splitArg []string, font map[rune][]string) string {
	var generatedArt string
	for _, line := range splitArg {
		if line != "" {
			generatedArt += GenerateAsciiArt(line, font) + "\n"
		} else {
			generatedArt += "\n"
		}
	}
	return generatedArt
}
